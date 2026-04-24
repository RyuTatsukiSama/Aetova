package butcher

import (
	"aetova/server/utils"
	"fmt"
	"os"
)

func saveChunk(_data []byte, _path string, _name string, _part int64) (_rname string, _err error) {

	_rname = fmt.Sprintf("part%d_%s.chk", _part, _name)

	path := "chop/" + _path + "/" + _rname // TODO : found a better naming way

	if ok, err := utils.Exists(path); !ok {
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
