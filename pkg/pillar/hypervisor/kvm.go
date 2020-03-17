// Copyright (c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package hypervisor

import (
	"fmt"
	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/wrap"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

const qemuConfTemplate = `# This file is automatically generated by domainmgr uuid = {{.UUIDandVersion.UUID}}
[msg]
  timestamp = "on"

[machine]
  type = "pc-q35-3.1"
  accel = "kvm"
  usb = "off"
  vmport = "off"
  dump-guest-core = "off"
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

[realtime]
  mlock = "off"

[rtc]
  base = "localtime"
  driftfix = "slew"

[chardev "charmonitor"]
  backend = "socket"
  path = "` + kvmStateDir + `{{.DisplayName}}.qmp"
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

[chardev "charserial0"]
  backend = "pty"

[device "serial0"]
  driver = "isa-serial"
  chardev = "charserial0"

{{if .EnableVnc}}
[vnc "default"]
  vnc = "0.0.0.0:{{if .VncDisplay}}{{.VncDisplay}}{{else}}0{{end}}"
  to = "99"
{{if .VncPasswd}}
  password = "on"
{{end}}
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

[device "pci.1"]
  driver = "pcie-root-port"
  port = "0x11"
  chassis = "2"
  bus = "pcie.0"
  addr = "0x2.0x1"

[device "usb"]
  driver = "qemu-xhci"
  p2 = "15"
  p3 = "15"
  bus = "pci.1"
  addr = "0x0"

[device "input0"]
  driver = "usb-tablet"
  bus = "usb.0"
  port = "1"
`

const qemuDiskTemplate = `
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "0x12"
  chassis = "3"
  bus = "pcie.0"
  addr = "0x2.0x2"

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
`

const qemu9PTemplate = `
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "0x12"
  chassis = "3"
  bus = "pcie.0"
  addr = "0x2.0x2"

[fsdev "fsdev{{.DiskID}}"]
  fsdriver = "local"
  security_model = "none"
  path = "{{.DiskFile}}"

[device "fs{{.DiskID}}"]
  driver = "virtio-9p-pci"
  fsdev = "fsdev{{.DiskID}}"
  mount_tag = "hostshare"
  bus = "pci.{{.PCIId}}"
  addr = "0x0"
`

const qemuNetTemplate = `
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "0x10"
  chassis = "1"
  bus = "pcie.0"
  multifunction = "on"
  addr = "0x2"

[netdev "hostnet{{.NetID}}"]
  type = "tap"
  br = "{{.Bridge}}"
  script = "no"
  downscript = "no"

[device "net{{.NetID}}"]
  driver = "e1000e"
  netdev = "hostnet{{.NetID}}"
  mac = "{{.Mac}}"
  bus = "pci.{{.PCIId}}"
  addr = "0x0"
`

const kvmStateDir = "/var/run/hypervisor/kvm/"

type kvmContext struct {
}

func newKvm() Hypervisor {
	return kvmContext{}
}

func (ctx kvmContext) Name() string {
	return "kvm"
}

func (ctx kvmContext) CreateDomConfig(domainName string, config types.DomainConfig, diskStatusList []types.DiskStatus,
	aa *types.AssignableAdapters, file *os.File) error {

	// ---> BIOS /usr/lib/xen/boot/ovmf-pvh.bin
	// ---> bootLoader
	// var serialAssignments []string
	// serialAssignments = append(serialAssignments, "pty")

	// Always prefer CDROM vdisk over disk
	// file.WriteString(fmt.Sprintf("boot = \"%s\"\n", "dc"))

	// we are hijacking some of the config fields for template creation
	// this is ok, since config is passed here by value
	config.Memory = (config.Memory + 1023) / 1024
	config.DisplayName = domainName
	if config.IsContainer {
		config.Kernel = "/hostfs/boot/kernel"
		config.Ramdisk = "/usr/lib/xen/boot/runx-initrd"
		config.ExtraArgs = config.ExtraArgs + " root=9p dhcp=1"
	}

	// render global device model settings
	t, _ := template.New("qemu").Parse(qemuConfTemplate)
	t.Execute(file, config)

	// render disk device model settings
	diskContext := struct {
		PCIId, DiskID        int
		DiskFile, DiskFormat string
		ro                   bool
	}{PCIId: 2, DiskID: 0}
	t, _ = template.New("qemuDisk").Parse(qemuDiskTemplate)
	t9p, _ := template.New("qemuDisk").Parse(qemu9PTemplate)
	for _, ds := range diskStatusList {
		if ds.Format == zconfig.Format_CONTAINER {
			diskContext.DiskFile = ds.FSVolumeLocation
			t9p.Execute(file, diskContext)
		} else {
			diskContext.DiskFile = ds.ActiveFileLocation
			diskContext.ro = ds.ReadOnly
			diskContext.DiskFormat = strings.ToLower(ds.Format.String())
			t.Execute(file, diskContext)
		}
		diskContext.PCIId = diskContext.PCIId + 1
		diskContext.DiskID = diskContext.DiskID + 1
	}

	// render network device model settings
	netContext := struct {
		PCIId, NetID int
		Mac, Bridge  string
	}{PCIId: diskContext.PCIId, NetID: 0}
	t, _ = template.New("qemuNet").Parse(qemuNetTemplate)
	for _, net := range config.VifList {
		// net.Bridge
		// net.Vif
		netContext.Mac = net.Mac
		netContext.Bridge = net.Bridge
		t.Execute(file, netContext)
		netContext.PCIId = netContext.PCIId + 1
		netContext.NetID = netContext.NetID + 1
	}

	return nil
	// imString := ""
	// for _, im := range config.IOMem {
	//	if imString != "" {
	//		imString += ","
	//	}
	//	imString += fmt.Sprintf("\"%s\"", im)
	// }
	// if imString != "" {
	//	file.WriteString(fmt.Sprintf("iomem = [%s]\n", imString))
	// }
	//
	// Gather all PCI assignments into a single line
	// Also irqs, ioports, and serials
	// irqs and ioports are used if we are pv; serials if hvm
	// var pciAssignments []typeAndPCI
	// var irqAssignments []string
	// var ioportsAssignments []string
	//
	// for _, irq := range config.IRQs {
	//	irqString := fmt.Sprintf("%d", irq)
	//	irqAssignments = addNoDuplicate(irqAssignments, irqString)
	// }
	// for _, adapter := range config.IoAdapterList {
	//	log.Debugf("configToXenCfg processing adapter %d %s\n",
	//		adapter.Type, adapter.Name)
	//	list := aa.LookupIoBundleGroup(adapter.Name)
	//	// We reserved it in handleCreate so nobody could have stolen it
	//	if len(list) == 0 {
	//		log.Fatalf("configToXencfg IoBundle disappeared %d %s for %s\n",
	//			adapter.Type, adapter.Name, domainName)
	//	}
	//	for _, ib := range list {
	//		if ib == nil {
	//			continue
	//		}
	//		if ib.UsedByUUID != config.UUIDandVersion.UUID {
	//			log.Fatalf("configToXencfg IoBundle not ours %s: %d %s for %s\n",
	//				ib.UsedByUUID, adapter.Type, adapter.Name,
	//				domainName)
	//		}
	//		if ib.PciLong != "" {
	//			tap := typeAndPCI{pciLong: ib.PciLong, ioType: ib.Type}
	//			pciAssignments = addNoDuplicatePCI(pciAssignments, tap)
	//		}
	//		if ib.Irq != "" && config.VirtualizationMode == types.PV {
	//			log.Infof("Adding irq <%s>\n", ib.Irq)
	//			irqAssignments = addNoDuplicate(irqAssignments,
	//				ib.Irq)
	//		}
	//		if ib.Ioports != "" && config.VirtualizationMode == types.PV {
	//			log.Infof("Adding ioport <%s>\n", ib.Ioports)
	//			ioportsAssignments = addNoDuplicate(ioportsAssignments, ib.Ioports)
	//		}
	//		if ib.Serial != "" && (config.VirtualizationMode == types.HVM || config.VirtualizationMode == types.FML) {
	//			log.Infof("Adding serial <%s>\n", ib.Serial)
	//			serialAssignments = addNoDuplicate(serialAssignments, ib.Serial)
	//		}
	//	}
	//}
	//if len(pciAssignments) != 0 {
	//	log.Infof("PCI assignments %v\n", pciAssignments)
	//	cfg := fmt.Sprintf("pci = [ ")
	//	for i, pa := range pciAssignments {
	//		if i != 0 {
	//			cfg = cfg + ", "
	//		}
	//		short := types.PCILongToShort(pa.pciLong)
	//		// USB controller are subject to legacy USB support from
	//		// some BIOS. Use relaxed to get past that.
	//		if pa.ioType == types.IoUSB {
	//			cfg = cfg + fmt.Sprintf("'%s,rdm_policy=relaxed'",
	//				short)
	//		} else {
	//			cfg = cfg + fmt.Sprintf("'%s'", short)
	//		}
	//	}
	//	cfg = cfg + "]"
	//	log.Debugf("Adding pci config <%s>\n", cfg)
	//	file.WriteString(fmt.Sprintf("%s\n", cfg))
	//}
	//irqString := ""
	//for _, irq := range irqAssignments {
	//	if irqString != "" {
	//		irqString += ","
	//	}
	//	irqString += irq
	//}
	//if irqString != "" {
	//	file.WriteString(fmt.Sprintf("irqs = [%s]\n", irqString))
	//}
	//ioportString := ""
	//for _, ioports := range ioportsAssignments {
	//	if ioportString != "" {
	//		ioportString += ","
	//	}
	//	ioportString += ioports
	//}
	//if ioportString != "" {
	//	file.WriteString(fmt.Sprintf("ioports = [%s]\n", ioportString))
	//}
	//serialString := ""
	//for _, serial := range serialAssignments {
	//	if serialString != "" {
	//		serialString += ","
	//	}
	//	serialString += "'" + serial + "'"
	//}
	//if serialString != "" {
	//	file.WriteString(fmt.Sprintf("serial = [%s]\n", serialString))
	//}
	// XXX log file content: log.Infof("Created %s: %s
}

func (ctx kvmContext) Create(domainName string, cfgFilename string) (int, error) {
	log.Infof("starting qemu for %s %s\n", domainName, cfgFilename)

	os.MkdirAll(kvmStateDir, 0777)

	pidFile := kvmStateDir + domainName + ".pid"
	qmpFile := kvmStateDir + domainName + ".qmp"
	cmd := "qemu-system-x86_64"
	args := []string{
		"-name", domainName,
		"-display", "none", // it seems to be impossible to wrap -display none into
		"-readconfig", cfgFilename,
		"-pidfile", pidFile,
		"-daemonize", "-S", "-no-user-config", "-nodefaults", "-no-shutdown", "-no-hpet",
		// "-cpu host",
	}
	// -uuid 87a53c25-d061-49f5-b5aa-78730b4e300f
	// -cpu IvyBridge-IBRS,ss=on,vmx=on,movbe=on,hypervisor=on,arat=on,tsc_adjust=on,mpx=on,rdseed=on,smap=on,clflushopt=on,sha-ni=on,umip=on,md-clear=on,arch-capabilities=on,xsaveopt=on,xsavec=on,xgetbv1=on,xsaves=on,pdpe1gb=on,3dnowprefetch=on,avx=off,f16c=off,hv_time,hv_relaxed,hv_vapic,hv_spinlocks=0x1fff
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Errorln("starting qemu failed ", err)
		log.Errorln("qemu output ", string(stdoutStderr))
		return 0, fmt.Errorf("starting qemu failed: %s\n", string(stdoutStderr))
	}
	log.Infof("done launching qemu device model")

	pidStr, err := ioutil.ReadFile(pidFile)
	if err != nil {
		log.Errorln("can't read from PID file ", err)
		return 0, fmt.Errorf("can't read from PID file %s (%v)", pidFile, err)
	}

	domainID, err := strconv.Atoi(strings.TrimSpace(string(pidStr)))
	if err != nil {
		log.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
		return 0, fmt.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
	}

	if status, err := getQemuStatus(qmpFile); err != nil || status != "prelaunch" {
		log.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
		return 0, fmt.Errorf("Can't extract domainID from %s: %s\n", string(pidStr), err)
	}
	return domainID, nil
}

func (ctx kvmContext) Start(domainName string, domainID int) error {
	qmpFile := kvmStateDir + domainName + ".qmp"

	if err := execCont(qmpFile); err != nil {
		log.Errorf("Can't extract domainID from %s\n", err)
		return fmt.Errorf("Can't extract domainID from %s\n", err)
	}

	if status, err := getQemuStatus(qmpFile); err != nil || status != "running" {
		log.Errorf("Can't extract domainID from %s\n", err)
		return fmt.Errorf("Can't extract domainID from %s\n", err)
	}

	return nil
}

func (ctx kvmContext) Stop(domainName string, domainID int, force bool) error {
	if err := execShutdown(kvmStateDir + domainName + ".qmp"); err != nil {
		log.Errorf("Can't extract domainID from %s\n", err)
		return fmt.Errorf("Can't extract domainID from %s\n", err)
	}
	return nil
}

func (ctx kvmContext) Delete(domainName string, domainID int) error {
	if err := execQuit(kvmStateDir + domainName + ".qmp"); err != nil {
		log.Errorf("Can't extract domainID from %s\n", err)
		return fmt.Errorf("Can't extract domainID from %s\n", err)
	}
	// we may want to wait a little bit here and actually kill qemu process if it gets wedged
	return nil
}

func (ctx kvmContext) Info(domainName string, domainID int) error {
	return nil
}

func (ctx kvmContext) LookupByName(domainName string, domainID int) (int, error) {
	return 0, nil
}

func (ctx kvmContext) Tune(domainName string, domainID int, vifCount int) error {
	return nil
}

func (ctx kvmContext) PCIReserve(long string) error {
	return nil
}

func (ctx kvmContext) PCIRelease(long string) error {
	return nil
}

func (ctx kvmContext) IsDeviceModelAlive(domid int) bool {
	return true
}

func (ctx kvmContext) GetHostCPUMem() (types.HostMemory, error) {
	return selfDomCPUMem()
}

func (ctx kvmContext) GetDomsCPUMem() (map[string]types.DomainMetric, error) {
	return nil, nil
}
