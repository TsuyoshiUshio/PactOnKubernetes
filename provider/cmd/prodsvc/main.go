package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/TsuyoshiUshio/PactOnKubernetes/provider"
)

func main() {
	mux := http.NewServeMux()
	var port = 9000
	mux.HandleFunc("/products/search", provider.ProductSearch)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))

}
