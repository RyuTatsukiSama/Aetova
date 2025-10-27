package downloader

import (
	"Aetova/util"
	"os"
	"strconv"
	"strings"
)

const (
	sizeChunk int = 1000 // 1ko
)

func ChopDir(_path string) (_maDir util.ManifestDir, _err error) {
	dir, _err := os.ReadDir(_path)
	if _err != nil {
		return util.ManifestDir{}, _err
	}

	_maDir.Name = strings.Split(_path, "/")[len(strings.Split(_path, "/"))-1]
	saveDir(_path)

	// browse file & subDir
	for _, entry := range dir {
		if entry.IsDir() {
			subDir, _err := ChopDir(_path + "/" + entry.Name())
			if _err != nil {
				return util.ManifestDir{}, _err
			}
			_maDir.SubDir = append(_maDir.SubDir, subDir)

		} else {
			file, _err := ChopFile(_path, entry.Name())
			if _err != nil {
				return util.ManifestDir{}, _err
			}
			_maDir.SubFiles = append(_maDir.SubFiles, file...)
		}
	}

	return _maDir, _err
}

func saveDir(_path string) (_err error) {
	if _err = os.MkdirAll("chop/", 0700); _err != nil {
		return _err
	}
	if _err = os.MkdirAll("chop/"+_path+"/", 0700); _err != nil {
		return _err
	}
	return _err
}

func ChopFile(_path string, _name string) ([]util.ManifestFile, error) {

	var allPath string = _path + "/" + _name
	// open the .zip file
	file, err := os.Open(allPath)
	if err != nil {
		return []util.ManifestFile{}, err
	}
	defer file.Close()

	// get the info for size
	info, err := os.Stat(file.Name())
	if err != nil {
		return []util.ManifestFile{}, err
	}

	// put all the bytes into a slice
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		return []util.ManifestFile{}, err
	}

	jsonData := []util.ManifestFile{}

	// cut the data
	previousI := 0
	for i := sizeChunk; i < len(data); i += sizeChunk {
		current := data[previousI:i]
		res, err := saveChunk(current, _path, _name, i/sizeChunk)
		if err != nil {
			return []util.ManifestFile{}, err
		}
		jsonData = append(jsonData, util.ManifestFile{
			Name:       res,
			IsDownload: false,
		})
		previousI = i
	}

	// last part
	current := data[len(data)/sizeChunk*sizeChunk:]
	res, err := saveChunk(current, _path, _name, len(data)/sizeChunk+1)
	if err != nil {
		return []util.ManifestFile{}, err
	}
	jsonData = append(jsonData, util.ManifestFile{
		Name:       res,
		IsDownload: false,
	})

	return jsonData, nil
}

func saveChunk(_data []byte, _path string, _name string, _part int) (_rname string, _err error) {
	_rname = "part" + strconv.Itoa(_part) + "_" + _name + ".bin"
	path := "chop/" + _path + "/" + _rname // TODO : found a better naming way
	if ok, err := util.Exists(path); !ok {
		// create the file
		file, err := os.Create(path)
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
		return _rname, err
	} else if err != nil {
		return "", err // return an error
	} else {
		return _rname, nil // return if the file already exist
	}
}
