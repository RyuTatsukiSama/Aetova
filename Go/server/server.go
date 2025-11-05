package main

import (
	server "aetova/server/src"
	"log"
	"os"
)

func main() {
	var err error = nil

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	if err = server.ChopGame("BuildOranys.zip"); err != nil {
		log.Fatal(err)
	}
}
