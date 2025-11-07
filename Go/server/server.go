package main

import (
	server_api "aetova/server/src/API"
	butcher "aetova/server/src/Butcher"
	"log"
	"os"
)

func main() {
	var err error = nil

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	if err = butcher.ChopGame("BuildOranys.zip"); err != nil {
		log.Fatal(err)
	}
	log.Fatal(server_api.LaunchAPI())
}
