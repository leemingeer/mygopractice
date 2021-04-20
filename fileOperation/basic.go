package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main(){
	cidr := []string{"x","y", "z"}
	format_cidr := strings.Join(cidr, "-")
	fmt.Println(format_cidr)
	outDir := "/home/cloud"
	filename := "author.txt"
	kubeConfigFilePath := filepath.Join(outDir, filename)
	fmt.Println(kubeConfigFilePath)
	pkiPath := "/etc/kubernetes/ssl/"
	name := "ca"
	pathForCert := filepath.Join(pkiPath, fmt.Sprintf("%s.crt", name))
	fmt.Println(pathForCert)

	file := "basic.go"
	pemBlock, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	fmt.Println(string(pemBlock))

}
