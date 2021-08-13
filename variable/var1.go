package main

import (
	"fmt"
	"github.com/golang/glog"
	"path"
	"reflect"
	"strconv"
	"strings"
	"encoding/json"
)

func setName(n *string){
	*n = "Lee"
}

type createSnap struct{
	SourceFilePath string `json:"sourceFilePath"`
}

func RoundUpSize(volumeSizeBytes int64, allocationUnitBytes int64) int64 {
	roundedUp := volumeSizeBytes / allocationUnitBytes
	fmt.Println("roundedUp1", roundedUp)
	if volumeSizeBytes%allocationUnitBytes > 0 {
		roundedUp++
	}
	fmt.Println("roundedUp2", roundedUp)
	return roundedUp
}

func main(){
	var name string
	setName(&name)
	fmt.Println(name)

	var v string

	fmt.Println(v == "")

	var s string = "13.11.12.5-default-mingk8starget-mingk8svdisk-iqn.1994-05.com.redhat:3a6083908637-1"
	lun := strings.Split(s, "-")
	fmt.Println(len(lun), lun)

	var s1 ="xyz"
	id := strings.Split(s1, ":")[0]
	fmt.Println(id)
	ret := path.Join("poolName", "vDiskName")

	fmt.Println(ret)

	size := "21474836480"
	i, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %v with type %s!\n", i, reflect.TypeOf(i))

	poolName := "xy"
	vDiskName := "XY"
	sourceFilePath := path.Join(poolName, vDiskName)
	glog.Infof("join:%s, %v", sourceFilePath,createSnap{sourceFilePath})
	createSnapshotReq, err := json.Marshal(createSnap{sourceFilePath})
	fmt.Printf("%v, af %s",createSnapshotReq, string(createSnapshotReq))

    l1 := "1234"
    l2, _ := strconv.Atoi(l1)
    fmt.Println(l2, reflect.TypeOf(l2))
    lunlun := "1"
	var lunVal int32
	if lunlun != "" {
		l, err := strconv.Atoi(lunlun)
     fmt.Println(err)
		lunVal = int32(l)
	}
	fmt.Println(lunVal, reflect.TypeOf(lunVal))
	volSizeBytes := int64(26843545600)
	fmt.Println(volSizeBytes, reflect.TypeOf(volSizeBytes))
	expandSizeInGB := string(RoundUpSize(volSizeBytes, 1024*1024*1024))
	fmt.Println(expandSizeInGB)
	fmt.Println(strconv.Itoa(35))

	ss := string(97)
	fmt.Println(ss)
}
