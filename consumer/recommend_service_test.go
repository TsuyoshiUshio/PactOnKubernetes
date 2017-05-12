package consumer

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	"net/http"

	"net/http/httptest"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

// test data
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var pact dsl.Pact
var pactBrokerURL = os.Getenv("PACT_BROKER_URL")
var searchRequest = ` { "keyword":"protein"}`
var form url.Values
var rr http.ResponseWriter
var req *http.Request

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	pact.WritePact()

	publishRequest := types.PublishRequest{
		PactBroker:      pactBrokerURL,
		PactURLs:        []string{"../pacts/recommendation-product.json"},
		ConsumerVersion: "1.0.0",
		Tags:            []string{"latest", "dev"},
	}
	p := dsl.Publisher{}
	err := p.Publish(publishRequest)
	if err != nil {
		log.Fatal(err)
	}

	pact.Teardown()
	os.Exit(code)
}

func setup() {
	pact = createPact()

	form = url.Values{}
	form.Add("keyword", "protein")

	req, _ = http.NewRequest("POST", "/search", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = form

	rr = httptest.NewRecorder()
}

func createPact() dsl.Pact {
	pactDaemonPort := 6666
	return dsl.Pact{
		Port:     pactDaemonPort,
		Consumer: "recommendation",
		Provider: "product",
		LogDir:   logDir,
		PactDir:  pactDir,
	}
}

func TestPactConsumerSearchHandler_ProductExists(t *testing.T) {
	var testProteinExists = func() error {
		client := Client{
			Host: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
		}
		client.searchHandler(rr, req)
		if client.Product == nil {
			return errors.New("Expected product not to be nil")
		}
		return nil
	}
	pact.
		AddInteraction().
		Given("Product protein exists").
		UponReceiving("A request to search with keyword 'protein'").
		WithRequest(dsl.Request{
			Method: "POST",
			Path:   "/products/search",
			Body:   searchRequest,
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Headers: map[string]string{
				"Content-Type": "application/json; charset=utf-8",
			},
			Body: `
				{
					"product": {
						"Id":    1,
						"Name":  "Protein",
						"Price": 40
					}
				}
				`,
		})

	err := pact.Verify(testProteinExists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}
