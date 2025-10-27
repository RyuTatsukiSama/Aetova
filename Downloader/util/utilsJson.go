package util

import (
	"encoding/json"
	"os"
)

func FromJson[T any](_jsonData *T, _jsonFile *os.File) (_err error) {
	decoder := json.NewDecoder(_jsonFile)
	_err = decoder.Decode(_jsonData)
	if _err != nil {
		return _err
	}

	return _err
}

func ToJson[T any](_jsonData T, _jsonFile *os.File) (_err error) {
	encoder := json.NewEncoder(_jsonFile)
	encoder.SetIndent("", "  ")
	_err = encoder.Encode(_jsonData)
	if _err != nil {
		return _err
	}
	return _err
}
