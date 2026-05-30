package main

import (
	server_api "aetova/server/src/API"
	butcher "aetova/server/src/Butcher"
	update "aetova/server/src/Update"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func main() {
	opts := docLogger.NewOptionsBuilder().Build()
	docLogger.SetGlobalLoggerOptions(opts)
	dLog := docLogger.NewLoggerWithGOpts("Server")

	err := createDirectory()
	if err != nil {
		dLog.Critical(err.Error())
		return
	}

	guid := uint(0) // TODO: Generate it with postrgre cf https://github.com/RyuTatsukiSama/Aetova/issues/33
	version := uint(0)
	gname := "BuildOranys"
	err = butcher.RegisterGameManifest(gname, guid, version)
	if err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	if err = butcher.ChopGame("v1_0_0.zip", gname, guid, version); err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	err = update.UpdateGame("v1_0_1.zip", guid)
	if err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	dLog.Critical(server_api.LaunchAPI().Error())
}

func createDirectory() error {
	err := os.MkdirAll("Shipyard", 0700)
	if err != nil {
		return err
	}

	err = os.Chdir("Shipyard")
	if err != nil {
		return err
	}

	err = os.MkdirAll("unzip", 0700)
	if err != nil {
		return err
	}

	err = os.MkdirAll("data", 0700)
	if err != nil {
		return err
	}

	return nil
}
