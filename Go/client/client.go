package main

import (
	client_api "aetova/client/src/API"
	"log"
)

func main() {
	log.Fatal(client_api.LaunchAPI())
}
