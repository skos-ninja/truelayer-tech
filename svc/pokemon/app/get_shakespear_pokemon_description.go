package app

import (
	"context"
)

func (a *app) GetShakespearPokemonDescription(ctx context.Context, pokemon string) (string, error) {
	species, err := a.GetPokemonSpecies(ctx, pokemon)
	if err != nil {
		return "", err
	}

	// TODO: filter here to find the correct description
	description, err := a.GetShakespearText(ctx, species.FlavorTextEntries[0].FlavorText)
	if err != nil {
		return "", err
	}

	return description, nil
}
