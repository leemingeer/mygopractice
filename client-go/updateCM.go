package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"
)

/*
apiVersion: v1
data:
config.json: '{"arstors":[{"arstorid": "314f39fb-9dab-42fb-b755-7f0c5aab449e"", "portal":"10.0.211.5", "status":"Delete"}]}'
kind: ConfigMap
metadata:
name: test-configmap
namespace: default
*/

// fix structure
type arstorParams struct{
	TargetPortal string `json:"targetPortal"`
	Portal string `json:"portals"`
	Iqn string `json:"iqn"`
	IscsiInterface string `json:"iscsiInterface"`
	PoolName string `json:"poolName"`
	MaxVolumeSize string `json:"maxVolumeSize"`
	IopsLimit string `json:"iopsLimit"`
	BandwidthLimit string `json:"bandwidthLimit"`
}


type provider struct {
	Client *kubernetes.Clientset
	mu sync.RWMutex
}

func NewarstorParams(targetPortal, portal string) arstorParams{
	return arstorParams{
		TargetPortal: targetPortal,
		Portal: portal,
		Iqn: "iqn.2021-11.com.node01:k8starget",
		PoolName: "default",
		MaxVolumeSize: "50",
		IopsLimit: "25000",
		BandwidthLimit: "1000",
	}
}

func (p *provider) getFromConfigmap() error {
	defer wg.Done()
	client := p.Client
	configmap, err := client.CoreV1().ConfigMaps("default").Get(context.TODO(), "test-configmap", metav1.GetOptions{})
	if err != nil {
		glog.Errorf("Failed to get configmap : %v", err)
		return fmt.Errorf("Failed to get configmap")
	}
	datamap, ok := (*configmap).Data["arstors"]
	if !ok {
		glog.Errorf("Failed to get configmap : %v", err)
		return fmt.Errorf("Failed to get configmap")
	}
	glog.Infof("read: %v, %T", datamap, reflect.TypeOf(datamap))
	return nil
}

func (p *provider) addToConfigmap(arstorid, targetportal, portal string) error {
	defer wg.Done()
	p.mu.Lock()
	defer p.mu.Unlock()
	glog.Infof("addToConfigmap arstorid: %v, volumeID: %v, status: %v", arstorid, targetportal, portal)
	client := p.Client
	configmap, err := client.CoreV1().ConfigMaps("default").Get(context.TODO(), "arstor-conf", metav1.GetOptions{})
	if err != nil {
		glog.Errorf("Failed to get configmap : %v", err)
		return fmt.Errorf("Failed to get configmap")
	}
	datamap, ok := configmap.Data["arstors"]
	if !ok {
		glog.Errorf("Failed to get configmap : %v", err)
		return fmt.Errorf("Failed to get configmap")
	}
    //glog.Infof("%v, %v", reflect.TypeOf(datamap), datamap)
	arstors := map[string]arstorParams{}
	// map string to map
	err = yaml.Unmarshal([]byte(datamap), &arstors)
	if err != nil {
		glog.Errorf("yaml decode failed : %v", err)
		return fmt.Errorf("Failed to get configmap")
	}
	glog.Infof("arstors: %#v", arstors)

	arstors[arstorid] = NewarstorParams(targetportal, portal)
	// to byte stream
	arstorsNew, err1 := yaml.Marshal(arstors)
	if err1 != nil {
		glog.Errorf("Failed to get configmap : %v", err1)
		return fmt.Errorf("Failed to get configmap")
	}

	configmap.Data["arstors"]= string(arstorsNew)
	configmap, err2 := client.CoreV1().ConfigMaps("default").Update(context.TODO(), configmap, metav1.UpdateOptions{})
	if err2 != nil {
		glog.Errorf("update configmap failed : %v", err2)
		return fmt.Errorf("update configmap error")
	}
	glog.Infof("update: %v", configmap)
	return nil
}

var wg sync.WaitGroup
func main(){
	var err error
	var config *rest.Config
	var kubeconfig *string
	var mu = sync.RWMutex{}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "[可选] kubeconfig 绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 绝对路径")
	}
	// 初始化 rest.Config 对象
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}
	// 创建 Clientset 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}


	p := provider{clientset, mu}

	for i:=0;i<10;i++{
		wg.Add(1)
		arstorid := "id" + strconv.Itoa(i)
		portalid := "[192.168.2." + strconv.Itoa(i) + ":80" + "]"
		targetPortal := "10.0.211." + strconv.Itoa(i) + ":3268"
		if i == 5{
			targetPortal = strconv.Itoa(i) + ":3269"
		}
		go p.addToConfigmap(arstorid, targetPortal, portalid)
	}
	for i:=0;i<10;i++{
       wg.Add(1)
		go p.getFromConfigmap()
	}

	glog.Info("main goroutine waiting\n")
	wg.Wait()
}
