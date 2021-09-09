# 分两种场景，在vm中的数据平面和在物理host上的数据平面

# 第一种在vm
# 1. 在openstack中创建两个外部网络，这里假设是extvlan1420和extvlan1421，这里extvlan1420作为主管集群的管理网络，extvlan1421作为数据网络
neutron net-create <net-name> --provider:physical_network self --provider:segmentation_id <vlan-id> --provider:network_type vlan
neutron subnet-create <net-id> <cidr>

# 示例：
neutron net-create extvlan1420 --provider:physical_network self --provider:segmentation_id 1420 --provider:network_type vlan
neutron subnet-create extvlan1420 178.119.220.0/24
neutron net-create extvlan1421 --provider:physical_network self --provider:segmentation_id 1421 --provider:network_type vlan
neutron subnet-create extvlan1421 178.119.221.0/24

# 2. 在数据外部网络（extvlan1421）中创建三个端口，给物理机备用
neutron port-create <ext-net-id> --fixed-ip subnet_id=<ext-subnet-id>,ip_address=<data-ip> --binding:host_id=<hostname>
# 示例：
neutron port-create 3482bdcf-c9c2-45c3-9ccc-a92612ffefda --fixed-ip subnet_id=5c69ecdc-872c-4328-8e67-690fee319112,ip_address=178.119.221.11 --binding:host_id=host1
neutron port-create 3482bdcf-c9c2-45c3-9ccc-a92612ffefda --fixed-ip subnet_id=5c69ecdc-872c-4328-8e67-690fee319112,ip_address=178.119.221.12 --binding:host_id=host2
neutron port-create 3482bdcf-c9c2-45c3-9ccc-a92612ffefda --fixed-ip subnet_id=5c69ecdc-872c-4328-8e67-690fee319112,ip_address=178.119.221.13 --binding:host_id=host3

# 3. 给三台master虚拟机添加数据网卡
nova interface-attach <vm-id> --net-id <data-ext-net-id>
# 示例：
nova interface-attach fa51a775-bb6b-4c68-8669-c6a9089dc1a3 --net-id 3482bdcf-c9c2-45c3-9ccc-a92612ffefda
nova interface-attach 3e8bca4b-6948-4b0a-bd66-b8d584e77af8 --net-id 3482bdcf-c9c2-45c3-9ccc-a92612ffefda
nova interface-attach 804a7aff-734c-46bc-bf2e-703b50a1d279 --net-id 3482bdcf-c9c2-45c3-9ccc-a92612ffefda

# 第二种 在物理host上

# 创建ovs port
# brVlan是ovs bridge名称
# tag=1421 和数据网络的vlan保持一致，需要二层通信
# data是ovs port的名称，calico中数据网卡都统一改成data
ovs-vsctl add-port brVlan data tag=1421 -- set interface data type=internal

# 为ovs port配置网卡文件
# cat <<EOF>/etc/sysconfig/network-scripts/ifcfg-data
TYPE=OVSPort
DEVICE=data
ONBOOT=yes
BOOTPROTO=static
IPADDR=178.119.221.11
PREFIX=24
OVS_BRIDGE=brVlan
<EOF>
# 启动数据网卡
ifup data