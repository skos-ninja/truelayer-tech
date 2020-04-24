package app

import (
	"context"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare"

	"github.com/hashicorp/golang-lru/simplelru"
)

type App interface {
	GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error)
	GetShakespeareText(ctx context.Context, text string) (string, error)

	GetShakespearePokemonDescription(ctx context.Context, pokemon string) (string, error)
}

type app struct {
	pokeAPI pokeapi.Service
	pokeLRU *simplelru.LRU

	shakespeare    shakespeare.Service
	shakespeareLRU *simplelru.LRU
}

func New(cacheSize int) (App, error) {
	pLRU, err := simplelru.NewLRU(cacheSize, nil)
	if err != nil {
		return nil, err
	}

	sLRU, err := simplelru.NewLRU(cacheSize, nil)
	if err != nil {
		return nil, err
	}

	return &app{
		pokeAPI: pokeapi.New(),
		pokeLRU: pLRU,

		shakespeare:    shakespeare.New(),
		shakespeareLRU: sLRU,
	}, nil
}
