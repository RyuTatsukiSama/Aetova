package butcher

import (
	"aetova/server/utils"
	"fmt"
	"os"
)

func chopFile(_path string, _name string) (utils.ManifestFile, error) {

	fmt.Println("Chop file", _name, "start") // ? Here for debug purpose ( replace by docLogger later )

	var allPath string = _path + "/" + _name
	// open the file
	file, err := os.Open(allPath)
	if err != nil {
		return utils.ManifestFile{}, err
	}
	defer file.Close()

	// get the info for size
	info, err := os.Stat(file.Name())
	if err != nil {
		return utils.ManifestFile{}, err
	}

	// put all the bytes into a slice
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		return utils.ManifestFile{}, err
	}

	manifestFile := utils.ManifestFile{
		Name:     _name,
		NbChunks: 0,
	}

	// cut the data
	previousI := 0
	for i := sizeChunk; i < len(data); i += sizeChunk {
		current := data[previousI:i]
		_, err := saveChunk(current, _path, _name, i/sizeChunk)
		if err != nil {
			return utils.ManifestFile{}, err
		}
		manifestFile.NbChunks++

		previousI = i
	}

	// last part
	current := data[len(data)/sizeChunk*sizeChunk:]
	_, err = saveChunk(current, _path, _name, len(data)/sizeChunk+1)
	if err != nil {
		return utils.ManifestFile{}, err
	}
	manifestFile.NbChunks++

	fmt.Println("Chop file", _name, "done") // ? Here for debug purpose ( replace by docLogger later )

	return manifestFile, nil
}
