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
  type = "pc-q35-3.1"
  accel = "kvm"
  usb = "off"
  vmport = "off"
  dump-guest-core = "off"
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

[chardev "charserial0"]
  backend = "pty"

[device "serial0"]
  driver = "isa-serial"
  chardev = "charserial0"

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
  port = "0x11"
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
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "0x12"
  chassis = "3"
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
`

const qemu9PTemplate = `
[fsdev "fsdev{{.DiskID}}"]
  fsdriver = "local"
  security_model = "none"
  path = "{{.DiskFile}}"

[device "fs{{.DiskID}}"]
  driver = "virtio-9p-pci"
  fsdev = "fsdev{{.DiskID}}"
  mount_tag = "hostshare"
  addr = "{{.PCIId}}"
`

const qemuNetTemplate = `
[device "pci.{{.PCIId}}"]
  driver = "pcie-root-port"
  port = "0x10"
  chassis = "1"
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

const kvmStateDir = "/var/run/hypervisor/kvm/"

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
}

func newKvm() Hypervisor {
	return kvmContext{
		domains: map[string]int{},
	}
}

func (ctx kvmContext) Name() string {
	return "kvm"
}

func (ctx kvmContext) CreateDomConfig(domainName string, config types.DomainConfig, diskStatusList []types.DiskStatus,
	aa *types.AssignableAdapters, file *os.File) error {
	// we are hijacking some of the config fields for template creation
	// this is ok, since config is passed here by value
	config.Memory = (config.Memory + 1023) / 1024
	config.DisplayName = domainName
	config.BootLoader = "" // could be /usr/lib/xen/boot/ovmf.bin
	if config.IsContainer {
		config.Kernel = "/hostfs/boot/kernel"
		config.Ramdisk = "/usr/lib/xen/boot/runx-initrd"
		config.ExtraArgs = config.ExtraArgs + " root=9p-kvm dhcp=1"
	}

	// render global device model settings
	t, _ := template.New("qemu").Parse(qemuConfTemplate)
	if err := t.Execute(file, config); err != nil {
		return logError("can't write to config file %s (%v)", file.Name(), err)
	}

	// render disk device model settings
	diskContext := struct {
		PCIId, DiskID        int
		DiskFile, DiskFormat string
		ro                   bool
	}{PCIId: 3, DiskID: 0}
	t, _ = template.New("qemuDisk").Parse(qemuDiskTemplate)
	t9p, _ := template.New("qemuDisk").Parse(qemu9PTemplate)
	for _, ds := range diskStatusList {
		if ds.Format == zconfig.Format_CONTAINER {
			diskContext.DiskFile = ds.FSVolumeLocation
			if err := t9p.Execute(file, diskContext); err != nil {
				return logError("can't write to config file %s (%v)", file.Name(), err)
			}
		} else {
			diskContext.DiskFile = ds.ActiveFileLocation
			diskContext.ro = ds.ReadOnly
			diskContext.DiskFormat = strings.ToLower(ds.Format.String())
			if err := t.Execute(file, diskContext); err != nil {
				return logError("can't write to config file %s (%v)", file.Name(), err)
			}
		}
		diskContext.PCIId = diskContext.PCIId + 1
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

	return nil
}

func (ctx kvmContext) Create(domainName string, cfgFilename string) (int, error) {
	log.Infof("starting KVM domain %s with config %s\n", domainName, cfgFilename)

	os.MkdirAll(kvmStateDir+domainName, 0777)

	pidFile := kvmStateDir + domainName + "/pid"
	qmpFile := kvmStateDir + domainName + "/qmp"
	consFile := kvmStateDir + domainName + "/cons"
	cmd := "qemu-system-x86_64"
	args := []string{
		"-name", domainName,
		"-display", "none",
		"-readconfig", cfgFilename,
		"-pidfile", pidFile,
		"-daemonize", "-S", "-no-user-config", "-nodefaults", "-no-shutdown", "-no-hpet",
	}
	// "-cpu host",
	// -cpu IvyBridge-IBRS,ss=on,vmx=on,movbe=on,hypervisor=on,arat=on,tsc_adjust=on,mpx=on,rdseed=on,smap=on,clflushopt=on,sha-ni=on,umip=on,md-clear=on,arch-capabilities=on,xsaveopt=on,xsavec=on,xgetbv1=on,xsaves=on,pdpe1gb=on,3dnowprefetch=on,avx=off,f16c=off,hv_time,hv_relaxed,hv_vapic,hv_spinlocks=0x1fff
	stdoutStderr, err := wrap.Command(cmd, args...).CombinedOutput()
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
