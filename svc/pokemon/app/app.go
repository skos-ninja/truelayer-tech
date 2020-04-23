package app

import (
	"context"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

type App interface {
	GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error)
	GetShakespearText(ctx context.Context, text string) (string, error)

	GetShakespearPokemonDescription(ctx context.Context, pokemon string) (string, error)
}

type app struct {
	pokeAPI pokeapi.PokeAPI
}

func New() App {
	return &app{
		pokeAPI: pokeapi.New(128),
	}
}
