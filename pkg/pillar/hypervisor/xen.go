// Copyright (c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package hypervisor

import (
	"context"
	"errors"
	"fmt"
	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/wrap"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	dom0Name = "Domain-0"
)

type typeAndPCI struct {
	pciLong string
	ioType  types.IoType
}

func addNoDuplicatePCI(list []typeAndPCI, tap typeAndPCI) []typeAndPCI {

	for _, t := range list {
		if t.pciLong == tap.pciLong {
			return list
		}
	}
	return append(list, tap)
}

func addNoDuplicate(list []string, add string) []string {

	for _, s := range list {
		if s == add {
			return list
		}
	}
	return append(list, add)
}

type xenContext struct {
}

func newXen() Hypervisor {
	return xenContext{}
}

func (ctx xenContext) Name() string {
	return "xen"
}

func (ctx xenContext) CreateDomConfig(domainName string, config types.DomainConfig, diskStatusList []types.DiskStatus,
	aa *types.AssignableAdapters, file *os.File) error {
	xen_type := "pvh"
	rootDev := ""
	extra := ""
	bootLoader := ""
	kernel := ""
	ramdisk := config.Ramdisk
	vif_type := "vif"
	xen_global := ""
	uuidStr := fmt.Sprintf("appuuid=%s ", config.UUIDandVersion.UUID)

	switch config.VirtualizationMode {
	case types.PV:
		xen_type = "pvh"
		extra = "console=hvc0 " + uuidStr + config.ExtraArgs
		kernel = "/usr/lib/xen/boot/ovmf-pvh.bin"
	case types.HVM:
		xen_type = "hvm"
		if config.Kernel != "" {
			kernel = config.Kernel
		}
	case types.FML:
		xen_type = "hvm"
		vif_type = "ioemu"
		xen_global = "hdtype = \"ahci\"\nspoof_xen = 1\npci_permissive = 1\n"
	default:
		log.Errorf("Internal error: Unknown virtualizationMode %d",
			config.VirtualizationMode)
	}

	if config.IsContainer {
		kernel = "/hostfs/boot/kernel"
		ramdisk = "/usr/lib/xen/boot/runx-initrd"
		extra = extra + " root=9p-xen dhcp=1"
	}

	file.WriteString("# This file is automatically generated by domainmgr\n")
	file.WriteString(fmt.Sprintf("name = \"%s\"\n", domainName))
	file.WriteString(fmt.Sprintf("type = \"%s\"\n", xen_type))
	file.WriteString(fmt.Sprintf("uuid = \"%s\"\n",
		config.UUIDandVersion.UUID))
	file.WriteString(xen_global)

	if kernel != "" {
		file.WriteString(fmt.Sprintf("kernel = \"%s\"\n", kernel))
	}

	if ramdisk != "" {
		file.WriteString(fmt.Sprintf("ramdisk = \"%s\"\n", ramdisk))
	}

	if bootLoader != "" {
		file.WriteString(fmt.Sprintf("bootloader = \"%s\"\n",
			bootLoader))
	}
	if config.EnableVnc {
		file.WriteString(fmt.Sprintf("vnc = 1\n"))
		file.WriteString(fmt.Sprintf("vnclisten = \"0.0.0.0\"\n"))
		file.WriteString(fmt.Sprintf("usb=1\n"))
		file.WriteString(fmt.Sprintf("usbdevice=[\"tablet\"]\n"))

		if config.VncDisplay != 0 {
			file.WriteString(fmt.Sprintf("vncdisplay = %d\n",
				config.VncDisplay))
		}
		if config.VncPasswd != "" {
			file.WriteString(fmt.Sprintf("vncpasswd = \"%s\"\n",
				config.VncPasswd))
		}
	} else {
		file.WriteString(fmt.Sprintf("vnc = 0\n"))
	}

	// Go from kbytes to mbytes
	kbyte2mbyte := func(kbyte int) int {
		return (kbyte + 1023) / 1024
	}
	file.WriteString(fmt.Sprintf("memory = %d\n",
		kbyte2mbyte(config.Memory)))
	if config.MaxMem != 0 {
		file.WriteString(fmt.Sprintf("maxmem = %d\n",
			kbyte2mbyte(config.MaxMem)))
	}
	vCpus := config.VCpus
	if vCpus == 0 {
		vCpus = 1
	}
	file.WriteString(fmt.Sprintf("vcpus = %d\n", vCpus))
	maxCpus := config.MaxCpus
	if maxCpus == 0 {
		maxCpus = vCpus
	}
	file.WriteString(fmt.Sprintf("maxcpus = %d\n", maxCpus))
	if config.CPUs != "" {
		file.WriteString(fmt.Sprintf("cpus = \"%s\"\n", config.CPUs))
	}
	if config.DeviceTree != "" {
		file.WriteString(fmt.Sprintf("device_tree = \"%s\"\n",
			config.DeviceTree))
	}
	dtString := ""
	for _, dt := range config.DtDev {
		if dtString != "" {
			dtString += ","
		}
		dtString += fmt.Sprintf("\"%s\"", dt)
	}
	if dtString != "" {
		file.WriteString(fmt.Sprintf("dtdev = [%s]\n", dtString))
	}
	// Note that qcow2 images might have partitions hence xvda1 by default
	if rootDev != "" {
		file.WriteString(fmt.Sprintf("root = \"%s\"\n", rootDev))
	}
	if extra != "" {
		file.WriteString(fmt.Sprintf("extra = \"%s\"\n", extra))
	}
	// XXX Should one be able to disable the serial console? Would need
	// knob in manifest

	var serialAssignments []string
	serialAssignments = append(serialAssignments, "pty")

	// Always prefer CDROM vdisk over disk
	file.WriteString(fmt.Sprintf("boot = \"%s\"\n", "dc"))

	var diskStrings []string
	var p9Strings []string
	for i, ds := range diskStatusList {
		if ds.Format == zconfig.Format_CONTAINER {
			p9Strings = append(p9Strings,
				fmt.Sprintf("'tag=share_dir,security_model=none,path=%s'", ds.FSVolumeLocation))
		} else {
			access := "rw"
			if ds.ReadOnly {
				access = "ro"
			}
			oneDisk := fmt.Sprintf("'%s,%s,%s,%s'",
				ds.ActiveFileLocation, strings.ToLower(ds.Format.String()), ds.Vdev, access)
			log.Debugf("Processing disk %d: %s\n", i, oneDisk)
			diskStrings = append(diskStrings, oneDisk)
		}
	}
	if len(diskStrings) > 0 {
		file.WriteString(fmt.Sprintf("disk = [%s]\n", strings.Join(diskStrings, ",")))
	}
	if len(p9Strings) > 0 {
		file.WriteString(fmt.Sprintf("p9 = [%s]\n", strings.Join(p9Strings, ",")))
	}

	vifString := ""
	for _, net := range config.VifList {
		oneVif := fmt.Sprintf("'bridge=%s,vifname=%s,mac=%s,type=%s'",
			net.Bridge, net.Vif, net.Mac, vif_type)
		if vifString == "" {
			vifString = oneVif
		} else {
			vifString = vifString + ", " + oneVif
		}
	}
	file.WriteString(fmt.Sprintf("vif = [%s]\n", vifString))

	imString := ""
	for _, im := range config.IOMem {
		if imString != "" {
			imString += ","
		}
		imString += fmt.Sprintf("\"%s\"", im)
	}
	if imString != "" {
		file.WriteString(fmt.Sprintf("iomem = [%s]\n", imString))
	}

	// Gather all PCI assignments into a single line
	// Also irqs, ioports, and serials
	// irqs and ioports are used if we are pv; serials if hvm
	var pciAssignments []typeAndPCI
	var irqAssignments []string
	var ioportsAssignments []string

	for _, irq := range config.IRQs {
		irqString := fmt.Sprintf("%d", irq)
		irqAssignments = addNoDuplicate(irqAssignments, irqString)
	}
	for _, adapter := range config.IoAdapterList {
		log.Debugf("configToXenCfg processing adapter %d %s\n",
			adapter.Type, adapter.Name)
		list := aa.LookupIoBundleAny(adapter.Name)
		// We reserved it in handleCreate so nobody could have stolen it
		if len(list) == 0 {
			log.Fatalf("configToXencfg IoBundle disappeared %d %s for %s\n",
				adapter.Type, adapter.Name, domainName)
		}
		for _, ib := range list {
			if ib == nil {
				continue
			}
			if ib.UsedByUUID != config.UUIDandVersion.UUID {
				log.Fatalf("configToXencfg IoBundle not ours %s: %d %s for %s\n",
					ib.UsedByUUID, adapter.Type, adapter.Name,
					domainName)
			}
			if ib.PciLong != "" {
				tap := typeAndPCI{pciLong: ib.PciLong, ioType: ib.Type}
				pciAssignments = addNoDuplicatePCI(pciAssignments, tap)
			}
			if ib.Irq != "" && config.VirtualizationMode == types.PV {
				log.Infof("Adding irq <%s>\n", ib.Irq)
				irqAssignments = addNoDuplicate(irqAssignments,
					ib.Irq)
			}
			if ib.Ioports != "" && config.VirtualizationMode == types.PV {
				log.Infof("Adding ioport <%s>\n", ib.Ioports)
				ioportsAssignments = addNoDuplicate(ioportsAssignments, ib.Ioports)
			}
			if ib.Serial != "" && (config.VirtualizationMode == types.HVM || config.VirtualizationMode == types.FML) {
				log.Infof("Adding serial <%s>\n", ib.Serial)
				serialAssignments = addNoDuplicate(serialAssignments, ib.Serial)
			}
		}
	}
	if len(pciAssignments) != 0 {
		log.Infof("PCI assignments %v\n", pciAssignments)
		cfg := fmt.Sprintf("pci = [ ")
		for i, pa := range pciAssignments {
			if i != 0 {
				cfg = cfg + ", "
			}
			short := types.PCILongToShort(pa.pciLong)
			// USB controller are subject to legacy USB support from
			// some BIOS. Use relaxed to get past that.
			if pa.ioType == types.IoUSB {
				cfg = cfg + fmt.Sprintf("'%s,rdm_policy=relaxed'",
					short)
			} else {
				cfg = cfg + fmt.Sprintf("'%s'", short)
			}
		}
		cfg = cfg + "]"
		log.Debugf("Adding pci config <%s>\n", cfg)
		file.WriteString(fmt.Sprintf("%s\n", cfg))
	}
	irqString := ""
	for _, irq := range irqAssignments {
		if irqString != "" {
			irqString += ","
		}
		irqString += irq
	}
	if irqString != "" {
		file.WriteString(fmt.Sprintf("irqs = [%s]\n", irqString))
	}
	ioportString := ""
	for _, ioports := range ioportsAssignments {
		if ioportString != "" {
			ioportString += ","
		}
		ioportString += ioports
	}
	if ioportString != "" {
		file.WriteString(fmt.Sprintf("ioports = [%s]\n", ioportString))
	}
	serialString := ""
	for _, serial := range serialAssignments {
		if serialString != "" {
			serialString += ","
		}
		serialString += "'" + serial + "'"
	}
	if serialString != "" {
		file.WriteString(fmt.Sprintf("serial = [%s]\n", serialString))
	}
	// XXX log file content: log.Infof("Created %s: %s
	return nil
}

func (ctx xenContext) Create(domainName string, xenCfgFilename string) (int, error) {
	log.Infof("xlCreate %s %s\n", domainName, xenCfgFilename)
	cmd := "xl"
	args := []string{
		"create",
		xenCfgFilename,
		"-p",
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl create failed ", err)
		log.Errorln("xl create output ", string(stdoutStderr))
		return 0, fmt.Errorf("xl create failed: %s\n",
			string(stdoutStderr))
	}
	log.Infof("xl create done\n")

	args = []string{
		"domid",
		domainName,
	}
	stdoutStderr, err = wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl domid failed ", err)
		log.Errorln("xl domid output ", string(stdoutStderr))
		return 0, fmt.Errorf("xl domid failed: %s\n",
			string(stdoutStderr))
	}
	res := strings.TrimSpace(string(stdoutStderr))
	domainID, err := strconv.Atoi(res)
	if err != nil {
		log.Errorf("Can't extract domainID from %s: %s\n", res, err)
		return 0, fmt.Errorf("Can't extract domainID from %s: %s\n", res, err)
	}
	return domainID, nil
}

func (ctx xenContext) Start(domainName string, domainID int) error {
	log.Infof("xlUnpause %s %d\n", domainName, domainID)
	cmd := "xl"
	args := []string{
		"unpause",
		domainName,
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl unpause failed ", err)
		log.Errorln("xl unpause output ", string(stdoutStderr))
		return fmt.Errorf("xl unpause failed: %s\n",
			string(stdoutStderr))
	}
	log.Infof("xlUnpause done. Result %s\n", string(stdoutStderr))
	return nil
}

func (ctx xenContext) Stop(domainName string, domainID int, force bool) error {
	log.Infof("xlShutdown %s %d\n", domainName, domainID)
	cmd := "xl"
	var args []string
	if force {
		args = []string{
			"shutdown",
			"-F",
			domainName,
		}
	} else {
		args = []string{
			"shutdown",
			domainName,
		}
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl shutdown failed ", err)
		log.Errorln("xl shutdown output ", string(stdoutStderr))
		return fmt.Errorf("xl shutdown failed: %s\n",
			string(stdoutStderr))
	}
	log.Infof("xl shutdown done\n")
	return nil
}

func (ctx xenContext) Delete(domainName string, domainID int) error {
	log.Infof("xlDestroy %s %d\n", domainName, domainID)
	cmd := "xl"
	args := []string{
		"destroy",
		domainName,
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl destroy failed ", err)
		log.Errorln("xl destroy output ", string(stdoutStderr))
		return fmt.Errorf("xl destroy failed: %s\n",
			string(stdoutStderr))
	}
	log.Infof("xl destroy done\n")
	return nil
}

func (ctx xenContext) Info(domainName string, domainID int) error {
	log.Infof("xlStatus %s %d\n", domainName, domainID)
	// XXX xl list -l domainName returns json. XXX but state not included!
	// Note that state is not very useful anyhow
	cmd := "xl"
	args := []string{
		"list",
		"-l",
		domainName,
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl list failed ", err)
		log.Errorln("xl list output ", string(stdoutStderr))
		return fmt.Errorf("xl list failed: %s\n",
			string(stdoutStderr))
	}
	// XXX parse json to look at state? Not currently included
	// XXX note that there is a warning at the top of the combined
	// output. If we want to parse the json we need to get Output()
	log.Infof("xl list done. Result %s\n", string(stdoutStderr))
	return nil
}

func (ctx xenContext) LookupByName(domainName string, domainID int) (int, error) {
	log.Debugf("xlDomid %s %d\n", domainName, domainID)
	cmd := "xl"
	args := []string{
		"domid",
		domainName,
	}
	// Avoid wrap since we are called periodically
	stdoutStderr, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("xl domid failed ", err)
		log.Errorln("xl domid output ", string(stdoutStderr))
		return domainID, fmt.Errorf("xl domid failed: %s\n",
			string(stdoutStderr))
	}
	res := strings.TrimSpace(string(stdoutStderr))
	domainID2, err := strconv.Atoi(res)
	if err != nil {
		log.Errorf("xl domid not integer %s: failed %s\n", res, err)
		return domainID, err
	}
	if domainID2 != domainID {
		log.Warningf("domainid changed from %d to %d for %s\n",
			domainID, domainID2, domainName)
	}
	if !isDomainRunning(domainID2) { //check if domain is shutting down.
		return domainID2, fmt.Errorf("xl domain %s not running", domainName)
	}
	return domainID2, nil
}

// Perform xenstore write to disable all of these for all VIFs
// feature-sg, feature-gso-tcpv4, feature-gso-tcpv6, feature-ipv6-csum-offload
func (ctx xenContext) Tune(domainName string, domainID int, vifCount int) error {
	log.Infof("xlDisableVifOffload %s %d %d\n",
		domainName, domainID, vifCount)
	pref := "/local/domain"
	for i := 0; i < vifCount; i += 1 {
		varNames := []string{
			fmt.Sprintf("%s/0/backend/vif/%d/%d/feature-sg",
				pref, domainID, i),
			fmt.Sprintf("%s/0/backend/vif/%d/%d/feature-gso-tcpv4",
				pref, domainID, i),
			fmt.Sprintf("%s/0/backend/vif/%d/%d/feature-gso-tcpv6",
				pref, domainID, i),
			fmt.Sprintf("%s/0/backend/vif/%d/%d/feature-ipv4-csum-offload",
				pref, domainID, i),
			fmt.Sprintf("%s/0/backend/vif/%d/%d/feature-ipv6-csum-offload",
				pref, domainID, i),
			fmt.Sprintf("%s/%d/device/vif/%d/feature-sg",
				pref, domainID, i),
			fmt.Sprintf("%s/%d/device/vif/%d/feature-gso-tcpv4",
				pref, domainID, i),
			fmt.Sprintf("%s/%d/device/vif/%d/feature-gso-tcpv6",
				pref, domainID, i),
			fmt.Sprintf("%s/%d/device/vif/%d/feature-ipv4-csum-offload",
				pref, domainID, i),
			fmt.Sprintf("%s/%d/device/vif/%d/feature-ipv6-csum-offload",
				pref, domainID, i),
		}
		for _, varName := range varNames {
			cmd := "xenstore"
			args := []string{
				"write",
				varName,
				"0",
			}
			stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
			if err != nil {
				log.Errorln("xenstore write failed ", err)
				log.Errorln("xenstore write output ", string(stdoutStderr))
				return fmt.Errorf("xenstore write failed: %s\n",
					string(stdoutStderr))
			}
			log.Debugf("xenstore write done. Result %s\n",
				string(stdoutStderr))
		}
	}

	log.Infof("xlDisableVifOffload done.\n")
	return nil
}

func (ctx xenContext) PCIReserve(long string) error {
	log.Infof("pciAssignableAdd %s\n", long)
	cmd := "xl"
	args := []string{
		"pci-assignable-add",
		long,
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("xl pci-assignable-add failed: %s\n",
			string(stdoutStderr))
		log.Errorln(errStr)
		return errors.New(errStr)
	}
	log.Infof("xl pci-assignable-add done\n")
	return nil
}

func (ctx xenContext) PCIRelease(long string) error {
	log.Infof("pciAssignableRemove %s\n", long)
	cmd := "xl"
	args := []string{
		"pci-assignable-rem",
		"-r",
		long,
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		errStr := fmt.Sprintf("xl pci-assignable-rem failed: %s\n",
			string(stdoutStderr))
		log.Errorln(errStr)
		return errors.New(errStr)
	}
	log.Infof("xl pci-assignable-rem done\n")
	return nil
}

func (ctx xenContext) IsDeviceModelAlive(domid int) bool {
	// create pgrep command to see if dataplane is running
	match := fmt.Sprintf("domid %d", domid)
	cmd := wrap.Command("pgrep", "-f", match)

	// pgrep returns 0 when there is atleast one matching program running
	// cmd.Output returns nil when pgrep returns 0, otherwise pids.
	out, err := cmd.Output()

	if err != nil {
		log.Infof("IsDeviceModelAlive: %s process is not running: %s",
			match, err)
		return false
	}
	log.Infof("IsDeviceModelAlive: Instances of %s is running: %s",
		match, out)
	return true
}

func (ctx xenContext) GetHostCPUMem() (types.HostMemory, error) {
	xlCmd := exec.Command("xl", "info")
	stdout, err := xlCmd.Output()
	if err != nil {
		log.Errorf("xl info failed %s\n falling back on Dom0 stats", err)
		return selfDomCPUMem()
	}

	xlInfo := string(stdout)
	splitXlInfo := strings.Split(xlInfo, "\n")

	dict := make(map[string]string, len(splitXlInfo)-1)
	for _, str := range splitXlInfo {
		res := strings.SplitN(str, ":", 2)
		if len(res) == 2 {
			dict[strings.TrimSpace(res[0])] = strings.TrimSpace(res[1])
		}
	}

	hm := types.HostMemory{}
	hm.TotalMemoryMB, err = strconv.ParseUint(dict["total_memory"], 10, 64)
	if err != nil {
		log.Errorf("Failed parsing total_memory: %s", err)
		hm.TotalMemoryMB = 0
	}
	hm.FreeMemoryMB, err = strconv.ParseUint(dict["free_memory"], 10, 64)
	if err != nil {
		log.Errorf("Failed parsing free_memory: %s", err)
		hm.FreeMemoryMB = 0
	}

	// Note that this is the set of physical CPUs which is different
	// than the set of CPUs assigned to dom0
	var ncpus uint64
	ncpus, err = strconv.ParseUint(dict["nr_cpus"], 10, 32)
	if err != nil {
		log.Errorln("error while converting ncpus to int: ", err)
		ncpus = 0
	}
	hm.Ncpus = uint32(ncpus)
	if false {
		// debug code to compare Xen and fallback
		// XXX remove debug code
		hm2, err := selfDomCPUMem()
		log.Infof("XXX xen %+v fallback %+v (%v)", hm, hm2, err)
	}
	return hm, nil
}

func (ctx xenContext) GetDomsCPUMem() (map[string]types.DomainMetric, error) {
	count := 0
	counter := 0
	stdout, _, _ := execWithTimeout("xentop", "-b", "-d", "1", "-i", "2", "-f")
	xentopInfo := string(stdout)

	splitXentopInfo := strings.Split(xentopInfo, "\n")
	splitXentopInfoLength := len(splitXentopInfo)
	var i int
	var start int

	for i = 0; i < splitXentopInfoLength; i++ {

		str := splitXentopInfo[i]
		re := regexp.MustCompile(" ")

		spaceRemovedsplitXentopInfo := re.ReplaceAllLiteralString(str, "")
		matched, err := regexp.MatchString("NAMESTATECPU.*", spaceRemovedsplitXentopInfo)

		if err != nil {
			log.Debugf("MatchString failed: %s", err)
		} else if matched {

			count++
			log.Debugf("string matched: %s", str)
			if count == 2 {
				start = i + 1
				log.Debugf("value of i: %d", start)
			}
		}
	}

	length := splitXentopInfoLength - 1 - start
	finalOutput := make([][]string, length)

	for j := start; j < splitXentopInfoLength-1; j++ {
		finalOutput[j-start] = strings.Fields(strings.TrimSpace(splitXentopInfo[j]))
	}

	cpuMemoryStat := make([][]string, length)

	for i := range cpuMemoryStat {
		cpuMemoryStat[i] = make([]string, 20)
	}

	// Need to treat "no limit" as one token
	for f := 0; f < length; f++ {

		// First name and state
		out := 0
		counter++
		cpuMemoryStat[f][counter] = finalOutput[f][out]
		out++
		counter++
		cpuMemoryStat[f][counter] = finalOutput[f][out]
		out++
		for ; out < len(finalOutput[f]); out++ {

			if finalOutput[f][out] == "no" {

			} else if finalOutput[f][out] == "limit" {
				counter++
				cpuMemoryStat[f][counter] = "no limit"
			} else {
				counter++
				cpuMemoryStat[f][counter] = finalOutput[f][out]
			}
		}
		counter = 0
	}
	log.Debugf("ExecuteXentopCmd return %+v", cpuMemoryStat)

	var dmList map[string]types.DomainMetric
	if len(cpuMemoryStat) == 0 {
		dmList = fallbackDomainMetric()
	} else {
		dmList = parseCPUMemoryStat(cpuMemoryStat)
	}
	// finally add host entry to dmList
	if false {
		// debug code to compare Xen and fallback
		log.Infof("XXX reported DomainMetric %+v", dmList)
		dmList = fallbackDomainMetric()
		log.Infof("XXX fallback DomainMetric %+v", dmList)
	}
	return dmList, nil
}

// Returns cpuTotal, usedMemory, availableMemory, usedPercentage
func parseCPUMemoryStat(cpuMemoryStat [][]string) map[string]types.DomainMetric {

	result := make(map[string]types.DomainMetric)
	for _, stat := range cpuMemoryStat {
		if len(stat) <= 2 {
			continue
		}
		domainname := strings.TrimSpace(stat[1])
		if len(stat) <= 6 {
			continue
		}
		log.Debugf("lookupCPUMemoryStat for %s %d elem: %+v",
			domainname, len(stat), stat)
		cpuTotal, err := strconv.ParseUint(stat[3], 10, 0)
		if err != nil {
			log.Errorf("ParseUint(%s) failed: %s",
				stat[3], err)
			cpuTotal = 0
		}
		// This is in kbytes
		totalMemory, err := strconv.ParseUint(stat[5], 10, 0)
		if err != nil {
			log.Errorf("ParseUint(%s) failed: %s",
				stat[5], err)
			totalMemory = 0
		}
		totalMemory = roundFromKbytesToMbytes(totalMemory)
		usedMemoryPercent, err := strconv.ParseFloat(stat[6], 10)
		if err != nil {
			log.Errorf("ParseFloat(%s) failed: %s",
				stat[6], err)
			usedMemoryPercent = 0
		}
		usedMemory := (float64(totalMemory) * (usedMemoryPercent)) / 100
		availableMemory := float64(totalMemory) - usedMemory

		dm := types.DomainMetric{
			CPUTotal:          cpuTotal,
			UsedMemory:        uint32(usedMemory),
			AvailableMemory:   uint32(availableMemory),
			UsedMemoryPercent: float64(usedMemoryPercent),
		}
		result[domainname] = dm
	}
	return result
}

// First approximation for a host without Xen
// XXX Assumes that all of the used memory in the host is overhead the same way dom0 is
// overhead, which is completely incorrect when running containers
func fallbackDomainMetric() map[string]types.DomainMetric {
	dmList := make(map[string]types.DomainMetric)
	vm, err := mem.VirtualMemory()
	if err != nil {
		return dmList
	}
	var usedMemoryPercent float64
	if vm.Total != 0 {
		usedMemoryPercent = float64(100 * (vm.Total - vm.Available) / vm.Total)
	}
	total := roundFromBytesToMbytes(vm.Total)
	available := roundFromBytesToMbytes(vm.Available)
	dm := types.DomainMetric{
		UsedMemory:        uint32(total - available),
		AvailableMemory:   uint32(available),
		UsedMemoryPercent: usedMemoryPercent,
	}
	// Ask for one total entry
	cpuStat, err := cpu.Times(false)
	if err != nil {
		return dmList
	}
	for _, cpu := range cpuStat {
		dm.CPUTotal = uint64(cpu.Total())
		break
	}
	dmList[dom0Name] = dm
	return dmList
}

func execWithTimeout(command string, args ...string) ([]byte, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)
	out, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, false, nil
	}
	return out, true, err
}

//verifyDomainState: Verifies if a domain is up and running by its state. Domain can have the following states:
//r - currently running
//b - blocked, and not running or runnable
//p - paused
//s - a shutdown command has been sent, but the domain isn't dying yet
//c - the domain has crashed
//d - the domain is dying, but hasn't properly shut down or crashed
func isDomainRunning(domainID int) bool {
	cmd := "xl"
	args := []string{
		"list",
		strconv.Itoa(domainID),
	}
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		//domain is not present
		log.Errorln("verifyDomainState: xl list failed ", err)
		log.Errorln("verifyDomainState: xl list output ", string(stdoutStderr))
		return false
	}
	//Removing all extra space between column result and split the result as array.
	xlDomainResult := regexp.MustCompile(`\s+`).ReplaceAllString(strings.Split(string(stdoutStderr), "\n")[1], " ")
	//Domain's status is 5th column in xl list <domain> result
	domainState := strings.Split(xlDomainResult, " ")[4]
	//Sometimes the domain can be in multiple states (logically possible) at the same time.
	// In such cases we can consider the last state. For example when a domain is shutting down its state will be
	//"--psc-", we can consider "c" to represent that domain is in middle of shutting down.
	var lastState string
	for _, state := range domainState {
		if state != '-' {
			lastState = string(state)
		}
	}
	log.Debugf("verifyDomainState: domain: %d domainState: %s.", domainID, domainState)
	log.Debugf("verifyDomainState: domain: %d lastState: %s.", domainID, lastState)
	if lastState == "r" || lastState == "b" {
		//domain is up and running.
		return true
	}
	return false
}
