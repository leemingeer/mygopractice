package main

import (
	"fmt"
	"reflect"
	"strings"
)

type ArstorVolume struct {
	ArstorHost   string
	TargetPortal string
	TargetIqn    string
	TargetName   string
	PoolName     string
	VdiskName    string
	LunName      string
	Iqn          string
}

func main() {
	volumeID := "192.168.10.5--default--k8starget--pvcef5551ccb80a4aac8327cffc99fd98a8--iqn.2010-01.com.huayun:k8starget--pvcef5551ccb80a4aac8327cffc99fd98a8--iqn.2010-01.com.huayun:k8starget--192.168.10.5:3260"
	var volume ArstorVolume
	splittedVol := strings.Split(volumeID, "--")
	rType := reflect.TypeOf(volume)

	fmt.Printf("before: %+v", volume)
	for i := 0; i < rType.NumField(); i++ {
		fmt.Printf("split vol:", splittedVol[i])
		v := reflect.ValueOf(&volume).Elem()
		v.Field(i).Set(reflect.ValueOf(splittedVol[i]))
		v.Type()

	}
	fmt.Printf("after: %+v", volume)

}
