package assemble

import (
	"aetova/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	sourcePath string = "chop/unzip"
	targetPath string = "assemble"
)

func AssembleGame(_manifest string) (err error) {
	start := time.Now()
	// open the json
	var jsonData utils.ManifestDir
	if err = readJson(_manifest, &jsonData); err != nil {
		return err
	}

	// assemble the binary into one file
	if err = assembleDir(jsonData, "/"); err != nil {
		return err
	}

	fmt.Println("The process takes", time.Since(start))
	return err
}

func readJson(manifest string, jsonData *utils.ManifestDir) (err error) {
	file, err := os.Open(manifest)
	if err != nil {
		return err
	}
	defer file.Close()

	// decode the json
	err = utils.FromJson(jsonData, file)
	if err != nil {
		return err
	}

	return err
}

func assembleDir(manifestDir utils.ManifestDir, path string) (err error) {
	fmt.Println("Start assemble", manifestDir.Name)

	// create the parent directory
	if err = os.MkdirAll(targetPath+path+manifestDir.Name+"/", 0700); err != nil {
		return err
	}

	// create the sub directory
	for _, dir := range manifestDir.SubDir {
		if err = assembleDir(dir, path+manifestDir.Name+"/"); err != nil {
			return err
		}
	}

	// create the file
	for _, file := range manifestDir.SubFiles {

		if err = assembleFile(file, path+manifestDir.Name); err != nil {
			return err
		}
	}

	fmt.Println("Done assemble", manifestDir.Name)

	return err
}

func assembleFile(file utils.ManifestFile, path string) (err error) {
	fmt.Println("Start assemble", file.Name)
	// assemble chunk
	var data []byte
	for idChunk := 1; idChunk < file.NbChunks; idChunk++ {
		var chunkData []byte
		if chunkData, err = loadChunk(sourcePath + path + "/" + "part" + strconv.Itoa(idChunk) + "_" + file.Name + ".bin"); err != nil {
			return err
		}
		data = append(data, chunkData...)
	}

	// create new file
	var newFile *os.File
	if newFile, err = os.Create(targetPath + path + "/" + file.Name); err != nil {
		return err
	}

	// write the byte in the new file
	if _, err = newFile.Write(data); err != nil {
		return err
	}

	fmt.Println("Done assemble", file.Name)

	return err
}

func loadChunk(_path string) (_data []byte, err error) {
	// open the chunk file
	file, err := os.Open(_path)
	if err != nil {
		return _data, err
	}
	defer file.Close()

	// get the info for size
	info, err := os.Stat(file.Name())
	if err != nil {
		return _data, err
	}

	// put all the bytes into a slice
	_data = make([]byte, info.Size())
	_, err = file.Read(_data)
	if err != nil {
		return _data, err
	}

	return _data, err
}
