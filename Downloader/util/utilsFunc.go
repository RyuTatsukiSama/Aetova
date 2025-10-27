package util

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
