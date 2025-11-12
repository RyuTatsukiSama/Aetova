package api_download

import (
	"aetova/client/utils"
	"fmt"
	"strconv"
)

func downloadFile(file utils.ManifestFile, path string) {
	for part := 1; part < file.NbChunks; part++ {
		currentPath := path + "part" + strconv.Itoa(part) + "_" + file.Name + ".bin"

		fmt.Println("Download file " + currentPath + " Start")

		saveChunk(currentPath, downloadChunk(currentPath))

		fmt.Println("Download file " + currentPath + " Done")
	}
}
