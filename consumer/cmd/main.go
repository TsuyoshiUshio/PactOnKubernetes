package main

import "github.com/TsuyoshiUshio/PactOnKubernetes/consumer"

func main() {
	client := consumer.Client{
		Host: "http://localhost:9000",
	}
	client.Run()
}
