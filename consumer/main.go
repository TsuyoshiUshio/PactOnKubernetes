package main

import (
	"fmt"
	"os"

	"github.com/TsuyoshiUshio/PactOnKubernetes/consumer/goconsumer"
)

func main() {
	var providerHost = os.Getenv("PROVIDER_HOST")
	if providerHost == "" {
		providerHost = "localhost"
	}
	client := goconsumer.Client{
		Host: fmt.Sprintf("http://%s:9000", providerHost),
	}
	client.Run()
}
