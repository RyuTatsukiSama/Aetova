package api_download

import (
	"aetova/client/utils"
	"fmt"
	"strconv"
)

func downloadFile(file utils.ManifestFile, path string) error {
	for part := 0; part < file.NbChunks; part++ {
		currentChunkPath := path + "part" + strconv.Itoa(part) + "_" + file.Name + ".bin"

		fmt.Println("Download file " + currentChunkPath + " Start")

		data, err := downloadChunk(currentChunkPath)
		if err != nil {
			return err
		}

		saveChunk(path+file.Name, data, int64(part*utils.SizeChunk))

		fmt.Println("Download file " + currentChunkPath + " Done")
	}

	return nil
}
