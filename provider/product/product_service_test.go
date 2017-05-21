package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"testing"

	ptypes "github.com/TsuyoshiUshio/PactOnKubernetes/types"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

func TestPact_Provider(t *testing.T) {
	go startInstrumentProvider()
	pact := createPact()

	var pactBrokerURL = os.Getenv("PACT_BROKER_URL")

	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs:        []string{fmt.Sprintf("%s/pacts/provider/product/consumer/recommendation/latest/dev", pactBrokerURL)},
		// PactURLs:               []string{fmt.Sprintf("%s/recommendation-product.json", pactDir)},
		ProviderStatesURL:      fmt.Sprintf("http://localhost:%d/states", port),
		ProviderStatesSetupURL: fmt.Sprintf("http://localhost:%d/setup", port),
	})

	if err != nil {
		t.Fatal("Error", err)
	}
}

func startInstrumentProvider() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products/search", ProductSearch)
	mux.HandleFunc("/setup", providerStateSetupFunc)
	mux.HandleFunc("/states", providerStatesFunc)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))
}

var providerStates = map[string][]string{
	"recommendation": []string{
		"Product protein exists"},
}

var proteinExists = &ptypes.ProductRepository{
	Product: map[string]*ptypes.Product{
		"protein": &ptypes.Product{
			Id:    1,
			Name:  "Protein",
			Price: 40,
		},
	},
}

var providerStatesFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := json.Marshal(providerStates)
	w.Write(body)
}

var providerStateSetupFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var state types.ProviderState

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	err = json.Unmarshal(body, &state)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Setup database for different states
	if state.State == "Product protein exists" {
		productRepository = proteinExists
	}
}

var dir, _ = os.Getwd()

var pactDir = fmt.Sprintf("%s/../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

func createPact() dsl.Pact {
	return dsl.Pact{
		Port:     6666,
		Consumer: "recommendation",
		Provider: "product",
	}
}
