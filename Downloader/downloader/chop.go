package downloader

import (
	"Aetova/util"
	"os"
	"strconv"
)

func ChopGame(_path string) (_err error) {
	dir, _err := os.ReadDir(_path)
	if _err != nil {
		return _err
	}

	for _, entry := range dir {
		if entry.IsDir() {

		} else {

		}
	}

	return _err
}

func createTree() (_err error) {
	return _err
}

func Chop(_path string) error {

	// open the .zip file
	file, err := os.Open(_path)
	if err != nil {
		return err
	}
	defer file.Close()

	// get the info for size
	info, err := os.Stat(file.Name())
	if err != nil {
		return err
	}

	// put all the bytes into a slice
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	if err := os.MkdirAll("chop/", 0700); err != nil {
		return err
	}

	jsonData := []util.ManifestFile{}

	// cut the data
	previousI := 0
	for i := 1000; i < len(data); i += 1000 {
		current := data[previousI:i]
		res, err := saveChopFile(current, i/1000)
		if err != nil {
			return err
		}
		jsonData = append(jsonData, util.ManifestFile{
			Name:       res,
			IsDownload: false,
		})
		previousI = i
	}

	// last part
	current := data[len(data)/1000*1000:]
	res, err := saveChopFile(current, len(data)/1000+1)
	if err != nil {
		return err
	}
	jsonData = append(jsonData, util.ManifestFile{
		Name:       res,
		IsDownload: false,
	})

	// for manifest.json
	jsonFile, err := os.Create("manifest.json")
	if err != nil {
		return err
	}

	err = util.ToJson(jsonData, jsonFile)
	if err != nil {
		return err
	}

	return nil
}

func saveChopFile(_data []byte, _part int) (_rpath string, _err error) {
	_rpath = "chop/part" + strconv.Itoa(_part) + "_" + "docLogger.bin"
	if ok, err := util.Exists(_rpath); !ok {
		// create the file
		file, err := os.Create(_rpath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		// write the data in it
		_, err = file.Write(_data)
		if err != nil {
			return "", err
		}

		// return it name
		return _rpath, err
	} else if err != nil {
		return "", err // return an error
	} else {
		return _rpath, nil // return if the fil already exist
	}
}
