module github.com/lf-edge/eve/pkg/pillar

go 1.16

require (
	cloud.google.com/go/storage v1.21.0 // indirect
	github.com/Focinfi/go-dns-resolver v1.0.0
	github.com/anatol/smart.go v0.0.0-20220615232124-371056cd18c3
	github.com/bicomsystems/go-libzfs v0.4.0
	github.com/containerd/cgroups v1.0.3
	github.com/containerd/containerd v1.6.6
	github.com/containerd/typeurl v1.0.2
	github.com/cshari-zededa/eve-tpm2-tools v0.0.4
	github.com/digitalocean/go-qemu v0.0.0-20181112162955-dd7bb9c771b8
	github.com/docker/docker v20.10.17+incompatible
	github.com/eriknordmark/ipinfo v0.0.0-20190220084921-7ee0839158f9
	github.com/fsnotify/fsnotify v1.5.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7
	github.com/google/go-containerregistry v0.8.0
	github.com/google/go-tpm v0.3.0
	github.com/google/gopacket v1.1.19
	github.com/gorilla/websocket v1.4.2
	github.com/grandcat/zeroconf v1.0.0
	github.com/jackwakefield/gopac v1.0.2
	github.com/jaypipes/ghw v0.8.0
	github.com/lf-edge/edge-containers v0.0.0-20220320131500-9d9f95d81e2c
	github.com/lf-edge/eve/api/go v0.0.0-00010101000000-000000000000
	github.com/lf-edge/eve/libs/depgraph v0.0.0-20220129022022-ba04fd269658
	github.com/lf-edge/eve/libs/reconciler v0.0.0-20220131150115-6941dbe72001
	github.com/lf-edge/eve/libs/zedUpload v0.0.0-20210120050122-276fea8f6efd
	github.com/linuxkit/linuxkit/src/cmd/linuxkit v0.0.0-20220913135124-e532e7310810
	github.com/moby/sys/mountinfo v0.6.0
	github.com/onsi/gomega v1.17.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/packetcap/go-pcap v0.0.0-20210809221331-e2e6b14e1812
	github.com/prometheus/procfs v0.7.3
	github.com/rackn/gohai v0.0.0-20190321191141-5053e7f1fa36
	github.com/satori/go.uuid v1.2.1-0.20180404165556-75cca531ea76
	github.com/shirou/gopsutil v0.0.0-20190901111213-e4ec7b275ada
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.2
	github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
	github.com/vishvananda/netlink v1.1.1-0.20210924202909-187053b97868
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/lf-edge/eve/api/go => ../../api/go

replace github.com/lf-edge/eve/libs/zedUpload => ../../libs/zedUpload

replace github.com/lf-edge/eve/libs/depgraph => ../../libs/depgraph

replace github.com/lf-edge/eve/libs/reconciler => ../../libs/reconciler
