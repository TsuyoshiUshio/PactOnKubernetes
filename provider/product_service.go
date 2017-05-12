package provider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TsuyoshiUshio/PactOnKubernetes/types"
)

var productRepository = &types.ProductRepository{
	Product: map[string]*types.Product{
		"protein": &types.Product{
			Id:    1,
			Name:  "Protein",
			Price: 40,
		},
		"fatburn": &types.Product{
			Id:    2,
			Name:  "Fat Burn",
			Price: 39,
		},
		"teststrone": &types.Product{
			Id:    3,
			Name:  "Teststrone Booster",
			Price: 110,
		},
	},
}

// ProductSearch is the sarch by keyword
var ProductSearch = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	log.Printf("ProductSearch:")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var request types.SearchRequest

	err = json.Unmarshal(body, &request)
	log.Printf("ProductSearch: Keyword: %s", request.Keyword)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	product, err := productRepository.ByKeyword(request.Keyword)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		res := types.SearchResponse{Product: product}
		resBody, _ := json.Marshal(res)
		log.Printf("*****************")
		log.Printf(string(resBody))
		w.Write(resBody)
	}

}
