package downloader

import (
	"Aetova/util"
	"encoding/json"
	"os"
)

func Assemble() error {
	// open the json
	var jsonData []util.Manifest
	file, err := os.Open("manifest.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// decode the json
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return err
	}

	allData := []byte{}
	for i := 0; i < len(jsonData); i++ {
		currentData, err := loadFile(jsonData[i].Name)
		if err != nil {
			return err
		}

		allData = append(allData, currentData...)
	}

	path := "newdocLogger.zip"
	if ok, err := util.Exists(path); !ok {
		// create the file
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// write the data in it
		_, err = file.Write(allData)
		if err != nil {
			return err
		}

		return err
	} else if err != nil {
		return err // return an error
	}

	return nil
}

func loadFile(_path string) (_data []byte, err error) {
	// open the .zip file
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
