package downloader

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func Unzip(path string) (unzipPath string, err error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var newFilePath string
	for _, file := range reader.File {
		fmt.Printf("Unzipping %s\n", file.Name)

		// open the stream link to the file/dir
		var rc io.ReadCloser
		if rc, err = file.Open(); err != nil {
			return "", err
		}
		defer rc.Close()

		// define the new file path
		newFilePath = fmt.Sprintf("unzip/%s", file.Name)

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

	return newFilePath, err
}
