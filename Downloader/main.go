package main

import (
	"Aetova/downloader"
	"Aetova/util"
	"log"
	"os"
)

func main() {
	var err error = nil

	err = os.Chdir("WorkingDirectory")

	checkErr(err)

	err = downloader.Unzip("docLogger_v1-1-1.zip")

	checkErr(err)

	manifest, err := downloader.ChopGame("unzip/docLogger_v1-1-1")
	checkErr(err)

	// for manifest.json
	jsonFile, err := os.Create("manifest.json")
	checkErr(err)

	err = util.ToJson(manifest, jsonFile)
	checkErr(err)

	// err = downloader.Assemble()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
