package main

import (
	"k8s.io/klog/v2"
	"os/exec"
	"strings"
)
type nodeServer struct{}
// getFSType returns the filesystem for the supplied devName.
func (s *nodeServer) getFSType(devName string) string {
	//l := s.log.WithField("func", "getFSType()")
	cmd := exec.Command("blkid", devName)
	klog.Infof("Executing command: %+v", cmd)
	out, err := cmd.CombinedOutput()
	fsType := ""
	if err != nil {
		klog.Infof("Could not get FSType for device.")
		return fsType
	}

	if strings.Contains(string(out), "TYPE=") {
		for _, v := range strings.Split(string(out), " ") {
			if strings.Contains(v, "TYPE=") {
				fsType = strings.Split(v, "=")[1]
				fsType = strings.Replace(fsType, "\"", "", -1)
				fsType = strings.TrimSpace(fsType)
			}
		}
	}
	return fsType
}