package main

import (
	"encoding/base64"
	"log"

	//"encoding/base64"
	//"gopkg.in/yaml.v2"
	"fmt"
	"gopkg.in/yaml.v3"

	//"log"
)

const testData = `
  apikey: 1-OjrNqlteLPz8zsLItlLBsB3sJNqKjNPszJQsW2aJNo48iIyTDzEwn2ZbVWBUjWh7 
  apiurl: http://178.104.162.80/api/v2.0
  configurations:
    default:
      dataset: testpool1/ljxtest1
      deletePolicy: delete
  iscsi:
    portal: 192.168.0.11:3260
    portalid: 1
  nfs:
    allowedhosts: []
    #allow client ip range
    allowednetworks: []
    server: 178.104.162.80
  password: 1
  username: root
`
type CSIConfiguration map[string]*FreeNAS

type FreeNAS struct {
	APIUrl string `yaml:"apiurl"`

	// FreeNAS : User & password
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`

	// TrueNAS Core : ApiKey
	APIKey string `yaml:"apikey,omitempty"`

	// NFS holds global nfs configuration
	NFS *NFS `yaml:"nfs,omitempty"`

	// ISCSI holds global iSCSI configuration
	ISCSI *ISCSI `yaml:"iscsi,omitempty"`

	Configurations map[string]*Configuration `yaml:"configurations,omitempty"`

	name                  string
	rootDsToConfiguration map[string]*Configuration
}

// NFS holds configuration for Filesystem Volumes
type NFS struct {
	Server string `yaml:"server"`

	AllowedHosts    []string `yaml:"allowedhosts"`
	AllowedNetworks []string `yaml:"allowednetworks"`
}

type ISCSI struct {
	Portal                 string `yaml:"portal"`
	PortalID               int    `yaml:"portalid"`
	VolBlockSize           string `yaml:"volblocksize,omitempty"`
	DisableReportBlockSize bool   `yaml:"disableReportBlockSize,omitempty"`
}


// Configuration holds common configuration for nfs/iscsi shares
type Configuration struct {
	// Dataset specifies the dataset to hold volumes
	Dataset string `yaml:"dataset"`

	// Sparse means to allocate sparse datasets/volumes (i.e. without refreservation)
	// Quota will be enforced always.
	Sparse bool `yaml:"sparse,omitempty"`

	// DeletePolicy specifies delete policy for this configuration
	DeletePolicy DeletePolicy `yaml:"deletePolicy"`

	// NFS holds nfs sub-configuration
	NFS *NFS `yaml:"nfs,omitempty"`

	// ISCSI holds iSCSI sub-configuration
	ISCSI *ISCSI `yaml:"iscsi,omitempty"`
}

type DeletePolicy string


func main(){
    //var nascfg CSIConfiguration
	var nascfg *FreeNAS
	testData1 :="ICBhcGlrZXk6IDEtT2pyTnFsdGVMUHo4enNMSXRsTEJzQjNzSk5xS2pOUHN6SlFzVzJhSk5vNDhpSXlURHpFd24yWmJWV0JValdoNw0KICBhcGl1cmw6IGh0dHA6Ly8xNzguMTA0LjE2Mi44MC9hcGkvdjIuMA0KICBjb25maWd1cmF0aW9uczoNCiAgICBkZWZhdWx0Og0KICAgICAgZGF0YXNldDogdGVzdHBvb2wxL2xqeHRlc3QxDQogICAgICBkZWxldGVQb2xpY3k6IGRlbGV0ZQ0KICBpc2NzaToNCiAgICBwb3J0YWw6IDE5Mi4xNjguMC4xMTozMjYwDQogICAgcG9ydGFsaWQ6IDENCiAgbmZzOg0KICAgIGFsbG93ZWRob3N0czogW10NCiAgICAjYWxsb3cgY2xpZW50IGlwIHJhbmdlDQogICAgYWxsb3dlZG5ldHdvcmtzOiBbXQ0KICAgIHNlcnZlcjogMTc4LjEwNC4xNjIuODANCiAgcGFzc3dvcmQ6IDENCiAgdXNlcm5hbWU6IHJvb3QNCg=="

//	testData2 := `{"apikey":"1-OjrNqlteLPz8zsLItlLBsB3sJNqKjNPszJQsW2aJNo48iIyTDzEwn2ZbVWBUjWh7",
//	"apiurl":"http:\/\/178.104.162.80\/api\/v2.0",
//    "configurations":{
//        "default":{
//            "dataset":"testpool1\/ljxtest1",
//            "deletePolicy":"delete"
//        }
//    },
//    "iscsi":{
//        "portal":"192.168.0.11:3260",
//        "portalid":1
//    },
//    "nfs":{
//        "allowedhosts":[
//
//        ],
//        "allowednetworks":[
//
//        ],
//        "server":"178.104.162.80"
//    },
//    "password":"1",
//    "username":"root"
//}`
	decodeBytes, err := base64.StdEncoding.DecodeString(testData1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))
	err = yaml.Unmarshal(decodeBytes, nascfg)
    fmt.Println(err)
	//fmt.Println(nascfg["default"].APIKey)
    //fmt.Println(nascfg["default"].Configurations["default"].Dataset)
	fmt.Println(nascfg.APIKey)
	fmt.Println(nascfg.Configurations["default"].Dataset)

	//encodeString := base64.StdEncoding.EncodeToString([]byte(testData))
	//fmt.Println(encodeString)
	//

}