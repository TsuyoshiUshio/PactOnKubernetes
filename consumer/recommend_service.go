package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type searchResponse struct {
	Product Product `json:"product"`
}

type Product struct {
	ID    int64
	Name  string `json:"name"`
	Price int64
}

type Client struct {
	Product *Product
	Host    string
	Err     error
}
type templateData struct {
	Product *Product
	Err     error
}

var searchTemplatePath = "search.html"
var templates = template.Must(template.ParseFiles(searchTemplatePath))

func renderTemplate(w http.ResponseWriter, tmpl string, u templateData) {
	err := templates.ExecuteTemplate(w, tmpl+".html", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Client) searchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	searchRequest := fmt.Sprintf(`
	{
		"keyword":"%s"
}`, keyword)

	res, err := http.Post(fmt.Sprintf("%s/products/search", c.Host), "application/json", bytes.NewReader([]byte(searchRequest)))
	if res.StatusCode != 200 || err != nil {
		c.Err = fmt.Errorf("Search failed")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Err = fmt.Errorf("Search failed")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var response searchResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		c.Err = fmt.Errorf("Search failed")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	result := templateData{
		Product: &response.Product,
		Err:     c.Err,
	}
	renderTemplate(w, "search", result)
}

func (c *Client) viewHandler(w http.ResponseWriter, r *http.Request) {
	data := templateData{
		Product: c.Product,
		Err:     c.Err,
	}
	renderTemplate(w, "search", data)

}

func (c *Client) Run() {
	http.HandleFunc("/search", c.searchHandler)
	http.HandleFunc("/", c.viewHandler)
	fmt.Println("Product svc client runnning on port 8081")
	http.ListenAndServe(":8081", nil)
}
