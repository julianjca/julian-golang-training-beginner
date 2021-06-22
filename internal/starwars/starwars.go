package starwars

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewStarWarsClient(httpClient *http.Client) *Client {
	return &Client{
		BaseURL:    "https://swapi.dev/api/people",
		HTTPClient: httpClient,
	}
}

func (r *Client) GetCharacters() (*golangtraining.StarWarsResponse, error) {
	u, err := url.Parse(r.BaseURL)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	jbyt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res *golangtraining.StarWarsResponse
	err = json.Unmarshal(jbyt, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
