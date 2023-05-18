#!/bin/ash
#
# Copyright (c) 2023 Zededa, Inc.
# SPDX-License-Identifier: Apache-2.0

K3S_VERSION=v1.26.3+k3s1
KUBEVIRT_VERSION=v0.59.0
LONGHORN_VERSION=v1.4.2
CDI_VERSION=v1.56.0

INSTALL_LOG=/var/log/install.log
date >> $INSTALL_LOG

logmsg() {
   local MSG
   local TIME
   MSG="$*"
   TIME=$(date +"%F %T")
   echo "$TIME : $MSG"  >> $INSTALL_LOG
}

setup_cgroup () {
   echo "cgroup /sys/fs/cgroup cgroup defaults 0 0" >> /etc/fstab
}


                       
check_network_connection () {
 while true; do

   IP="8.8.8.8"
   if ping -c 5  $IP > /dev/null 2>&1; then
        logmsg "Network available, ping to $IP was successful."
        break;
   else
        logmsg "Ping to $IP failed."
   fi

   sleep 5
 done
}                                                               

wait_for_default_route() {
  while read -r iface dest gw flags refcnt use metric mask mtu window irtt; do
    if [ "$dest" = "00000000" ] && [ "$mask" = "00000000" ]; then
      logmsg "Default route found"
      return 0
    fi
    logmsg "waiting for default route"
    sleep 1
  done < /proc/net/route

  return 1
}

#Prereqs
setup_prereqs () {
	modprobe tun                                                        
	modprobe vhost_net    
	modprobe fuse         
	modprobe iscsi_tcp
	#Needed for iscsi tools
	mkdir -p /run/lock
	/usr/sbin/iscsid start
	mount --make-rshared /
	setup_cgroup                                                   

	#Check network and default routes are up
	wait_for_default_route
	check_network_connection
}

#Make sure all prereqs are set
setup_prereqs
                                                              

#Is this ZFS block device ?
if [ -b /dev/zvol/persist/clustered-storage ]; then
        mount /dev/zvol/persist/clustered-storage /var/lib  ## This is where we persist the cluster components (k3s containers)
fi

HOSTNAME=`/bin/hostname`
#Forever loop every 15 secs
while true;
do
if [ ! -f /var/lib/all_components_initialized ]; then
	if [ ! -f /var/lib/k3s_initialized ]; then
		#/var/lib is where all kubernetes components get installed.
		logmsg "Installing K3S version $K3S_VERSION"
		mkdir -p /var/lib/k3s/bin
		/usr/bin/curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=${K3S_VERSION} INSTALL_K3S_SKIP_ENABLE=true INSTALL_K3S_BIN_DIR=/var/lib/k3s/bin sh -
		ln -s /var/lib/k3s/bin/* /usr/bin
		logmsg "Initializing K3S version $K3S_VERSION"
		nohup /usr/bin/k3s server --config /etc/rancher/k3s/config.yaml &
		#wait until k3s is ready
		logmsg "Looping until k3s is ready"
		until kubectl get node | grep $HOSTNAME | awk '{print $2}' | grep 'Ready'; do sleep 5; done
		logmsg "k3s is ready on this node"
		# Default location where clients will look for config
		ln -s /etc/rancher/k3s/k3s.yaml ~/.kube/config
		touch /var/lib/k3s_initialized
	fi

	if [ ! -f /var/lib/kubevirt_initialized ]; then
		# This patched version will be removed once the following PR https://github.com/kubevirt/kubevirt/pull/9668 is merged
		logmsg "Installing patched Kubevirt"
		kubectl apply -f /etc/kubevirt-operator.yaml
		kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/${KUBEVIRT_VERSION}/kubevirt-cr.yaml

		#CDI (containerzed data importer) is need to convert qcow2/raw formats to Persistent Volumes and Data volumes
		#Since CDI goes with kubevirt we install with that.
		logmsg "Installing CDI version $CDI_VERSION" 
		kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/$CDI_VERSION/cdi-operator.yaml
		kubectl create -f https://github.com/kubevirt/containerized-data-importer/releases/download/$CDI_VERSION/cdi-cr.yaml
		touch /var/lib/kubevirt_initialized
	fi

	if [ ! -f /var/lib/longhorn_initialized ]; then
		logmsg "Installing longhorn version ${LONGHORN_VERSION}"
		kubectl apply -f  https://raw.githubusercontent.com/longhorn/longhorn/v1.4.2/deploy/longhorn.yaml
		touch /var/lib/longhorn_initialized
	fi

	if [ -f /var/lib/k3s_initialized -a -f /var/lib/kubevirt_initialized -a -f /var/lib/longhorn_initialized ]; then
		logmsg "All components initialized"
		touch /var/lib/all_components_initialized
	fi
else
	ps -ef | grep "k3s server" | grep -v "grep" >> $INSTALL_LOG
	if [ $? -eq 0 ]; then
		logmsg "k3s is alive "
	else
		## Must be after reboot
		ln -s /var/lib/k3s/bin/* /usr/bin
		logmsg "Starting k3s server after reboot"
		nohup /usr/bin/k3s server --config /etc/rancher/k3s/config.yaml &
		logmsg "Looping until k3s is ready"
		until kubectl get node | grep $HOSTNAME | awk '{print $2}' | grep 'Ready'; do sleep 5; done
		logmsg "k3s is ready on this node"
		# Default location where clients will look for config
		ln -s /etc/rancher/k3s/k3s.yaml ~/.kube/config
	fi
fi
	sleep 15
done
