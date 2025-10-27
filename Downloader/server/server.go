package main

import (
	"Aetova/downloader"
	"Aetova/util"
	"os"
)

func main() {
	var err error = nil

	err = os.Chdir("wd")
	util.CheckErr(err)

	err = downloader.Unzip("docLogger_v1-1-1.zip")
	util.CheckErr(err)

	manifest, err := downloader.ChopGame("unzip/docLogger_v1-1-1")
	util.CheckErr(err)

	// for manifest.json
	jsonFile, err := os.Create("manifest.json")
	util.CheckErr(err)

	err = util.ToJson(manifest, jsonFile)
	util.CheckErr(err)
}
