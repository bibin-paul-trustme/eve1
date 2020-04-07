// Copyright (c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package hypervisor

import (
	"fmt"
	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/wrap"
	"github.com/prometheus/procfs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

// We build device model around PCIe topology according to best practices
//    https://github.com/qemu/qemu/blob/master/docs/pcie.txt
// and
//    https://libvirt.org/pci-hotplug.html
// Thus the only PCI devices plugged directly into the root (pci.0) bus are:
//    00:01.0 cirrus-vga
//    00:02.0 pcie-root-port for QEMU XHCI Host Controller
//    00:0x.0 pcie-root-port for block or network device #x (where x > 2)
//    00:0y.0 virtio-9p-pci
//
// This makes everything but 9P volumes be separated from root pci bus
// and effectively hang off the bus of its own:
//     01:00.0 QEMU XHCI Host Controller (behind pcie-root-port 00:02.0)
//     xx:00.0 block or network device #x (behind pcie-root-port 00:0x.0)
//
// It would be nice to figure out how to do the same with virtio-9p-pci
// eventually, but for now this is not a high priority.

const qemuConfTemplate = `# This file is automatically generated by domainmgr
[msg]
  timestamp = "on"

[machine]
  type = "{{.Machine}}"
  accel = "kvm"
  usb = "off"
  dump-guest-core = "off"
{{- if ne .Machine "virt" }}
  vmport = "off"
{{- end -}}
{{- if ne .Machine "virt" }}
  kernel-irqchip = "on"
{{- end -}}
{{- if .BootLoader }}
  firmware = "{{.BootLoader}}"
{{- end -}}
{{- if .Kernel }}
  kernel = "{{.Kernel}}"
{{- end -}}
{{- if .Ramdisk }}
  initrd = "{{.Ramdisk}}"
{{- end -}}
{{- if .DeviceTree }}
  dtb = "{{.DeviceTree}}"
{{- end -}}
{{- if .ExtraArgs }}
  append = "{{.ExtraArgs}}"
{{ end }}
{{if ne .Machine "virt" }}
[global]
  driver = "kvm-pit"
  property = "lost_tick_policy"
  value = "delay"

[global]
  driver = "ICH9-LPC"
  property = "disable_s3"
  value = "1"

[global]
  driver = "ICH9-LPC"
  property = "disable_s4"
  value = "1"

[rtc]
  base = "localtime"
  driftfix = "slew"

[device]
  driver = "intel-iommu"
{{ end }}
[realtime]
  mlock = "off"

[chardev "charmonitor"]
  backend = "socket"
  path = "` + kvmStateDir + `{{.DisplayName}}/qmp"
  server = "on"
  wait = "off"

[mon "monitor"]
  chardev = "charmonitor"
  mode = "control"

[memory]
  size = "{{.Memory}}"

[smp-opts]
  cpus = "{{.VCpus}}"
  sockets = "{{.VCpus}}"
  cores = "1"
  threads = "1"
{{if ne .Machine "virt" }}
[chardev "charserial0"]
  backend = "pty"

[device "serial0"]
  driver = "isa-serial"
  chardev = "charserial0"
{{ end }}

{{if .EnableVnc}}
[vnc "default"]
  vnc = "0.0.0.0:{{if .VncDisplay}}{{.VncDisplay}}{{else}}0{{end}}"
  to = "99"
{{- if .VncPasswd}}
  password = "on"
{{- end -}}
{{end}}
#[device "video0"]
#  driver = "qxl-vga"
#  ram_size = "67108864"
#  vram_size = "67108864"
#  vram64_size_mb = "0"
#  vgamem_mb = "16"
#  max_outputs = "1"
#  bus = "pcie.0"
#  addr = "0x1"
[device "video0"]
  driver = "cirrus-vga"
  vgamem_mb = "16"
  bus = "pcie.0"
  addr = "0x1"

[device "pci.2"]
  driver = "pcie-root-port"
  port = "12"
  chassis = "2"
  bus = "pcie.0"
  addr = "0x2"

[device "usb"]
  driver = "qemu-xhci"
  p2 = "15"
  p3 = "15"
  bus = "pci.2"
  addr = "0x0"

[device "input0"]
  driver = "usb-tablet"
  bus = "usb.0"
  port = "1"
`

const qemuDiskTemplate = `
{{if eq .Devtype "CDROM"}}
[drive "drive-sata0-{{.DiskID}}"]
  file = "{{.DiskFile}}"
  format = "{{.DiskFormat}}"
  if = "none"
  media = "cdrom"
  readonly = "on"

[device "sata0-{{.SATAId}}"]
  driver = "ide-cd"
  bus = "ide.{{.SATAId}}"
  drive = "drive-sata0-{{.DiskID}}"
{{else if eq .Devtype "CONTAINER"}}
[fsdev "fsdev{{.DiskID}}"]
  fsdriver = "local"
  security_model = "none"
  path = "{{.DiskFile}}"

[device "fs{{.DiskID}}"]
  driver = "virtio-9p-pci"
  fsdev = "fsdev{{.DiskID}}"
  mount_tag = "hostshare"
  addr = "{{.PCIId}}"
{{else}}
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "1{{.PCIId}}"
  chassis = "{{.PCIId}}"
  bus = "pcie.0"
  addr = "{{.PCIId}}"

[drive "drive-virtio-disk{{.DiskID}}"]
  file = "{{.DiskFile}}"
  format = "{{.DiskFormat}}"
  if = "none"

[device "virtio-disk{{.DiskID}}"]
  driver = "virtio-blk-pci"
  scsi = "off"
  bus = "pci.{{.PCIId}}"
  addr = "0x0"
  drive = "drive-virtio-disk{{.DiskID}}"
  bootindex = "{{.DiskID}}"
{{end}}`

const qemuNetTemplate = `
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "1{{.PCIId}}"
  chassis = "{{.PCIId}}"
  bus = "pcie.0"
  multifunction = "on"
  addr = "{{.PCIId}}"

[netdev "hostnet{{.NetID}}"]
  type = "tap"
  ifname = "{{.Vif}}"
  br = "{{.Bridge}}"
  script = "/etc/xen/scripts/qemu-ifup"
  downscript = "no"

[device "net{{.NetID}}"]
  driver = "virtio-net-pci"
  netdev = "hostnet{{.NetID}}"
  mac = "{{.Mac}}"
  bus = "pci.{{.PCIId}}"
  addr = "0x0"
`

const qemuPciPassthruTemplate = `
[device]
  driver = "vfio-pci"
  host = "{{.PciShortAddr}}"
{{- if .Xvga }}
  x-vga = "on"
{{- end -}}
`
const qemuSerialTemplate = `
[chardev "charserial-usr{{.ID}}"]
  backend = "tty"
  path = "{{.SerialPortName}}"

[device "serial-usr{{.ID}}"]
  driver = "isa-serial"
  chardev = "charserial-usr{{.ID}}"
`

const kvmStateDir = "/var/run/hypervisor/kvm/"
const sysfsPciDevices = "/sys/bus/pci/devices/"
const sysfsVfioPciBind = "/sys/bus/pci/drivers/vfio-pci/bind"
const sysfsPciDriversProbe = "/sys/bus/pci/drivers_probe"

func logError(format string, a ...interface{}) error {
	log.Errorf(format, a...)
	return fmt.Errorf(format, a...)
}

// KVM domains map 1-1 to anchor device model UNIX processes (qemu or firecracker)
// For every anchor process we maintain the following entry points in the
// /var/run/hypervisor/kvm/DOMAIN_NAME:
//    pid - contains PID of the anchor process
//    qmp - UNIX domain socket that allows us to talk to anchor process
//   cons - symlink to /dev/pts/X that allows us to talk to the serial console of the domain
// In addition to that, we also maintain DOMAIN_NAME -> PID mapping in kvmContext, so we don't
// have to look things up in the filesystem all the time (this also allows us to filter domains
// that may be created by others)
type kvmContext struct {
	domains map[string]int
	// for now the following is statically configured and can not be changed per domain
	devicemodel  string
	dmExec       string
	dmArgs       []string
	dmCPUArgs    []string
	dmFmlCPUArgs []string
}

func newKvm() Hypervisor {
	// later on we may want to pass device model machine type in DomainConfig directly;
	// for now -- lets just pick a static device model based on the host architecture
	// "-cpu host",
	// -cpu IvyBridge-IBRS,ss=on,vmx=on,movbe=on,hypervisor=on,arat=on,tsc_adjust=on,mpx=on,rdseed=on,smap=on,clflushopt=on,sha-ni=on,umip=on,md-clear=on,arch-capabilities=on,xsaveopt=on,xsavec=on,xgetbv1=on,xsaves=on,pdpe1gb=on,3dnowprefetch=on,avx=off,f16c=off,hv_time,hv_relaxed,hv_vapic,hv_spinlocks=0x1fff
	switch runtime.GOARCH {
	case "arm64":
		return kvmContext{
			domains:      map[string]int{},
			devicemodel:  "virt",
			dmExec:       "qemu-system-aarch64",
			dmArgs:       []string{"-display", "none", "-daemonize", "-S", "-no-user-config", "-nodefaults", "-no-shutdown", "-serial", "pty"},
			dmCPUArgs:    []string{"-cpu", "host"},
			dmFmlCPUArgs: []string{},
		}
	case "amd64":
		return kvmContext{
			domains:      map[string]int{},
			devicemodel:  "pc-q35-3.1",
			dmExec:       "qemu-system-x86_64",
			dmArgs:       []string{"-display", "none", "-daemonize", "-S", "-no-user-config", "-nodefaults", "-no-shutdown", "-no-hpet"},
			dmCPUArgs:    []string{},
			dmFmlCPUArgs: []string{"-cpu", "host,hv_time,hv_relaxed,hv_vendor_id=eveitis,hypervisor=off,kvm=off"},
		}
	}
	return nil
}

func (ctx kvmContext) Name() string {
	return "kvm"
}

func (ctx kvmContext) CreateDomConfig(domainName string, config types.DomainConfig, diskStatusList []types.DiskStatus,
	aa *types.AssignableAdapters, file *os.File) error {
	tmplCtx := struct {
		Machine string
		types.DomainConfig
	}{ctx.devicemodel, config}
	tmplCtx.Memory = (config.Memory + 1023) / 1024
	tmplCtx.DisplayName = domainName
	if config.VirtualizationMode == types.FML {
		tmplCtx.BootLoader = "/usr/lib/xen/boot/ovmf.bin"
	} else {
		tmplCtx.BootLoader = ""
	}
	if config.IsContainer {
		tmplCtx.Kernel = "/hostfs/boot/kernel"
		tmplCtx.Ramdisk = "/usr/lib/xen/boot/runx-initrd"
		tmplCtx.ExtraArgs = config.ExtraArgs + " root=9p-kvm dhcp=1"
	}

	// render global device model settings
	t, _ := template.New("qemu").Parse(qemuConfTemplate)
	if err := t.Execute(file, tmplCtx); err != nil {
		return logError("can't write to config file %s (%v)", file.Name(), err)
	}

	// render disk device model settings
	diskContext := struct {
		PCIId, DiskID, SATAId         int
		DiskFile, DiskFormat, Devtype string
		ro                            bool
	}{PCIId: 3, DiskID: 0, SATAId: 0}
	t, _ = template.New("qemuDisk").Parse(qemuDiskTemplate)
	for _, ds := range diskStatusList {
		if ds.Format == zconfig.Format_CONTAINER {
			diskContext.DiskFile = ds.FSVolumeLocation
			diskContext.Devtype = "CONTAINER"
		} else {
			diskContext.DiskFile = ds.ActiveFileLocation
			diskContext.Devtype = strings.ToUpper(ds.Devtype)
		}
		diskContext.ro = ds.ReadOnly
		diskContext.DiskFormat = strings.ToLower(ds.Format.String())
		if err := t.Execute(file, diskContext); err != nil {
			return logError("can't write to config file %s (%v)", file.Name(), err)
		}
		if diskContext.Devtype == "CDROM" {
			diskContext.SATAId = diskContext.SATAId + 1
		} else {
			diskContext.PCIId = diskContext.PCIId + 1
		}
		diskContext.DiskID = diskContext.DiskID + 1
	}

	// render network device model settings
	netContext := struct {
		PCIId, NetID     int
		Mac, Bridge, Vif string
	}{PCIId: diskContext.PCIId, NetID: 0}
	t, _ = template.New("qemuNet").Parse(qemuNetTemplate)
	for _, net := range config.VifList {
		netContext.Mac = net.Mac
		netContext.Bridge = net.Bridge
		netContext.Vif = net.Vif
		if err := t.Execute(file, netContext); err != nil {
			return logError("can't write to config file %s (%v)", file.Name(), err)
		}
		netContext.PCIId = netContext.PCIId + 1
		netContext.NetID = netContext.NetID + 1
	}

	// Gather all PCI assignments into a single line
	var pciAssignments []typeAndPCI
	// Gather all serial assignments into a single line
	var serialAssignments []string

	for _, adapter := range config.IoAdapterList {
		log.Debugf("processing adapter %d %s\n", adapter.Type, adapter.Name)
		list := aa.LookupIoBundleAny(adapter.Name)
		// We reserved it in handleCreate so nobody could have stolen it
		if len(list) == 0 {
			log.Fatalf("IoBundle disappeared %d %s for %s\n",
				adapter.Type, adapter.Name, domainName)
		}
		for _, ib := range list {
			if ib == nil {
				continue
			}
			if ib.UsedByUUID != config.UUIDandVersion.UUID {
				log.Fatalf("IoBundle not ours %s: %d %s for %s\n",
					ib.UsedByUUID, adapter.Type, adapter.Name,
					domainName)
			}
			if ib.PciLong != "" {
				log.Infof("Adding PCI device <%v>\n", ib.PciLong)
				tap := typeAndPCI{pciLong: ib.PciLong, ioType: ib.Type}
				pciAssignments = addNoDuplicatePCI(pciAssignments, tap)
			}
			if ib.Serial != "" {
				log.Infof("Adding serial <%s>\n", ib.Serial)
				serialAssignments = addNoDuplicate(serialAssignments, ib.Serial)
			}
		}
	}
	if len(pciAssignments) != 0 {
		pciPTContext := struct {
			PciShortAddr string
			Xvga         bool
		}{PciShortAddr: "", Xvga: false}

		t, _ = template.New("qemuPciPT").Parse(qemuPciPassthruTemplate)
		for _, pa := range pciAssignments {
			short := types.PCILongToShort(pa.pciLong)
			bootVgaFile := sysfsPciDevices + pa.pciLong + "/boot_vga"
			if _, err := os.Stat(bootVgaFile); err == nil {
				pciPTContext.Xvga = true
			}

			pciPTContext.PciShortAddr = short
			if err := t.Execute(file, pciPTContext); err != nil {
				return logError("can't write PCI Passthrough to config file %s (%v)", file.Name(), err)
			}
			pciPTContext.Xvga = false
		}
		serialPortContext := struct {
			SerialPortName string
			ID             int
		}{SerialPortName: "", ID: 0}

		t, _ = template.New("qemuSerial").Parse(qemuSerialTemplate)
		for id, serial := range serialAssignments {
			serialPortContext.SerialPortName = serial
			fmt.Printf("id for serial is %d\n", id)
			serialPortContext.ID = id
			if err := t.Execute(file, serialPortContext); err != nil {
				return logError("can't write serial assignment to config file %s (%v)", file.Name(), err)
			}
		}
	}
	return nil
}

func (ctx kvmContext) Create(domainName string, cfgFilename string, VirtualizationMode types.VmMode) (int, error) {
	log.Infof("starting KVM domain %s with config %s\n", domainName, cfgFilename)

	os.MkdirAll(kvmStateDir+domainName, 0777)

	pidFile := kvmStateDir + domainName + "/pid"
	qmpFile := kvmStateDir + domainName + "/qmp"
	consFile := kvmStateDir + domainName + "/cons"

	dmArgs := ctx.dmArgs
	if VirtualizationMode == types.FML {
		dmArgs = append(dmArgs, ctx.dmFmlCPUArgs...)
	} else {
		dmArgs = append(dmArgs, ctx.dmCPUArgs...)
	}
	stdoutStderr, err := wrap.Command(ctx.dmExec, append(dmArgs,
		"-name", domainName, "-readconfig", cfgFilename, "-pidfile", pidFile)...).CombinedOutput()
	if err != nil {
		return 0, logError("starting qemu failed %s (%v)", string(stdoutStderr), err)
	}
	log.Infof("done launching qemu device model")

	// lets capture console device; don't mind if it fails
	match := regexp.MustCompile(`char device redirected to ([^ ]*) \(label charserial0\)`).FindStringSubmatch(string(stdoutStderr))
	if len(match) == 2 {
		os.Symlink(match[1], consFile)
	}
	pidStr, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return 0, logError("failed to retrieve qemu PID file %s (%v)", pidFile, err)
	}

	domainID, err := strconv.Atoi(strings.TrimSpace(string(pidStr)))
	if err != nil {
		return 0, logError("Can't extract domainID from %s (%v)\n", string(pidStr), err)
	}

	if status, err := getQemuStatus(qmpFile); err != nil || status != "prelaunch" {
		log.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
		return 0, fmt.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
	}

	ctx.domains[domainName] = domainID
	return domainID, nil
}

func (ctx kvmContext) Start(domainName string, domainID int) error {
	qmpFile := kvmStateDir + domainName + "/qmp"

	if err := execContinue(qmpFile); err != nil {
		return logError("failed to start domain that is stopped %v", err)
	}

	if status, err := getQemuStatus(qmpFile); err != nil || status != "running" {
		return logError("domain status is not running but %s after cont command returned %v", status, err)
	}

	return nil
}

func (ctx kvmContext) Stop(domainName string, domainID int, force bool) error {
	if err := execShutdown(kvmStateDir + domainName + "/qmp"); err != nil {
		return logError("failed to execute shutdown command %v", err)
	}
	return nil
}

func (ctx kvmContext) Delete(domainName string, domainID int) error {
	if err := execQuit(kvmStateDir + domainName + "/qmp"); err != nil {
		return logError("failed to execute quite command %v", err)
	}
	// we may want to wait a little bit here and actually kill qemu process if it gets wedged
	if err := os.RemoveAll(kvmStateDir + domainName); err != nil {
		return logError("failed to clean up domain state directory %s (%v)", domainName, err)
	}
	return nil
}

func (ctx kvmContext) Info(domainName string, domainID int) error {
	res, err := execQueryCLIOptions(kvmStateDir + domainName + "/qmp")
	log.Infof("KVM Info for domain %s %d %s (%v)", domainName, domainID, res, err)
	return err
}

func (ctx kvmContext) LookupByName(domainName string, domainID int) (int, error) {
	if id, found := ctx.domains[domainName]; !found || id != domainID {
		return id, logError("couldn't find domain %s or new id %d != old one %d", domainName, id, domainID)
	}
	return domainID, nil
}

func (ctx kvmContext) Tune(domainName string, domainID int, vifCount int) error {
	return nil
}

func (ctx kvmContext) PCIReserve(long string) error {
	log.Errorf("PCIReserve long addr is %s", long)

	overrideFile := sysfsPciDevices + long + "/driver_override"
	unbindFile := sysfsPciDevices + long + "/driver/unbind"

	//Unbind the current driver, whatever it is, if there is one
	if _, err := os.Stat(unbindFile); err == nil {
		if err := ioutil.WriteFile(unbindFile, []byte(long), 0644); err != nil {
			log.Fatalf("unbind failure for PCI device %s: %v",
				long, err)
		}
	}

	//map vfio-pci as the driver_override for the device
	if err := ioutil.WriteFile(overrideFile, []byte("vfio-pci"), 0644); err != nil {
		log.Fatalf("driver_override failure for PCI device %s: %v",
			long, err)
	}

	//Write PCI DDDD:BB:DD.FF to /sys/bus/pci/drivers/vfio-pci/bind
	if err := ioutil.WriteFile(sysfsVfioPciBind, []byte(long), 0644); err != nil {
		log.Fatalf("bind failure for PCI device %s: %v",
			long, err)
	}

	return nil
}

func (ctx kvmContext) PCIRelease(long string) error {
	log.Errorf("PCIRelease long addr is %s", long)

	overrideFile := sysfsPciDevices + long + "/driver_override"
	unbindFile := sysfsPciDevices + long + "/driver/unbind"

	//Unbind vfio-pci, if unbind file is present
	if _, err := os.Stat(unbindFile); err == nil {
		if err := ioutil.WriteFile(unbindFile, []byte(long), 0644); err != nil {
			log.Fatalf("unbind failure for PCI device %s: %v",
				long, err)
		}
	}

	//Write Empty string, to clear driver_override for the device
	if err := ioutil.WriteFile(overrideFile, []byte(""), 0644); err != nil {
		log.Fatalf("driver_override failure for PCI device %s: %v",
			long, err)
	}

	//Write PCI DDDD:BB:DD.FF to /sys/bus/pci/drivers_probe,
	//as a best-effort to bring back original driver
	if err := ioutil.WriteFile(sysfsPciDriversProbe, []byte(long), 0644); err != nil {
		log.Fatalf("drivers_probe failure for PCI device %s: %v",
			long, err)
	}

	return nil
}

func (ctx kvmContext) IsDeviceModelAlive(domid int) bool {
	_, err := os.Stat(fmt.Sprintf("/proc/%d", domid))
	return err == nil
}

func (ctx kvmContext) GetHostCPUMem() (types.HostMemory, error) {
	return selfDomCPUMem()
}

func (ctx kvmContext) GetDomsCPUMem() (map[string]types.DomainMetric, error) {
	// for more precised measurements we should be using https://github.com/cha87de/kvmtop
	res := map[string]types.DomainMetric{}

	proc, err := procfs.NewFS("/proc")
	if err != nil {
		return res, logError("can't access /procfs %v", err)
	}

	for dom, pid := range ctx.domains {
		p, err := proc.Proc(pid)
		if err != nil {
			return res, logError("can't access stats for domain %s PID %d (%v)", dom, pid, err)
		}

		s, err := p.Stat()
		if err != nil {
			return res, logError("can't access stats for domain %s PID %d (%v)", dom, pid, err)
		}
		res[dom] = types.DomainMetric{
			UUIDandVersion:    types.UUIDandVersion{},
			CPUTotal:          0,
			UsedMemory:        uint32(roundFromBytesToMbytes(uint64(s.ResidentMemory()))),
			AvailableMemory:   uint32(roundFromBytesToMbytes(uint64(s.VirtualMemory()))),
			UsedMemoryPercent: float64((float32(s.ResidentMemory()) / float32(s.VirtualMemory())) * 100),
		}
	}
	return res, nil
}
