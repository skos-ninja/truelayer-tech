package app

import (
	"context"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

func (a *app) GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error) {
	// TODO: Add an lru cache here
	return a.pokeAPI.GetPokemonSpecies(ctx, pokemon)
}
