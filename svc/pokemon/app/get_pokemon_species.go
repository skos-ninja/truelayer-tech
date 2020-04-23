package app

import (
	"context"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

// pokeKey is a custom string type to ensure no collisions
type pokeKey string

func (a *app) GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error) {
	key := pokeKey(pokemon)
	species, ok := a.pokeLRU.Get(key)
	if !ok {
		var err error
		species, err = a.pokeAPI.GetPokemonSpecies(ctx, pokemon)
		if err != nil {
			return nil, err
		}

		a.pokeLRU.Add(key, species)
	}

	return species.(*models.PokemonSpecies), nil
}
