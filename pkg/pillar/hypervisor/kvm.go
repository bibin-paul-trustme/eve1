// Copyright (c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package hypervisor

import (
	"fmt"
	v1stat "github.com/containerd/cgroups/stats/v1"
	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/containerd"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"
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
  usb = "off"
  dump-guest-core = "off"
{{- if eq .Machine "virt" }}
  accel = "kvm:tcg"
  gic_version = "host"
{{- end -}}
{{- if ne .Machine "virt" }}
  accel = "kvm"
  vmport = "off"
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

[chardev "charlistener"]
  backend = "socket"
  path = "` + kvmStateDir + `{{.DisplayName}}/listener.qmp"
  server = "on"
  wait = "off"

[mon "listener"]
  chardev = "charlistener"
  mode = "control"

[memory]
  size = "{{.Memory}}"

[smp-opts]
  cpus = "{{.VCpus}}"
  sockets = "1"
  cores = "{{.VCpus}}"
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
const vfioDriverPath = "/sys/bus/pci/drivers/vfio-pci"

func logError(format string, a ...interface{}) error {
	log.Errorf(format, a...)
	return fmt.Errorf(format, a...)
}

//KvmContainerIntf wraps containerd methods. This helps
//in unit-testing, by mocking this interface with stubs
//to focus on kvm.go unit-testing
type KvmContainerIntf interface {
	InitContainerdClient() error
	GetMetrics(ctrID string) (*v1stat.Metrics, error)
}

//KvmContainerImpl implements KvmContainerIntf for containerd
type KvmContainerImpl struct{}

//InitContainerdClient implements InitContainerdClient interface of KvmContainerIntf
func (k *KvmContainerImpl) InitContainerdClient() error {
	return containerd.InitContainerdClient()
}

//GetMetrics implements GetMetrics interface of KvmContainerIntf
func (k *KvmContainerImpl) GetMetrics(ctrID string) (*v1stat.Metrics, error) {
	return containerd.GetMetrics(ctrID)
}

//Instantiate an object to call KvmContainerImpl methods
var kvmContainerImpl KvmContainerIntf = &KvmContainerImpl{}

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
	if err := kvmContainerImpl.InitContainerdClient(); err != nil {
		log.Errorf("InitContainerdClient failed: %v", err)
		return nil
	}
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
			dmExec:       "/usr/lib/xen/bin/qemu-system-x86_64",
			dmArgs:       []string{"-display", "none", "-S", "-no-user-config", "-nodefaults", "-no-shutdown", "-no-hpet", "-overcommit", "mem-lock=on", "-overcommit", "cpu-pm=on"},
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
	if config.VirtualizationMode == types.FML || config.VirtualizationMode == types.PV {
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
			diskContext.DiskFile = ds.FileLocation
			diskContext.Devtype = "CONTAINER"
		} else {
			diskContext.DiskFile = ds.FileLocation
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

func waitForQmp(domainName string) error {
	maxDelay := time.Second * 10
	delay := time.Second
	var waited time.Duration
	for {
		log.Infof("waitForQmp for %s: waiting for %v", domainName, delay)
		if delay != 0 {
			time.Sleep(delay)
			waited += delay
		}
		if _, err := getQemuStatus(getQmpExecutorSocket(domainName)); err == nil {
			log.Infof("waitForQmp for %s, found file", domainName)
			return nil
		} else {
			if waited > maxDelay {
				// Give up
				log.Warnf("waitForQmp for %s: giving up", domainName)
				return logError("Qmp not found")
			}
			delay = 2 * delay
			if delay > time.Minute {
				delay = time.Minute
			}
		}
	}
}

func (ctx kvmContext) Create(domainName string, cfgFilename string, config *types.DomainConfig) (int, error) {
	log.Infof("starting KVM domain %s with config %s", domainName, cfgFilename)
	if config == nil {
		return 0, logError("Empty config supplied for %s, %s", domainName, cfgFilename)
	}

	os.MkdirAll(kvmStateDir+domainName, 0777)

	pidFile := kvmStateDir + domainName + "/pid"
	//qmpFile := getQmpExecutorSocket(domainName)
	consFile := kvmStateDir + domainName + "/cons"

	dmArgs := ctx.dmArgs
	if config.VirtualizationMode == types.FML {
		dmArgs = append(dmArgs, ctx.dmFmlCPUArgs...)
	} else {
		dmArgs = append(dmArgs, ctx.dmCPUArgs...)
	}

	args := []string{ctx.dmExec}
	args = append(args, dmArgs...)
	args = append(args, "-name", domainName,
		"-readconfig", cfgFilename,
		"-pidfile", pidFile)

	domainID, err := containerd.LKTaskLaunch(domainName,
		"xen-tools", config, args)
	if err != nil {
		return 0, logError("starting LKTaskLaunch failed for %s, (%v)", domainName, err)
	}
	pid, status, err := containerd.CtrInfo(domainName)
	if err != nil {
		log.Errorf("Error getting status for container %s: %v", domainName, err)
		return 0, err
	}
	if domainID != pid || status != "running" {
		log.Errorf("domainID(%d) is not matching pid(%d), or status(%s) is not running",
			domainID, pid, status)
	}
	log.Infof("done launching qemu device model")
	if err := waitForQmp(domainName); err != nil {
		log.Errorf("Error waiting for Qmp for domain %s: %v", domainName, err)
		return 0, err
	}
	consolePty, err := getConsolePty(getQmpExecutorSocket(domainName))
	if err != nil {
		//don't bail, log an error and proceed
		log.Errorf("Error fetching charserial10 pty for domain %s: %v", domainName, err)
	} else {
		os.Symlink(consolePty, consFile)
	}

	log.Debugf("starting qmpEventHandler")
	go qmpEventHandler(getQmpListenerSocket(domainName), getQmpExecutorSocket(domainName))

	return domainID, nil
}

func (ctx kvmContext) Start(domainName string, domainID int) error {
	qmpFile := getQmpExecutorSocket(domainName)
	if err := execContinue(qmpFile); err != nil {
		return logError("failed to start domain that is stopped %v", err)
	}

	if status, err := getQemuStatus(qmpFile); err != nil || status != "running" {
		return logError("domain status is not running but %s after cont command returned %v", status, err)
	}
	ctx.domains[domainName] = domainID
	return nil
}

func (ctx kvmContext) Stop(domainName string, domainID int, force bool) error {
	if err := execShutdown(getQmpExecutorSocket(domainName)); err != nil {
		return logError("Stop: failed to execute shutdown command %v", err)
	}
	delete(ctx.domains, domainName)
	return nil
}

func (ctx kvmContext) Delete(domainName string, domainID int) error {
	//Sending a stop signal to then domain before quitting. This is done to freeze the domain before quitting it.
	execStop(getQmpExecutorSocket(domainName))
	if err := execQuit(getQmpExecutorSocket(domainName)); err != nil {
		return logError("failed to execute quit command %v", err)
	}
	// we may want to wait a little bit here and actually kill qemu process if it gets wedged
	if err := os.RemoveAll(kvmStateDir + domainName); err != nil {
		return logError("failed to clean up domain state directory %s (%v)", domainName, err)
	}
	if err := containerd.CtrDelete(domainName); err != nil {
		return logError("failed to delete container task for domain %s, %v",
			domainName, err)
	}
	return nil
}

func (ctx kvmContext) Info(domainName string, domainID int) error {
	res, err := execQueryCLIOptions(getQmpExecutorSocket(domainName))
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
	log.Infof("PCIReserve long addr is %s", long)

	overrideFile := sysfsPciDevices + long + "/driver_override"
	driverPath := sysfsPciDevices + long + "/driver"
	unbindFile := driverPath + "/unbind"

	//Check if already bound to vfio-pci
	driverPathInfo, driverPathErr := os.Stat(driverPath)
	vfioDriverPathInfo, vfioDriverPathErr := os.Stat(vfioDriverPath)
	if driverPathErr == nil && vfioDriverPathErr == nil &&
		os.SameFile(driverPathInfo, vfioDriverPathInfo) {
		log.Infof("Driver for %s is already bound to vfio-pci, skipping unbind", long)
		return nil
	}

	//map vfio-pci as the driver_override for the device
	if err := ioutil.WriteFile(overrideFile, []byte("vfio-pci"), 0644); err != nil {
		log.Fatalf("driver_override failure for PCI device %s: %v",
			long, err)
	}

	//Unbind the current driver, whatever it is, if there is one
	if _, err := os.Stat(unbindFile); err == nil {
		if err := ioutil.WriteFile(unbindFile, []byte(long), 0644); err != nil {
			log.Fatalf("unbind failure for PCI device %s: %v",
				long, err)
		}
	}

	if err := ioutil.WriteFile(sysfsPciDriversProbe, []byte(long), 0644); err != nil {
		log.Fatalf("drivers_probe failure for PCI device %s: %v",
			long, err)
	}

	return nil
}

func (ctx kvmContext) PCIRelease(long string) error {
	log.Errorf("PCIRelease long addr is %s", long)

	overrideFile := sysfsPciDevices + long + "/driver_override"
	unbindFile := sysfsPciDevices + long + "/driver/unbind"

	//Write Empty string, to clear driver_override for the device
	if err := ioutil.WriteFile(overrideFile, []byte(""), 0644); err != nil {
		log.Fatalf("driver_override failure for PCI device %s: %v",
			long, err)
	}

	//Unbind vfio-pci, if unbind file is present
	if _, err := os.Stat(unbindFile); err == nil {
		if err := ioutil.WriteFile(unbindFile, []byte(long), 0644); err != nil {
			log.Fatalf("unbind failure for PCI device %s: %v",
				long, err)
		}
	}

	//Write PCI DDDD:BB:DD.FF to /sys/bus/pci/drivers_probe,
	//as a best-effort to bring back original driver
	if err := ioutil.WriteFile(sysfsPciDriversProbe, []byte(long), 0644); err != nil {
		log.Fatalf("drivers_probe failure for PCI device %s: %v",
			long, err)
	}

	return nil
}

//IsDomainPotentiallyShuttingDown: returns false if domain's status is healthy (i.e. 'running')
func (ctx kvmContext) IsDomainPotentiallyShuttingDown(domainName string) bool {
	if status, err := getQemuStatus(getQmpExecutorSocket(domainName)); err != nil || status != "running" {
		log.Errorf("IsDomainPotentiallyShuttingDown: domain %s is not healthy. domainState: %s", domainName, status)
		return true
	}
	log.Debugf("IsDomainPotentiallyShuttingDown: domain %s is healthy", domainName)
	return false
}

func (ctx kvmContext) IsDeviceModelAlive(domid int) bool {
	_, err := os.Stat(fmt.Sprintf("/proc/%d", domid))
	return err == nil
}

func (ctx kvmContext) GetHostCPUMem() (types.HostMemory, error) {
	return selfDomCPUMem()
}

func readCPUUsage(pid int) (float64, error) {
	ps, err := process.NewProcess(int32(pid))
	if err != nil {
		return 0, err
	}

	cput, err := ps.Times()
	if err != nil {
		return 0, err
	}

	return cput.Total(), nil
}

func readMemUsage(pid int) (uint32, uint32, float64, error) {
	ps, err := process.NewProcess(int32(pid))
	if err != nil {
		return 0, 0, 0, err
	}

	processMemory, err := ps.MemoryInfo()
	if err != nil {
		return 0, 0, 0, err
	}

	usedMem := uint32(roundFromBytesToMbytes(processMemory.RSS))
	availMem := uint32(roundFromBytesToMbytes(processMemory.VMS))

	usedMemPerc, err := ps.MemoryPercent()
	if err != nil {
		return usedMem, availMem, 0, err
	}

	return usedMem, availMem, float64(usedMemPerc), nil
}

func (ctx kvmContext) GetDomsCPUMem() (map[string]types.DomainMetric, error) {
	// for more precised measurements we should be using a tool like https://github.com/cha87de/kvmtop
	res := map[string]types.DomainMetric{}

	for dom := range ctx.domains {
		var usedMem, availMem uint32
		var usedMemPerc float64
		var cpuTotal uint64

		if metric, err := kvmContainerImpl.GetMetrics(dom); err == nil {
			usedMem = uint32(roundFromBytesToMbytes(metric.Memory.Usage.Usage))
			availMem = uint32(roundFromBytesToMbytes(metric.Memory.Usage.Max))
			if availMem != 0 {
				usedMemPerc = float64(100 * float32(usedMem) / float32(availMem))
			} else {
				usedMemPerc = 0
			}
			cpuTotal = metric.CPU.Usage.Total / 1000000000
		} else {
			log.Errorf("GetDomsCPUMem failed with error %v", err)
		}

		res[dom] = types.DomainMetric{
			UUIDandVersion:    types.UUIDandVersion{},
			CPUTotal:          cpuTotal,
			UsedMemory:        usedMem,
			AvailableMemory:   availMem,
			UsedMemoryPercent: usedMemPerc,
		}
	}
	return res, nil
}

func getQmpExecutorSocket(domainName string) string {
	return kvmStateDir + domainName + "/qmp"
}

func getQmpListenerSocket(domainName string) string {
	return kvmStateDir + domainName + "/listener.qmp"
}
