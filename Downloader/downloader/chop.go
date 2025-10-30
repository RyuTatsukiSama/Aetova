package downloader

import (
	"Aetova/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	sizeChunk int = 4096 // bytes
)

func ChopGame(zip string) (err error) {
	start := time.Now()
	// unzip the game send by the dev
	var unzipPath string
	if unzipPath, err = Unzip(zip); err != nil {
		return err
	}

	fmt.Println("Unzip done !")

	// create the manifest.json for the game
	var manifestDir util.ManifestDir
	manifestDir.Name = strings.Split(unzipPath, "/")[len(strings.Split(unzipPath, "/"))-1]
	if manifestDir, err = chopDir(unzipPath); err != nil {
		return err
	}

	// create the manifest file
	var file *os.File
	if file, err = os.Create("manifest.json"); err != nil {
		return err
	}

	// save the data into the file
	if err = util.ToJson(manifestDir, file); err != nil {
		return err
	}

	fmt.Println("The process takes", time.Since(start))

	return err
}

func chopDir(dirPath string) (manifestDir util.ManifestDir, err error) {
	// get the data of the directory
	dir, _err := os.ReadDir(dirPath)
	if _err != nil {
		return util.ManifestDir{}, _err
	}

	manifestDir.Name = strings.Split(dirPath, "/")[len(strings.Split(dirPath, "/"))-1]
	createDir(dirPath)

	fmt.Println("Chop dir", manifestDir.Name, "start") // ? Here for debug purpose ( replace by docLogger later )

	// browse file & subDir
	for _, entry := range dir {
		if entry.IsDir() {
			subDir, _err := chopDir(dirPath + "/" + entry.Name())
			if _err != nil {
				return util.ManifestDir{}, _err
			}
			manifestDir.SubDir = append(manifestDir.SubDir, subDir)

		} else {
			file, _err := chopFile(dirPath, entry.Name())
			if _err != nil {
				return util.ManifestDir{}, _err
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

func chopFile(_path string, _name string) (util.ManifestFile, error) {

	fmt.Println("Chop file", _name, "start") // ? Here for debug purpose ( replace by docLogger later )

	var allPath string = _path + "/" + _name
	// open the file
	file, err := os.Open(allPath)
	if err != nil {
		return util.ManifestFile{}, err
	}
	defer file.Close()

	// get the info for size
	info, err := os.Stat(file.Name())
	if err != nil {
		return util.ManifestFile{}, err
	}

	// put all the bytes into a slice
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		return util.ManifestFile{}, err
	}

	manifestFile := util.ManifestFile{
		Name:       _name,
		IsDownload: false,
	}

	// cut the data
	previousI := 0
	for i := sizeChunk; i < len(data); i += sizeChunk {
		current := data[previousI:i]
		res, err := saveChunk(current, _path, _name, i/sizeChunk)
		if err != nil {
			return util.ManifestFile{}, err
		}
		manifestFile.Chunks = append(manifestFile.Chunks, util.Chunk{
			Name:       res,
			IsDownload: false,
		})
		previousI = i
	}

	// last part
	current := data[len(data)/sizeChunk*sizeChunk:]
	res, err := saveChunk(current, _path, _name, len(data)/sizeChunk+1)
	if err != nil {
		return util.ManifestFile{}, err
	}
	manifestFile.Chunks = append(manifestFile.Chunks, util.Chunk{
		Name:       res,
		IsDownload: false,
	})

	fmt.Println("Chop file", _name, "done") // ? Here for debug purpose ( replace by docLogger later )

	return manifestFile, nil
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
