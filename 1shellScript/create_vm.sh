#!/usr/bin/env bash
source /root/ArcherAdmin-openrc 
#source /root/admin-openrc

#network_id1420=693b3101-787c-4a56-8b13-8b67fcca0b9e
#network_id1421=3482bdcf-c9c2-45c3-9ccc-a92612ffefda
#subnet_id=cdb9b94b-6bf4-4450-9238-cf7e5ef8db28

extvlan1428_mgr=e0e74a5f-a15d-4830-9434-8cac9d7a1823
extvlan1429_data=cdc55fa6-f877-492f-98ca-51415a10e7d5
network_mgr=$extvlan1428_mgr
network_data=$extvlan1429_data

system_image_id=c8cf7587-11ec-4f12-b0a1-65f8f376944a  #7cc171a3-1ec5-451b-9243-9438ef01bf28
#volume_type=maxta_replica2
volume_type=basic-replica2
flavor_id=4
server_group_id=99178d12-7440-496b-ad2d-d307b67293af

vm_name=$1
if [ -z "$vm_name" ];then
	vm_name=phtest
fi

vm_count=$2
if [ -z "$vm_count" ];then
	vm_count=1
fi

vm_host="nova:archcnstcm8395"

#if [ -z "$3" ];then
#	vm_host=$vm_zone
#else
#	vm_host="${vm_zone}:node0"
#fi

function rand(){
	min=$1
	max=$(($2-$min+1))
	num=$(date +%s%N)
	echo $(($num%$max+$min))
}

priority=$(rand 1 10)
echo $priority

nova  --debug boot \
	--min-count $vm_count \
	--max-count $vm_count \
	--availability-zone $vm_host \
	--flavor  $flavor_id  \
	--block-device id=$system_image_id,source=image,device=vda,dest=volume,size=60,bootindex=0,volume_type=$volume_type,shutdown=remove \
	--nic net-id=$network_mgr \
	--nic net-id=$network_data \
	--priority $priority \
  --meta storage_type=ARSTOR \
  --security-groups=c545cde6-c3ff-4c65-ade6-75bfb452eafd \
	${vm_name}_${priority}

##--hint group=${server_group_id}
