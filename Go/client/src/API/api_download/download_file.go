package api_download

import (
	"aetova/client/utils"
)

func downloadFile(file utils.ManifestFile, path string, cd chan bool, ce chan error) {

	chanDone := make(chan bool)
	chanError := make(chan error)

	var part int = 0

	go downloadChunk(path, file.Name, part, chanDone, chanError)

	for range file.NbChunks {
		select {
		case <-Ctx.Done():
			return
		case err := <-chanError:
			ce <- err
			return
		case <-chanDone:
			if part < file.NbChunks {
				part++
				go downloadChunk(path, file.Name, part, chanDone, chanError)
			}
		}
	}

	cd <- true
}
