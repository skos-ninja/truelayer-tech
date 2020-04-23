package pokeapi

import (
	"context"
	"net/http"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

// For the purposes of this test we have implemented the calls to the pokeapi ourselves
// however it's worth noting that the pokeapi lists wrappers for their library available
// of which one includes https://github.com/mtslzr/pokeapi-go for a go wrapper.

const baseURL = "https://pokeapi.co"

type PokeAPI interface {
	GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error)
}

type pokeAPI struct {
	client *http.Client
}

func New(cacheSize int) PokeAPI {
	return &pokeAPI{
		client: http.DefaultClient,
	}
}
