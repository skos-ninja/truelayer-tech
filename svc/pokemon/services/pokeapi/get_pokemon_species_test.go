package pokeapi

import (
	"context"
	"net/http"
	"testing"

	"github.com/skos-ninja/truelayer-tech/lib/http/test"

	"github.com/stretchr/testify/assert"
)

const (
	okResponse = `{
		"name": "charizard"
	}`

	notFoundResponse = `Not Found`
)

func TestGetPokemonSpeciesValidPokemon(t *testing.T) {
	ctx := context.Background()
	client, close := test.CreateTestHTTPClient(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})
	defer close()

	svc := New()
	svc.(*service).client = client

	species, err := svc.GetPokemonSpecies(ctx, "charizard")
	assert.Nil(t, err)
	assert.Equal(t, "charizard", species.Name)
}

func TestGetPokemonSpeciesInvalidPokemon(t *testing.T) {
	ctx := context.Background()
	client, close := test.CreateTestHTTPClient(func(w http.ResponseWriter, r *http.Request) {
		// Return a 404 error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(notFoundResponse))
	})
	defer close()

	svc := New()
	svc.(*service).client = client

	species, err := svc.GetPokemonSpecies(ctx, "invalid-pokemon")
	assert.Nil(t, species)
	assert.Equal(t, ErrSpeciesNotFound, err, err.Error())
}
