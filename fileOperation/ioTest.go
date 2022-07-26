package main

import (
	"fmt"
	"io"
	"io/ioutil"
	utilpath "k8s.io/utils/path"
	"os"
	"syscall"
)


func RemoveAllOneFilesystemCommon(path string, remove func(string) error) error {
	// Simple case: if Remove works, we're done.
	err := remove(path)
	if err == nil || os.IsNotExist(err) {
		return nil
	}

	// Otherwise, is this a directory we need to recurse into?
	dir, serr := os.Lstat(path)
	if serr != nil {
		if serr, ok := serr.(*os.PathError); ok && (os.IsNotExist(serr.Err) || serr.Err == syscall.ENOTDIR) {
			return nil
		}
		return serr
	}
	if !dir.IsDir() {
		// Not a directory; return the error from remove.
		return err
	}



	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// Race. It was deleted between the Lstat and Open.
			// Return nil per RemoveAll's docs.
			return nil
		}
		return err
	}

	// Remove contents & return first error.
	err = nil
	for {
		names, err1 := fd.Readdirnames(100)
		for _, name := range names {
			err1 := RemoveAllOneFilesystemCommon(path+string(os.PathSeparator)+name, remove)
			if err == nil {
				err = err1
			}
		}
		if err1 == io.EOF {
			break
		}
		// If Readdirnames returned an error, use it.
		if err == nil {
			err = err1
		}
		if len(names) == 0 {
			break
		}
	}

	// Close directory, because windows won't remove opened directory.
	fd.Close()

	// Remove directory.
	err1 := remove(path)
	if err1 == nil || os.IsNotExist(err1) {
		return nil
	}
	if err == nil {
		err = err1
	}
	return err
}

func main(){
	podInfos, err := ioutil.ReadDir("/root")
	if err != nil{
		os.Exit(1)
	}
	for _, info := range podInfos {
		fmt.Println(info.Name())
	}
	fmt.Println("----------------------------")
	volumePluginDirs, _ := utilpath.ReadDirNoStat("/root")
	fmt.Println("ReadDirNoStat:", volumePluginDirs)
	fmt.Println("----------------------------")
	volumePath := "/root/ming/testrmdir"

	//if err := RemoveAllOneFilesystemCommon(volumePath, os.Remove); err != nil {
	//	fmt.Printf("os.remove,  path %v: error:%v", volumePath, err)
	//} else {
	//	fmt.Println("delete ok!") // 会递归进行删除，
	//
	//}

	if err := RemoveAllOneFilesystemCommon(volumePath, syscall.Rmdir); err != nil {
		fmt.Printf("syscall.Rmdir,  path %v: error:%v", volumePath, err)
	} else {
		fmt.Println("delete ok!") // 递归删除时，只会删掉空目录，如果目录非空，会报错
	}

	//if err := os.RemoveAll(volumePath); err != nil {
	//	fmt.Printf("orphaned pod found, but failed to rmdir() volume at path %v: %v", volumePath, err)
	//} else {
	//	fmt.Println("delete ok!")
	//
	//}
}