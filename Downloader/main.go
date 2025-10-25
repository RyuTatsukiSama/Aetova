package main

import (
	"log"
	"os"
)

func main() {
	var err error = nil

	err = os.Chdir("WorkingDirectory")

	checkErr(err)

	// err = downloader.Unzip("docLogger_v1-1-1.zip")

	checkErr(err)

	// err = downloader.ChopGame("uncompressed/docLogger_v1-1-1")

	checkErr(err)

	// err = downloader.Chop("docLogger_v1-1-1.zip")

	checkErr(err)

	// err = downloader.Assemble()

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
