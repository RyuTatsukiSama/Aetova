package server

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func Unzip(path string, guid int) (unzipPath string, err error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var newFilePath string
	for _, file := range reader.File {

		// open the stream link to the file/dir
		var rc io.ReadCloser
		if rc, err = file.Open(); err != nil {
			return "", err
		}
		defer rc.Close()

		// define the new file path
		newFilePath = fmt.Sprintf("unzip/%d/%s", guid, file.Name)

		if file.FileInfo().IsDir() {
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				return "", err
			}
		} else {
			unzipFile, err := os.Create(newFilePath)
			if err != nil {
				return "", err
			}
			if _, err = io.Copy(unzipFile, rc); err != nil {
				return "", err
			}
		}
	}

	return fmt.Sprintf("%d/%s", guid, strings.Split(path, ".")[0]), err
}
