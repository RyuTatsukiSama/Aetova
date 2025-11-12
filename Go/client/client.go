package main

import (
	client_api "aetova/client/src/API"
	"log"
	"os"
)

func main() {
	if err := os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	log.Fatal(client_api.LaunchAPI())
}
