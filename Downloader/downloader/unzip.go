package downloader

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func Unzip(_path string) (_err error) {
	reader, _err := zip.OpenReader(_path)
	if _err != nil {
		return _err
	}
	defer reader.Close()

	for _, file := range reader.File {
		fmt.Printf("Unzipping %s\n", file.Name)

		// open the stream link to the file/dir
		rc, _err := file.Open()
		if _err != nil {
			return _err
		}
		defer rc.Close()

		// define the new file path
		newFilePath := fmt.Sprintf("uncompressed/%s", file.Name)

		if file.FileInfo().IsDir() {
			_err = os.MkdirAll(newFilePath, 0777)
			if _err != nil {
				return _err
			}
		} else {
			unzipFile, _err := os.Create(newFilePath)
			if _err != nil {
				return _err
			}
			_, _err = io.Copy(unzipFile, rc)
			if _err != nil {
				return _err
			}
		}
	}

	return _err
}
