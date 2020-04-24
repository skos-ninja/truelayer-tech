package test

import (
	"context"
	"fmt"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

type testService struct {
	success bool
}

func (t *testService) GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error) {
	if !t.success {
		return nil, pokeapi.ErrSpeciesNotFound
	}

	return &models.PokemonSpecies{
		Name: pokemon,
		FlavorTextEntries: []models.FlavorText{
			{
				FlavorText: fmt.Sprint(pokemon, " is cool"),
				Language: models.NamedAPIResource{
					Name: "en",
					URL:  "",
				},
				Version: models.NamedAPIResource{
					Name: "25",
					URL:  "",
				},
			},
		},
	}, nil
}

func New(success bool) pokeapi.Service {
	return &testService{success}
}
