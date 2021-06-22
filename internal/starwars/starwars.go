package starwars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewStarWarsClient(httpClient *http.Client) *Client {
	return &Client{
		BaseURL:    "https://swapi.dev/api",
		HTTPClient: httpClient,
	}
}

func (r *Client) GetCharacters() (resp *http.Response, err error) {
	u, err := url.Parse(r.BaseURL)
	req, err := http.NewRequest(http.MethodGet, u.Path, nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return
	}

	u.Path = path.Join(u.Path, fmt.Sprintf("/people"))

	resp, err = r.HTTPClient.Do(req)
	if err != nil {
		return
	}

	jbyt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var res []golangtraining.StarWarsResponse
	err = json.Unmarshal(jbyt, &res)
	if err != nil {
		return
	}

	return
}
