package cas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/images"
	"github.com/lf-edge/eve/pkg/pillar/containerd"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/opencontainers/go-digest"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	v1types "github.com/google/go-containerregistry/pkg/v1/types"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	spec "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus" // XXX add log argument
)

const (
	casClientType = "containerd"
	// relative path to rootfs for an individual container
	containerRootfsPath = "rootfs/"
	// container config file name
	imageConfigFilename = "image-config.json"
	// contains conatiner's image name.
	imageNameFilename = "image-name"
	// start of containerd gc ref label for children in content store
	containerdGCRef = "containerd.io/gc.ref.content"
)

type containerdCAS struct {
}

//CheckBlobExists: returns true if the blob exists. Arg 'blobHash' should be of format sha256:<hash>.
func (c *containerdCAS) CheckBlobExists(blobHash string) bool {
	_, err := containerd.CtrGetBlobInfo(blobHash)
	return err == nil
}

//GetBlobInfo: returns BlobInfo of type BlobInfo for the given blobHash.
// Arg 'blobHash' should be of format sha256:<hash>.
//Returns error if no blob is found for the given 'blobHash'.
func (c *containerdCAS) GetBlobInfo(blobHash string) (*BlobInfo, error) {
	info, err := containerd.CtrGetBlobInfo(blobHash)
	if err != nil {
		return nil, fmt.Errorf("GetBlobInfo: Exception while getting size of blob: %s. %s", blobHash, err.Error())
	}

	return &BlobInfo{
		Digest: info.Digest.String(),
		Size:   info.Size,
		Labels: info.Labels,
	}, nil
}

//ListBlobInfo: returns list of BlobInfo for all the blob present in CAS
func (c *containerdCAS) ListBlobInfo() ([]*BlobInfo, error) {
	infos, err := containerd.CtrListBlobInfo()
	if err != nil {
		return nil, fmt.Errorf("ListBlobInfo: Exception while getting blob list. %s", err.Error())
	}
	blobInfos := make([]*BlobInfo, 0)
	for _, info := range infos {
		blobInfos = append(blobInfos, &BlobInfo{
			Digest: info.Digest.String(),
			Size:   info.Size,
			Labels: info.Labels,
		})
	}
	return blobInfos, nil
}

// ListBlobsMediaTypes get a map of all blobs and their media types.
// If a blob does not have a media type, it is not returned here.
// If you want *all* blobs, whether or not it has a type, use ListBlobInfo
func (c *containerdCAS) ListBlobsMediaTypes() (map[string]string, error) {
	hashMap := map[string]string{}

	// start with all of the images
	imageObjectList, err := containerd.CtrListImages()
	if err != nil {
		return nil, fmt.Errorf("ListBlobsMediaTypes: Exception while getting image list. %s", err.Error())
	}
	// save the root and type of each image
	for _, i := range imageObjectList {
		dig, mediaType := i.Target.Digest.String(), i.Target.MediaType
		hashMap[dig] = mediaType
		switch v1types.MediaType(mediaType) {
		case v1types.OCIImageIndex, v1types.DockerManifestList:
			index, err := getIndexManifest(c, dig)
			if err != nil {
				return nil, fmt.Errorf("ListBlobsMediaTypes: could not get index for %s", dig)
			}
			// save all of the manifests
			for _, m := range index.Manifests {
				hashMap[m.Digest.String()] = string(m.MediaType)
				// and now read each manifest
				manifest, err := getManifest(c, m.Digest.String())
				if err != nil {
					return nil, fmt.Errorf("ListBlobsMediaTypes: could not get manifest for %s", dig)
				}
				// read the config and the layers
				hashMap[manifest.Config.Digest.String()] = string(manifest.Config.MediaType)
				for _, l := range manifest.Layers {
					hashMap[l.Digest.String()] = string(l.MediaType)
				}
			}
		case v1types.OCIManifestSchema1, v1types.DockerManifestSchema1, v1types.DockerManifestSchema2, v1types.DockerManifestSchema1Signed:
			manifest, err := getManifest(c, dig)
			if err != nil {
				return nil, fmt.Errorf("ListBlobsMediaTypes: could not get manifest for %s", dig)
			}
			// read the config and the layers
			hashMap[manifest.Config.Digest.String()] = string(manifest.Config.MediaType)
			for _, l := range manifest.Layers {
				hashMap[l.Digest.String()] = string(l.MediaType)
			}
		}
	}
	return hashMap, nil
}

// IngestBlob: parses the given one or more `blobs` (BlobStatus) and for each blob reads the blob data from
// BlobStatus.Path and ingests it into CAS's blob store.
// Returns a list of loaded BlobStatus and an error is thrown if the read blob's hash does not match with the
// respective BlobStatus.Sha256 or if there is an exception while reading the blob data.
// In case of exception, the returned list of loaded blobs will contain all the blob that were loaded until that point.
func (c *containerdCAS) IngestBlob(blobs ...*types.BlobStatus) ([]*types.BlobStatus, error) {
	var (
		index          *ocispec.Index
		indexHash      string
		manifests      = make([]*ocispec.Manifest, 0)
		manifestHashes = make([]string, 0)
	)
	loadedBlobs := make([]*types.BlobStatus, 0)

	//Step 1: Load blobs into CAS
	for _, blob := range blobs {
		log.Infof("IngestBlob(%s): processing blob %+v", blob.Sha256, blob)
		//Process the blob only if its not loaded already
		if blob.State == types.LOADED {
			log.Infof("IngestBlob(%s): Not loading blob as it is already loaded", blob.Sha256)
			loadedBlobs = append(loadedBlobs, blob)
			continue
		}
		log.Infof("IngestBlob(%s): Attempting to load blob", blob.Sha256)
		var (
			r        io.Reader
			blobFile = blob.Path
			// the sha MUST be lower-case for it to work with the ocispec utils
			sha = fmt.Sprintf("%s:%s", digest.SHA256, strings.ToLower(blob.Sha256))
		)

		//Step 1.1: Read the blob from verified dir
		fileReader, err := os.Open(blobFile)
		if err != nil {
			err = fmt.Errorf("IngestBlob(%s): could not open blob file for reading at %s: %+s",
				blob.Sha256, blobFile, err.Error())
			log.Errorf(err.Error())
			return loadedBlobs, err
		}
		defer fileReader.Close()

		//Step 1.2: Resolve blob type and if this is a manifest or index, we will need to process (parse) it accordingly
		switch blob.BlobType {
		case types.BlobIndex:
			// read it in so we can process it
			data, err := ioutil.ReadAll(fileReader)
			if err != nil {
				err = fmt.Errorf("IngestBlob(%s): could not read data at %s: %+s",
					blob.Sha256, blobFile, err.Error())
				log.Errorf(err.Error())
				return loadedBlobs, err
			}
			fileReader.Close()
			// create a new reader for the content.WriteBlob
			r = bytes.NewReader(data)
			// try to parse the index
			if err := json.Unmarshal(data, &index); err != nil {
				err = fmt.Errorf("IngestBlob(%s): could not parse index at %s: %+s",
					blob.Sha256, blobFile, err.Error())
				log.Errorf(err.Error())
				return loadedBlobs, err
			}
			indexHash = sha
		case types.BlobManifest:
			// read it in so we can process it
			data, err := ioutil.ReadAll(fileReader)
			if err != nil {
				err = fmt.Errorf("IngestBlob(%s): could not read data at %s: %+s",
					blob.Sha256, blobFile, err.Error())
				log.Errorf(err.Error())
				return loadedBlobs, err
			}
			fileReader.Close()
			// create a new reader for the content.WriteBlob
			r = bytes.NewReader(data)
			// try to parse the index
			mfst := ocispec.Manifest{}
			if err := json.Unmarshal(data, &mfst); err != nil {
				err = fmt.Errorf("IngestBlob(%s): could not parse manifest at %s: %+s",
					blob.Sha256, blobFile, err.Error())
				log.Errorf(err.Error())
				return loadedBlobs, err
			}
			manifests = append(manifests, &mfst)
			manifestHashes = append(manifestHashes, sha)
		default:
			// do nothing special, just pass it on
			r = fileReader
		}

		//Step 1.3: Ingest the blob into CAS
		if err := containerd.CtrWriteBlob(sha, blob.Size, r); err != nil {
			err = fmt.Errorf("IngestBlob(%s): could not load blob file into containerd at %s: %+s",
				blob.Sha256, blobFile, err.Error())
			log.Errorf(err.Error())
			return loadedBlobs, err
		}
		log.Infof("IngestBlob(%s): Loaded the blob successfully", blob.Sha256)
		blob.State = types.LOADED
		loadedBlobs = append(loadedBlobs, blob)
	}

	//Step 2: Walk the tree from the root to add the necessary labels
	if index != nil {
		info := BlobInfo{
			Digest: indexHash,
			Labels: map[string]string{},
		}
		for i, m := range index.Manifests {
			info.Labels[fmt.Sprintf("%s.%d", containerdGCRef, i)] = m.Digest.String()
		}
		if err := c.UpdateBlobInfo(info); err != nil {
			err = fmt.Errorf("IngestBlob(%s): could not update labels on index: %v", info.Digest, err.Error())
			log.Errorf(err.Error())
			return loadedBlobs, err
		}
	}

	if len(manifests) > 0 {
		for j, m := range manifests {
			info := BlobInfo{
				Digest: manifestHashes[j],
				Labels: map[string]string{},
			}
			for i, l := range m.Layers {
				info.Labels[fmt.Sprintf("%s.%d", containerdGCRef, i)] = l.Digest.String()
			}
			i := len(m.Layers)
			info.Labels[fmt.Sprintf("%s.%d", containerdGCRef, i)] = m.Config.Digest.String()

			if err := c.UpdateBlobInfo(info); err != nil {
				err = fmt.Errorf("IngestBlob(%s): could not update labels on manifest: %v",
					info.Digest, err.Error())
				log.Errorf(err.Error())
				return loadedBlobs, err
			}
		}

	}
	return loadedBlobs, nil
}

//UpdateBlobInfo updates BlobInfo of a blob in CAS.
//Arg is BlobInfo type struct in which BlobInfo.Digest is mandatory, and other field are to be filled
// only if its needed to be updated
//Returns error is no blob is found match blobInfo.Digest
func (c *containerdCAS) UpdateBlobInfo(blobInfo BlobInfo) error {
	existingBlobIfo, err := c.GetBlobInfo(blobInfo.Digest)
	if err != nil {
		err = fmt.Errorf("UpdateBlobInfo: Exception while fetching existing blobInfo of %s: %s", blobInfo.Digest, err.Error())
		log.Error(err.Error())
		return err
	}

	changed := false
	updatedContentInfo := content.Info{
		Digest: digest.Digest(blobInfo.Digest),
	}

	updatedFields := make([]string, 0)
	if blobInfo.Size > 0 && blobInfo.Size != existingBlobIfo.Size {
		updatedFields = append(updatedFields, "size")
		updatedContentInfo.Size = blobInfo.Size
		changed = true
	}

	if blobInfo.Labels != nil {
		for k := range blobInfo.Labels {
			updatedFields = append(updatedFields, fmt.Sprintf("labels.%s", k))
		}
		updatedContentInfo.Labels = blobInfo.Labels
		changed = true
	}

	if changed {
		if err := containerd.CtrUpdateBlobInfo(updatedContentInfo, updatedFields); err != nil {
			err = fmt.Errorf("UpdateBlobInfo: Exception while updating blobInfo of %s: %s",
				blobInfo.Digest, err.Error())
			log.Error(err.Error())
			return err
		}
	}
	return nil
}

//ReadBlob: returns a reader to consume the raw data of the blob which matches the given arg 'blobHash'.
//Returns error if no blob is found for the given 'blobHash'.
//Arg 'blobHash' should be of format sha256:<hash>.
func (c *containerdCAS) ReadBlob(blobHash string) (io.Reader, error) {
	reader, err := containerd.CtrReadBlob(blobHash)
	if err != nil {
		log.Errorf("ReadBlob: Exception while reading blob: %s. %s", blobHash, err.Error())
		return nil, err
	}
	return reader, nil
}

//RemoveBlob: removes a blob which matches the given arg 'blobHash'.
//To keep this method idempotent, no error is returned if the given arg 'blobHash' does not match any blob.
//Arg 'blobHash' should be of format sha256:<hash>.
func (c *containerdCAS) RemoveBlob(blobHash string) error {
	if err := containerd.CtrDeleteBlob(blobHash); err != nil && !isNotFoundError(err) {
		return fmt.Errorf("RemoveBlob: Exception while removing blob: %s. %s", blobHash, err.Error())
	}
	return nil
}

//Children: returns a list of child blob hashes if the given arg 'blobHash' belongs to a
// index or a manifest blob, else an empty list is returned.
//Format of returned blob hash list and arg 'blobHash' is sha256:<hash>.
func (c *containerdCAS) Children(blobHash string) ([]string, error) {
	if _, err := c.ReadBlob(blobHash); err != nil {
		return nil, fmt.Errorf("Children: Exception while reading blob %s. %s", blobHash, err.Error())
	}
	childBlobSha256 := make([]string, 0)
	index, err := getIndexManifest(c, blobHash)
	if err == nil && index.Manifests != nil {
		for _, manifest := range index.Manifests {
			childBlobSha256 = append(childBlobSha256, manifest.Digest.String())
		}
	} else {
		manifest, err := getManifest(c, blobHash)
		if err != nil {
			return childBlobSha256, nil
		}
		childBlobSha256 = append(childBlobSha256, manifest.Config.Digest.String())
		for _, layer := range manifest.Layers {
			childBlobSha256 = append(childBlobSha256, layer.Digest.String())
		}
	}
	return childBlobSha256, nil
}

//CreateImage: creates a reference which points to a blob with 'blobHash'. 'blobHash' must belong to a index blob
//Arg 'blobHash' should be of format sha256:<hash>.
//Returns error if no blob is found matching the given 'blobHash' or if the given 'blobHash' does not belong to an index.
func (c *containerdCAS) CreateImage(reference, blobHash string) error {
	manifest, mediaType, err := getManifestAndMediaType(c, blobHash)
	if err != nil {
		return fmt.Errorf("CreateImage: exception while parsing blob %s: %s", blobHash, err.Error())
	}
	image := images.Image{
		Name:   reference,
		Labels: nil,
		Target: spec.Descriptor{
			MediaType: mediaType,
			Digest:    digest.Digest(blobHash),
			Size:      manifest.Config.Size,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}
	_, err = containerd.CtrCreateImage(image)
	if err != nil {
		return fmt.Errorf("CreateImage: Exception while creating reference: %s. %s", reference, err.Error())
	}
	return nil
}

//GetImageHash: returns a blob hash of format sha256:<hash> which the given 'reference' is pointing to.
// Returns error if the given 'reference' is not found.
func (c *containerdCAS) GetImageHash(reference string) (string, error) {
	image, err := containerd.CtrGetImage(reference)
	if err != nil {
		return "", fmt.Errorf("GetImageHash: Exception while getting image: %s. %s", reference, err.Error())
	}
	return image.Target().Digest.String(), nil
}

//ListImages: returns a list of references
func (c *containerdCAS) ListImages() ([]string, error) {
	imageObjectList, err := containerd.CtrListImages()
	if err != nil {
		return nil, fmt.Errorf("ListImages: Exception while getting image list. %s", err.Error())
	}
	imageNameList := make([]string, 0)
	for _, image := range imageObjectList {
		imageNameList = append(imageNameList, image.Name)
	}
	return imageNameList, nil
}

//RemoveImage removes an reference from CAS
//To keep this method idempotent, no error  is returned if the given 'reference' is not found.
func (c *containerdCAS) RemoveImage(reference string) error {
	if err := containerd.CtrDeleteImage(reference); err != nil {
		return fmt.Errorf("RemoveImage: Exception while removing image. %s", err.Error())
	}
	return nil
}

//ReplaceImage: replaces the blob hash to which the given 'reference' is pointing to with the given 'blobHash'.
//Returns error if the given 'reference' or a blob matching the given arg 'blobHash' is not found.
//Returns if the given 'blobHash' does not belong to an index.
//Arg 'blobHash' should be of format sha256:<hash>.
func (c *containerdCAS) ReplaceImage(reference, blobHash string) error {
	manifest, mediaType, err := getManifestAndMediaType(c, blobHash)
	if err != nil {
		return fmt.Errorf("CreateImage: exception while parsing blob %s: %s", blobHash, err.Error())
	}
	image := images.Image{
		Name:   reference,
		Labels: nil,
		Target: spec.Descriptor{
			MediaType: mediaType,
			Digest:    digest.Digest(blobHash),
			Size:      manifest.Config.Size,
		},
	}
	if _, err := containerd.CtrUpdateImage(image, "target"); err != nil {
		return fmt.Errorf("ReplaceImage: Exception while updating reference: %s. %s", reference, err.Error())
	}
	return nil
}

//CreateSnapshotForImage: creates an snapshot with the given snapshotID for the given 'reference'
//Arg 'snapshotID' should be of format sha256:<hash>.
func (c *containerdCAS) CreateSnapshotForImage(snapshotID, reference string) error {
	clientImageObj, err := containerd.CtrGetImage(reference)
	if err != nil {
		return fmt.Errorf("CreateSnapshotForImage: Exception while getting clientImageObj: %s. %s", reference, err.Error())
	}
	if err := containerd.UnpackClientImage(clientImageObj); err != nil {
		err = fmt.Errorf("CreateSnapshotForImage: could not unpack clientImageObj %s: %+s",
			clientImageObj.Name(), err.Error())
		log.Errorf(err.Error())
		return err
	}

	if _, err := containerd.CtrPrepareSnapshot(snapshotID, clientImageObj); err != nil {
		return fmt.Errorf("CreateSnapshotForImage: Exception while creating snapshot: %s. %s", snapshotID, err.Error())
	}
	return nil
}

//MountSnapshot: mounts the snapshot on the given target path
//Arg 'snapshotID' should be of format sha256:<hash>.
func (c *containerdCAS) MountSnapshot(snapshotID, targetPath string) error {
	if err := containerd.CtrMountSnapshot(snapshotID, targetPath); err != nil {
		return fmt.Errorf("MountSnapshot: Exception while fetching mounts of snapshot: %s. %s", snapshotID, err)
	}
	return nil
}

//ListSnapshots: returns a list of snapshotIDs where each entry is of format sha256:<hash>.
func (c *containerdCAS) ListSnapshots() ([]string, error) {
	snapshotInfoList, err := containerd.CtrListSnapshotInfo()
	if err != nil {
		return nil, fmt.Errorf("ListSnapshots: unable to get snapshot info list: %s", err.Error())
	}
	snapshotIDList := make([]string, 0)
	for _, snapshotInfo := range snapshotInfoList {
		snapshotIDList = append(snapshotIDList, snapshotInfo.Name)
	}
	return snapshotIDList, nil
}

//ListSnapshots: removes a snapshot matching the given 'snapshotID'.
//Arg 'snapshotID' should be of format sha256:<hash>.
//To keep this method idempotent, no error  is returned if the given 'snapshotID' is not found.
func (c *containerdCAS) RemoveSnapshot(snapshotID string) error {
	if err := containerd.CtrRemoveSnapshot(snapshotID); err != nil && !isNotFoundError(err) {
		return fmt.Errorf("RemoveSnapshot: Exception while removing snapshot: %s. %s", snapshotID, err.Error())
	}
	return nil
}

// PrepareContainerRootDir prepares a writable snapshot from the reference. Before preparing container's root directory,
// this API removes any existing state that may have accumulated (like existing snapshots being available, etc.)
// This effectively voids any kind of caching, but on the flip side frees us
// from cache invalidation. Additionally this API should deposit an OCI config json file and image name
// next to the rootfs so that the effective structure becomes:
//    rootPath/rootfs, rootPath/image-config.json
// The rootPath is expected to end in a basename that becomes the snapshotID
func (c *containerdCAS) PrepareContainerRootDir(rootPath, reference, rootBlobSha string) error {
	//Step 1: On device restart, the existing bundle is not deleted, we need to delete the
	// existing bundle of the container and recreate it. This is safe to run even
	// when bundle doesn't exist
	if c.RemoveContainerRootDir(rootPath) != nil {
		log.Warnf("PrepareContainerRootDir: tried to clean up any existing state, hopefully it worked")
	}

	//Step 2: create snapshot of the image so that it can be mounted as container's rootfs.
	snapshotID := containerd.GetSnapshotID(rootPath)
	if err := c.CreateSnapshotForImage(snapshotID, reference); err != nil {
		err = fmt.Errorf("PrepareContainerRootDir: Could not create snapshot %s. %v", snapshotID, err)
		log.Errorf(err.Error())
		return err
	}

	//Step 3: write OCI image config/spec json under the container's rootPath.
	clientImageSpec, err := getImageConfig(c, reference)
	if err != nil {
		err = fmt.Errorf("PrepareContainerRootDir: exception while fetching image config for reference %s: %s",
			reference, err.Error())
		log.Errorf(err.Error())
		//return err
	}
	mountpoints := clientImageSpec.Config.Volumes
	execpath := clientImageSpec.Config.Entrypoint
	cmd := clientImageSpec.Config.Cmd
	workdir := clientImageSpec.Config.WorkingDir
	unProcessedEnv := clientImageSpec.Config.Env
	log.Infof("PrepareContainerRootDir: mountPoints %+v execpath %+v cmd %+v workdir %+v env %+v",
		mountpoints, execpath, cmd, workdir, unProcessedEnv)
	clientImageSpecJSON, err := getJSON(clientImageSpec)
	if err != nil {
		err = fmt.Errorf("PrepareContainerRootDir: Could not build json of image: %v. %v",
			reference, err.Error())
		log.Errorf(err.Error())
		return err
	}

	if err := os.MkdirAll(rootPath, 0766); err != nil {
		err = fmt.Errorf("PrepareContainerRootDir: Exception while creating rootPath dir. %v", err)
		log.Errorf(err.Error())
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(rootPath, imageConfigFilename), []byte(clientImageSpecJSON), 0666); err != nil {
		err = fmt.Errorf("PrepareContainerRootDir: Exception while writing image info to %v/%v. %v",
			rootPath, imageConfigFilename, err)
		log.Errorf(err.Error())
		return err
	}
	return nil
}

// RemoveContainerRootDir removes contents of a container's rootPath and snapshot.
func (c *containerdCAS) RemoveContainerRootDir(rootPath string) error {
	//Step 1: Un-mount container's rootfs
	if err := syscall.Unmount(filepath.Join(rootPath, containerRootfsPath), 0); err != nil {
		err = fmt.Errorf("RemoveContainerRootDir: exception while unmounting: %v/%v. %v",
			rootPath, containerRootfsPath, err)
		log.Error(err.Error())
		return err
	}

	//Step 2: Clean container rootPath
	if err := os.RemoveAll(rootPath); err != nil {
		err = fmt.Errorf("RemoveContainerRootDir: exception while deleting: %v. %v", rootPath, err)
		log.Error(err.Error())

		return err

	}

	//Step 3: Remove snapshot created for the image
	snapshotID := containerd.GetSnapshotID(rootPath)
	if err := c.RemoveSnapshot(snapshotID); err != nil {
		err = fmt.Errorf("RemoveContainerRootDir: unable to remove snapshot: %v. %v", snapshotID, err)
		log.Error(err.Error())

		return err

	}
	return nil
}

// IngestBlobsAndCreateImage is a combination of IngestBlobs and CreateImage APIs,
// but this API will add a lease, upload all the blobs, add reference to the blobs and release the lease.
// By adding a lock before uploading the blobs we prevent the unreferenced blobs from getting GCed.
// We will assume that the first blob in the list will be the root blob for which the reference will be created.
// Returns an an error if the read blob's hash does not match with the respective BlobStatus.Sha256 or
// if there is an exception while reading the blob data.
// NOTE: This API either loads all the blobs or loads nothing. In other words, in case of error,
// this API will GC all blobs that were loaded until that point
func (c *containerdCAS) IngestBlobsAndCreateImage(reference string, blobs ...*types.BlobStatus) ([]*types.BlobStatus, error) {

	log.Infof("IngestBlobsAndCreateImage: Attempting to Ingest %d blobs and add reference: %s", len(blobs), reference)
	loadedBlobs := make([]*types.BlobStatus, 0)
	// This will be called in case of any error.
	// This is to make sure that we delete loadedBlobs in case of an error as we wouldn't have added a reference
	// to the blobs, which makes the loadedBlobs vulnerable for garbage collection.
	var cleanUpLoadedBlobs = func() {
		for _, blob := range loadedBlobs {
			c.RemoveBlob(blob.Sha256)
		}
	}
	deleteLease, err := containerd.CtrCreateLease()
	if err != nil {
		err = fmt.Errorf("IngestBlobsAndCreateImage: Unable load blobs for reference %s. "+
			"Exception while creating lease: %v", reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
	defer deleteLease()
	loadedBlobs, err = c.IngestBlob(blobs...)
	if err != nil {
		err = fmt.Errorf("IngestBlobsAndCreateImage: Exception while loading blobs into CAS: %v", err.Error())
		log.Errorf(err.Error())
		cleanUpLoadedBlobs()
		return nil, err
	}
	rootBlobSha := fmt.Sprintf("%s:%s", digest.SHA256, strings.ToLower(blobs[0].Sha256))
	imageHash, err := c.GetImageHash(reference)
	log.Infof("IngestBlobsAndCreateImage: creating/updating reference: %s for rootBlob %s", reference, rootBlobSha)
	if err != nil || imageHash == "" {
		if err := c.CreateImage(reference, rootBlobSha); err != nil {
			err = fmt.Errorf("IngestBlobsAndCreateImage: could not reference %s with rootBlob %s: %v",
				reference, rootBlobSha, err.Error())
			log.Errorf(err.Error())
			cleanUpLoadedBlobs()
			return nil, err
		}
	} else {
		if err := c.ReplaceImage(reference, rootBlobSha); err != nil {
			err = fmt.Errorf("IngestBlobsAndCreateImage: could not update reference %s with rootBlob %s: %v",
				reference, rootBlobSha, err.Error())
			log.Errorf(err.Error())
			cleanUpLoadedBlobs()
			return nil, err
		}
	}
	return loadedBlobs, nil
}

//CloseClient closes the containerd CAS client initialized while calling `NewCAS()`
func (c *containerdCAS) CloseClient() error {
	if err := containerd.CloseClient(); err != nil {
		err = fmt.Errorf("CloseClient: Exception while closinn %s CAS client: %s", casClientType, err.Error())
		log.Error(err.Error())
		return err
	}
	return nil
}

//newContainerdCAS: constructor for containerd CAS
func newContainerdCAS() CAS {
	if err := containerd.InitContainerdClient(); err != nil {
		log.Fatalf("newContainerdCAS: excpetion while initializing containerd client: %s", err.Error())
	}
	return &containerdCAS{}
}

//getIndexManifest: returns a indexManifest by parsing the given blobSha256
func getIndexManifest(c *containerdCAS, blobSha256 string) (*v1.IndexManifest, error) {
	reader, err := c.ReadBlob(blobSha256)
	if err != nil {
		return nil, fmt.Errorf("getIndexManifest: Exception while reading blob: %s. %s", blobSha256, err)
	}
	index, err := v1.ParseIndexManifest(reader)
	if err != nil {
		return nil, fmt.Errorf("getIndexManifest: Exception while reading blob Index: %s. %s", blobSha256, err.Error())
	}
	return index, nil
}

//getManifestFromIndex: returns Manifest for the current architecture from IndexManifest
func getManifestFromIndex(c *containerdCAS, indexManifest *v1.IndexManifest) (*v1.Manifest, error) {
	manifestSha256, err := getManifestBlobSha256FromIndex(indexManifest)
	if err != nil {
		return nil, fmt.Errorf("getManifestFromIndex: Exception while fetching manifest sha256: %s", err.Error())
	}
	return getManifest(c, manifestSha256)
}

//getManifest: returns manifest as type v1.Manifest byr parsing the given blobSha256
func getManifest(c *containerdCAS, blobSha256 string) (*v1.Manifest, error) {
	reader, err := c.ReadBlob(blobSha256)
	if err != nil {
		return nil, fmt.Errorf("getManifest: Exception while reading blob: %s. %s", blobSha256, err.Error())
	}
	manifest, err := v1.ParseManifest(reader)
	if err != nil {
		return nil, fmt.Errorf("getManifest: Exception while reading blob Manifest: %s. %s", blobSha256, err.Error())
	}
	return manifest, nil
}

//getManifestBlobSha256FromIndex: return blobSha256 of an manifest for the current architecture
func getManifestBlobSha256FromIndex(indexManifest *v1.IndexManifest) (string, error) {
	if indexManifest.Manifests == nil {
		return "", fmt.Errorf("getManifestBlobSha256FromIndex: No manifests found in index")
	}
	for _, m := range indexManifest.Manifests {
		if m.Platform.Architecture == runtime.GOARCH {
			return m.Digest.String(), nil
		}
	}
	return "", fmt.Errorf("getManifestBlobSha256FromIndex: No manifest found in the Index for arch: %s", runtime.GOARCH)
}

//isNotFoundError: returns true if the given error is a "not found" error
func isNotFoundError(err error) bool {
	return strings.HasSuffix(err.Error(), "not found")
}

func getManifestAndMediaType(c *containerdCAS, blobHash string) (*v1.Manifest, string, error) {
	var (
		manifest  *v1.Manifest
		mediaType string
	)
	index, err := getIndexManifest(c, blobHash)
	if err != nil {
		return nil, "", fmt.Errorf("getManifestAndMediaType: Exception while parsing blob as IndexManifest. %s", err.Error())
	}
	if index.Manifests == nil {
		//Not an index. Check if the blob is manifest
		manifest, err = getManifest(c, blobHash)
		if err != nil {
			return nil, "", fmt.Errorf("getManifestAndMediaType: Exception while parsing blob as Manifest. %s", err.Error())
		}
		mediaType = images.MediaTypeDockerSchema2Manifest
	} else {
		manifest, err = getManifestFromIndex(c, index)
		if err != nil {
			return nil, "", fmt.Errorf("getManifestAndMediaType: Exception while fetching Manifest. %s", err.Error())
		}
		mediaType = images.MediaTypeDockerSchema2ManifestList
	}

	return manifest, mediaType, nil
}

//getImageConfig returns imageConfig for a reference
func getImageConfig(c *containerdCAS, reference string) (*ocispec.Image, error) {
	index := ocispec.Index{}
	manifests := ocispec.Manifest{}
	imageConfig := ocispec.Image{}

	//Step 1: Get the hash of parent blob
	imageParentHash, err := c.GetImageHash(reference)
	if err != nil {
		err = fmt.Errorf("getImageConfig: exception while fetching reference hash of %s: %s", reference, err.Error())

	}

	//Step 2: Read the parent blob data
	blobReader, err := c.ReadBlob(imageParentHash)
	if err != nil {
		err = fmt.Errorf("getImageConfig: exception while reading blob %s for reference %s: %s",
			imageParentHash, reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
	blobData, err := ioutil.ReadAll(blobReader)
	if err != nil {
		err = fmt.Errorf("getImageConfig: could not read blobdata %s for reference %s: %+s",
			imageParentHash, reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}

	//Step 3: Get the manifest of the image

	//Step 3.1: Check if the blob is an index
	if err := json.Unmarshal(blobData, &index); err != nil || index.Manifests == nil {
		//Step 3.2: Check if the blob is an manifest
		if err := json.Unmarshal(blobData, &manifests); err != nil {
			err = fmt.Errorf("getImageConfig: could not read imageManifest %s for reference %s: %+s",
				imageParentHash, reference, err.Error())
			log.Errorf(err.Error())
			return nil, err
		}
	} else {
		//Step 3.1.1: Fetch manifest hash from index
		for _, m := range index.Manifests {
			//Step 3.1.2:  get the appropriate manifest has from index
			if m.Platform.Architecture == runtime.GOARCH {
				blobReader, err = c.ReadBlob(m.Digest.String())
				if err != nil {
					err = fmt.Errorf("getImageConfig: exception while reading manifest blob %s for reference %s: %s",
						m.Digest.String(), reference, err.Error())
					log.Errorf(err.Error())
					return nil, err
				}
				//Step 3.1.3: Read the manifest data
				blobData, err = ioutil.ReadAll(blobReader)
				if err != nil {
					err = fmt.Errorf("getImageConfig: could not parsr manifestBlob %s for reference %s: %+s",
						m.Digest.String(), reference, err.Error())
					log.Errorf(err.Error())
					return nil, err
				}
				if err := json.Unmarshal(blobData, &manifests); err != nil {
					err = fmt.Errorf("getImageConfig: could not parse manifestBlob %s for reference %s: %+s",
						m.Digest.String(), reference, err.Error())
					log.Errorf(err.Error())
					return nil, err
				}
				break
			}
		}
	}

	//Step 4: Get the config hash from manifest and read the config data
	configHash := manifests.Config.Digest.String()
	blobReader, err = c.ReadBlob(configHash)
	if err != nil {
		err = fmt.Errorf("getImageConfig: exception while reading config blob %s for reference %s: %s",
			configHash, reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
	blobData, err = ioutil.ReadAll(blobReader)
	if err != nil {
		err = fmt.Errorf("getImageConfig: could not read config blobdata %s for reference %s: %+s",
			configHash, reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
	if err := json.Unmarshal(blobData, &imageConfig); err != nil {
		err = fmt.Errorf("getImageConfig: could not parse configBlob %s for reference %s: %+s",
			configHash, reference, err.Error())
		log.Errorf(err.Error())
		return nil, err
	}
	return &imageConfig, nil
}

// getJSON - returns input in JSON format
func getJSON(x interface{}) (string, error) {
	b, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		return "", fmt.Errorf("getJSON: Exception while marshalling container spec JSON. %v", err)
	}
	return fmt.Sprint(string(b)), nil
}
