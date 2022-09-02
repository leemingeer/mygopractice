package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// saveVolumeData persists parameter data as json file at the provided location
func saveVolumeData(dir string, fileName string, data map[string]string) error {
	dataFilePath := filepath.Join(dir, fileName)
	file, err := os.Create(dataFilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to save volume data file %s: %v", dataFilePath, err))
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(data); err != nil {
		return errors.New(fmt.Sprintf("failed to save volume data file %s: %v", dataFilePath, err))
	}
	return nil
}

// GenerateID generates a random unique id.
func GenerateID(length int32) string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)[:length]
}

func main(){
    var volumeContext = map[string]string{
    	"name": "ming",
	}
	arstorUID := volumeContext["arstorUID"]
	if arstorUID == "" {
		arstorUID = GenerateID(32)
		volumeContext["arstorUID"] = arstorUID
	}
	dataDir, _:= os.Getwd()
	volDataFileName := "volumeContext.txt"
	fmt.Println(saveVolumeData(dataDir, volDataFileName, volumeContext))

	//portals := `["10.0.211.110:80"]`
	// get arstor host
	portalList := []string{}
	if err := json.Unmarshal([]byte(""), &portalList); err != nil {
		fmt.Println(err)
	}
	fmt.Println(portalList)

}