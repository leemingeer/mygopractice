package main

import (
	"fmt"
	"reflect"
	"strings"
	//"reflect"
)

func main(){
	snapshotHandle := "125793bb-215f-4a40-ba8e-f49003c71258__10.0.211.100--default--k8starget--pvc77c4764afdfa46e79619c56f512593e9--iqn.2010-01.com.huayun:k8starget--pvc77c4764afdfa46e79619c56f512593e9--iqn.2010-01.com.huayun:k8starget--10.10.10.5:3260__snapshots/snapshot_261/snapshot-21d6c595-32f7-491d-b775-c3e8b9904ae0"
	fmt.Println(reflect.TypeOf((snapshotHandle)[strings.LastIndex(snapshotHandle, "__")+2:]))
}
