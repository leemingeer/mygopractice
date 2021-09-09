#!/bin/sh

vlan_id=1420
gateway=1
network_segment=178.119.220

physical_network=self
network_type=vlan

vlan_name=extvlan$vlan_id

neutron net-create --provider:segmentation_id ${vlan_id} --provider:physical_network ${physical_network} --provider:network_type ${network_type} ${vlan_name}

net_id=`neutron net-list |grep $vlan_name |awk '{print $2}'`

neutron subnet-create --allocation-pool start=${network_segment}.2,end=${network_segment}.253 --gateway ${network_segment}.${gateway} --dns-nameserver 114.114.114.114 ${net_id} ${network_segment}.0/24
