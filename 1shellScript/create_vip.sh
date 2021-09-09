#!/bin/bash

# control plane
#network_id=e0e74a5f-a15d-4830-9434-8cac9d7a1823
#subnet_id=f1faa88f-a040-4f51-8055-ded9abe10bef
#network_control_base=178.119.228
#vip=$network_control_base.166

# dataplane
network_id=cdc55fa6-f877-492f-98ca-51415a10e7d5
subnet_id=426a9592-fa3f-4610-a6e9-d8789adab0a0
network_dp_base=178.119.229
vip=$network_dp_base.166

# 1 创建port占位VIP对应的IP地址，使得此IP不会再被分配
neutron port-create ${network_id} --fixed-ip subnet_id=${subnet_id},ip_address=${vip}

masters=(
f30a652c-6f03-4eb5-a74f-2bde5fa362fd
60ede45c-1a5c-4c74-a8d1-60ed8ab0d456
fcd599df-9c4f-4345-bd0e-8d79b961c258
)

nodes=(
e5e19a7a-99a7-4cbf-99a9-e79499eefebc
c71b1d09-c04a-4022-b28f-ed3f8122bfbc
)
NODES=(${masters[@]} ${nodes[@]})

portids=()

# 2. 找到三个（有几个找几个）master节点对应的虚拟机端口
for N in ${NODES[@]};do
pid=`nova interface-list $N  | grep $network_dp_base | awk -F "|" '{print $3}'| tr -d ' \n'`
portids+=($pid)
done

# 3 更新所有虚拟机端口，设置allow_address_pair，使得虚拟机网卡允许VIP流量通过
for pid in ${portids[@]};do
neutron port-update $pid --allowed-address-pair ip_address=${vip}
done

