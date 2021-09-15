#!/usr/bin/env bash

# 主管集群node节点磁盘划分
#
#需要单独从系统盘上划一个60G的空间用于containerd，并且对其进行格式化，挂载到/var/lib/containerd目录下
##划分出来的分区以/dev/sdc1为例
$mkfs.xfs /dev/sdc1

$mount -t xfs  /dev/sdc1 /var/lib/containerd

##/etc/fstab中添加
/dev/sdc1 /var/lib/containerd  xfs noatime 0 0