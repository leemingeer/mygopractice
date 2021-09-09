#!/bin/bash env

VirtualEXTIP=`pcs resource show VirtualEXTIP | grep 'ip=' | awk -F "ip=" '{print $2}' | tr -d ' \n'`

# modify haproxy to listen external VIP
sed -i '/listen keystone_admin/a \ \ bind $VirtualEXTIP:45357' /etc/haproxy/haproxy.cfg
sed -i '/listen cinder_api/a \ \  bind $VirtualEXTIP:9776' /etc/haproxy/haproxy.cfg
sed -i '/listen neutron_api/a \ \ bind $VirtualEXTIP:10696' /etc/haproxy/haproxy.cfg
sed -i '/listen nova_compute_api/a \ \ bind $VirtualEXTIP:9774' /etc/haproxy/haproxy.cfg
systemctl daemon-reload && systemctl restart haproxy && systemctl status haproxy


# nova compute
id=`openstack endpoint list -f value |grep 'nova compute' |grep public |awk '{print $1}'`
openstack endpoint set $id --url "http://$VirtualEXTIP:9774/v2.1/%(tenant_id)s"

# neutron
id=`openstack endpoint list -f value |grep 'neutron' |grep public |awk '{print $1}'`
openstack endpoint set $id --url http://$VirtualEXTIP:10696

# cinderv2
id=`openstack endpoint list -f value |grep 'cinderv2' |grep public |awk '{print $1}'`
openstack endpoint set $id --url "http://$VirtualEXTIP:9776/v2/%(tenant_id)s"

# check
cat  /etc/haproxy/haproxy.cfg |egrep -A4 "keystone_admin|cinder_api|neutron_api|nova_compute_api"
openstack endpoint list -f value |grep public