package test

import (
	"context"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

var ExpectedModel = &models.PokemonSpecies{
	Name: "expected",
	FlavorTextEntries: []models.FlavorText{
		{
			FlavorText: "expected is cool",
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
}

type testService struct {
	success bool
}

func (t *testService) GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error) {
	if !t.success {
		return nil, pokeapi.ErrSpeciesNotFound
	}

	return ExpectedModel, nil
}

func New(success bool) pokeapi.Service {
	return &testService{success}
}
