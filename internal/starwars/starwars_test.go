package starwars_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	golangtraining "github.com/julianjca/julian-golang-training-beginner"
	"github.com/julianjca/julian-golang-training-beginner/internal/starwars"
	"github.com/stretchr/testify/require"
)

func TestGetCharacters(t *testing.T) {
	expectedRes := &golangtraining.StarWarsResponse{
		Results: []golangtraining.Characters{
			{
				Name: "Luke Skywalker",
			},
		},
	}

	jsonResp, err := json.Marshal(expectedRes)
	require.NoError(t, err)

	handler := func() (res http.Handler) {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "https://swapi.dev/api/people", r.URL.String())
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(jsonResp)
			require.NoError(t, err)
		})
	}()
	mockServer := httptest.NewServer(handler)
	defer mockServer.Close()

	starwarsClient := starwars.NewStarWarsClient(http.DefaultClient)
	res, err := starwarsClient.GetCharacters()
	require.NoError(t, err)
	require.Equal(t, expectedRes.Results[0], res.Results[0])
}
