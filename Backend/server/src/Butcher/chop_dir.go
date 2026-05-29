package butcher

import (
	"aetova/server/utils"
	"fmt"
	"os"
	"strings"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

const (
	FromDir string = "unzip/"
	ToDir   string = "data/"
)

func chopDir(dirPath string) (manifestDir utils.ManifestDir, err error) {
	dLog := docLogger.NewLoggerWithGOpts("Server/ChopDir")

	// get the data of the directory
	dir, _err := os.ReadDir(FromDir + dirPath)
	if _err != nil {
		return utils.ManifestDir{}, _err
	}

	manifestDir.Name = strings.Split(dirPath, "/")[len(strings.Split(dirPath, "/"))-1]
	_err = createDir(dirPath)
	if _err != nil {
		return utils.ManifestDir{}, _err
	}

	dLog.Debug(fmt.Sprintf("Chopdir %s start", manifestDir.Name))

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

	dLog.Debug(fmt.Sprintf("Chopdir %s done", manifestDir.Name))

	return manifestDir, err
}

func createDir(_path string) (_err error) {
	if _err = os.MkdirAll(ToDir, 0700); _err != nil {
		return _err
	}
	if _err = os.MkdirAll(ToDir+_path+"/", 0700); _err != nil {
		return _err
	}
	return _err
}
