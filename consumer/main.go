package main

import "github.com/TsuyoshiUshio/PactOnKubernetes/consumer/goconsumer"

func main() {
	client := goconsumer.Client{
		Host: "http://localhost:9000",
	}
	client.Run()
}
