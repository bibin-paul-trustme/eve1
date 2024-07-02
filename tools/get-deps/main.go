// Copyright(c) 2024 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/containerd/containerd/platforms"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/moby/buildkit/frontend/dockerfile/shell"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
)

// Target information
const (
	TARGETOS      = "linux"
	TARGETVARIANT = ""
)

var (
	outputImgFile  string
	outputMakeFile string
	rootfsDeps     bool
	targetArch     string
)

type printer interface {
	printSingleDep(dep string) string
	printDep(pkg string, dep string) string
	printHead() string
	printTail() string
	generate(filename string) string
}

type makefilePrinter struct {
}

type dotfilePrinter struct {
}

// Parse YML file and return all lfedge/* packages inside it
func parseYMLfile(fileName string) []string {
	var deps []string

	// Open and read the file
	bs, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Parses the file
	m, err := moby.NewConfig(bs, func(path string) (tag string, err error) {
		return path, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	deps = append(deps, m.Init...)
	for _, img := range m.Onboot {
		deps = append(deps, img.Image)
	}
	for _, img := range m.Onshutdown {
		deps = append(deps, img.Image)
	}
	for _, img := range m.Services {
		deps = append(deps, img.Image)
	}

	return deps
}

func parseDockerfile(f io.Reader) []string {
	dt, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	dockerfile, err := parser.Parse(bytes.NewReader(dt))
	if err != nil {
		panic(err)
	}
	stages, metaArgs, err := instructions.Parse(dockerfile.AST, nil)
	if err != nil {
		panic(err)
	}

	buildPlatform := []ocispecs.Platform{platforms.DefaultSpec()}[0]
	targetPlatform := ocispecs.Platform{
		Architecture: targetArch,
		OS:           TARGETOS,
		Variant:      TARGETVARIANT,
	}

	// from github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb/platform.go:getPlatformArgs
	args := map[string]string{
		"BUILDPLATFORM":  platforms.Format(buildPlatform),
		"BUILDOS":        buildPlatform.OS,
		"BUILDARCH":      buildPlatform.Architecture,
		"BUILDVARIANT":   buildPlatform.Variant,
		"TARGETPLATFORM": platforms.Format(targetPlatform),
		"TARGETOS":       targetPlatform.OS,
		"TARGETARCH":     targetPlatform.Architecture,
		"TARGETVARIANT":  targetPlatform.Variant,
	}
	for _, ma := range metaArgs {
		for _, arg := range ma.Args {
			key := arg.Key
			val := arg.ValueString()
			args[key] = val
		}
	}

	shlex := shell.NewLex(dockerfile.EscapeToken)

	targetsMap := make(map[string]struct{})
	for _, st := range stages {
		pResult, err := shlex.ProcessWordWithMatches(st.BaseName, args)
		if err != nil {
			panic(err)
		}
		if st.Platform != "" && st.Platform != targetArch {
			continue
		}
		targetsMap[pResult.Result] = struct{}{}
	}
	targets := make([]string, 0)
	for target := range targetsMap {
		targets = append(targets, target)
	}
	return targets
}

// Print a single dependency package, only suitable for dot file
func (mp makefilePrinter) printSingleDep(dep string) string {
	return ""
}

func (dp dotfilePrinter) printSingleDep(dep string) string {
	return "\"" + dep + "\";"
}

// Print a package and one of its dependency
func (mp makefilePrinter) printDep(pkg string, dep string) string {
	return pkg + ": " + dep + "\n"
}

func (dp dotfilePrinter) printDep(pkg string, dep string) string {
	return "\"" + pkg + "\" -> \"" + dep + "\";\n"
}

// Print the header and/or the initialization commands (for dot file)
func (mp makefilePrinter) printHead() string {
	return "#\n# This file was generated by EVE's build system. DO NOT EDIT.\n#\n"
}

func (dp dotfilePrinter) printHead() string {
	return "digraph unix {\n   rankdir=\"LR\";\n"
}

// Print the end of the generated file
func (mp makefilePrinter) printTail() string {
	return ""
}

func (dp dotfilePrinter) printTail() string {
	return "\n}\n"
}

// Generate the output file
func (mp makefilePrinter) generate(outfileName string) string {
	return outfileName
}

func (dp dotfilePrinter) generate(outfileName string) string {
	_, err := exec.Command("dot", "-Tpng", outfileName, "-o", outputImgFile).Output()
	if err != nil {
		fmt.Println("Failed to run dot utility.")
	}

	// Remove temporary file
	os.Remove(outfileName)

	// Set the final output file name
	outfileName = outputImgFile
	return outfileName
}

// Filter packages from the list of dependencies
func filterPkg(deps []string) []string {
	var depList []string
	dpList := make(map[string]bool)

	reLF := regexp.MustCompile("lfedge/.*")
	rePkg := regexp.MustCompile("lfedge/(?:eve-)?(.*):.*")
	for _, s := range deps {
		// We are just interested on packages from lfegde (those that we
		// published)
		if reLF.MatchString(s) {
			str := rePkg.ReplaceAllString(s, "pkg/$1")
			if !dpList[str] {
				dpList[str] = true
				depList = append(depList, str)
			}
		}
	}

	return depList
}

// Return a list of all dependencies (packages) listed in a Dockerfile
func getDeps(dockerfile string) []string {
	var depList []string
	f, err := os.Open(dockerfile)
	if err != nil {
		fmt.Println(err)
		return depList
	}
	defer f.Close()
	ss := parseDockerfile(f)

	return filterPkg(ss)
}

// Write string to file with error checking
func writeToFile(f *os.File, str string) {
	_, err := f.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mkfile := false
	imgfile := false
	pkgName := ""
	var p printer

	// Build and validate the command line
	flag.Usage = func() {
		fmt.Printf("Create dependency packages tree\n\n")
		fmt.Printf("Use:\n    %s [-r] <-i|-m> <output_file>\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&outputImgFile, "i", "", "Generate a PNG image file")
	flag.StringVar(&outputMakeFile, "m", "", "Generate a Makefile auxiliary file")
	flag.StringVar(&targetArch, "t", "amd64", "Target Architecture amd64|arm64|riscv64")
	flag.BoolVar(&rootfsDeps, "r", false, "Also generates dependencies for rootfs image")
	flag.Parse()

	if len(outputImgFile) > 0 {
		imgfile = true
		p = dotfilePrinter{}
	}
	if len(outputMakeFile) > 0 {
		mkfile = true
		p = makefilePrinter{}
	}
	if !imgfile && !mkfile {
		flag.Usage()
		os.Exit(1)
	} else if imgfile && mkfile {
		flag.Usage()
		log.Fatal("Only one type of output dependency tree can be provided.\n")
	}

	// Create the output file, if we are generating image tree, then an
	// intermediate file (.dot) must be created
	var outfile *os.File
	var errF error
	if len(outputImgFile) > 0 {
		outfile, errF = os.CreateTemp("", "eve-dot-")
		if errF != nil {
			log.Fatal(errF)
		}
	} else {
		outfile, errF = os.Create(outputMakeFile)
		if errF != nil {
			log.Fatal(errF)
		}
	}

	// Beginning of the output file
	writeToFile(outfile, p.printHead())

	// Scan all directories of pkg/
	ent, err := os.ReadDir("./pkg")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range ent {
		if !e.IsDir() {
			continue
		}
		dockerFile := filepath.Join("./pkg/", e.Name(), "/Dockerfile")
		dockerFileIn := dockerFile + ".in"

		if _, err := os.Stat(dockerFile); err != nil {
			if _, errIn := os.Stat(dockerFileIn); errIn == nil {
				// Process Dockerfile.in
				cmd := exec.Command("./tools/parse-pkgs.sh", dockerFileIn)
				out, err := cmd.Output()
				if err != nil {
					log.Println("Failed to process", dockerFileIn)
					continue
				}

				f, errTmp := os.CreateTemp("", "eve-dockerfile-")
				if errTmp != nil {
					log.Println(errTmp)
					continue
				}
				dockerFile = f.Name()
				defer os.Remove(dockerFile)
				if _, errW := f.Write(out); errW != nil {
					log.Println("Failed to write to", dockerFile)
					continue
				}
				f.Close()
			} else {
				continue
			}

			pkgName = "pkg/" + e.Name()
		} else {
			pkgName = filepath.Dir(dockerFile)
		}

		// Get package dependencies from Dockerfile
		writeToFile(outfile, p.printSingleDep(pkgName))
		depList := getDeps(dockerFile)
		for _, d := range depList {
			if d != pkgName {
				// Write a single dependency of the package
				writeToFile(outfile, p.printDep(pkgName, d))
			}
		}
	}

	// Scan rootfs dependencies
	if rootfsDeps {
		ent, err = os.ReadDir("./images/out/")
		if err == nil {
			for _, e := range ent {
				if !e.IsDir() {
					// Process yml file
					ymlFile := filepath.Join("images/out/", e.Name())
					depYML := parseYMLfile(ymlFile)
					depList := filterPkg(depYML)
					for _, d := range depList {
						writeToFile(outfile, p.printDep(ymlFile, d))
					}
				}
			}
		}
	}

	// We reach the end of the file
	writeToFile(outfile, p.printTail())
	outfileName := outfile.Name()
	outfile.Close()

	outfileName = p.generate(outfileName)

	fmt.Println("Done. Output file written to " + outfileName + ".")
}
