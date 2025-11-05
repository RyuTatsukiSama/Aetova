package main

import (
	server "aetova/server/src"
	server_api "aetova/server/src/API"
	"log"
	"os"
)

func main() {
	var err error = nil

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}
	log.Fatal(server_api.LaunchAPI())

	if err = server.ChopGame("BuildOranys.zip"); err != nil {
		log.Fatal(err)
	}
}
