package main

import (
	server "aetova/server/src"
	"log"
	"os"
)

func main() {
	var err error = nil

	if err = os.MkdirAll("wd", 0700); err != nil {
		log.Fatal(err)
	}

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	if err = server.ChopGame("BuildOranys.zip"); err != nil {
		log.Fatal(err)
	}
	// log.Fatal(server_api.LaunchAPI())
}
