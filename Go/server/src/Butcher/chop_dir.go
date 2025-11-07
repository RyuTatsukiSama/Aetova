package butcher

import (
	"aetova/server/utils"
	"fmt"
	"os"
	"strings"
)

func chopDir(dirPath string) (manifestDir utils.ManifestDir, err error) {
	// get the data of the directory
	dir, _err := os.ReadDir(dirPath)
	if _err != nil {
		return utils.ManifestDir{}, _err
	}

	manifestDir.Name = strings.Split(dirPath, "/")[len(strings.Split(dirPath, "/"))-1]
	createDir(dirPath)

	fmt.Println("Chop dir", manifestDir.Name, "start") // ? Here for debug purpose ( replace by docLogger later )

	// browse file & subDir
	for _, entry := range dir {
		if entry.IsDir() {
			subDir, _err := chopDir(dirPath + "/" + entry.Name())
			if _err != nil {
				return utils.ManifestDir{}, _err
			}
			manifestDir.SubDir = append(manifestDir.SubDir, subDir)

		} else {
			file, _err := chopFile(dirPath, entry.Name())
			if _err != nil {
				return utils.ManifestDir{}, _err
			}
			manifestDir.SubFiles = append(manifestDir.SubFiles, file)
		}
	}

	fmt.Println("Chop dir", manifestDir.Name, "done") // ? Here for debug purpose ( replace by docLogger later )

	return manifestDir, err
}

func createDir(_path string) (_err error) {
	if _err = os.MkdirAll("chop/", 0700); _err != nil {
		return _err
	}
	if _err = os.MkdirAll("chop/"+_path+"/", 0700); _err != nil {
		return _err
	}
	return _err
}
