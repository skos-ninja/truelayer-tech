package app

import (
	"context"
)

func (a *app) GetShakespearePokemonDescription(ctx context.Context, pokemon string) (string, error) {
	species, err := a.GetPokemonSpecies(ctx, pokemon)
	if err != nil {
		return "", err
	}

	// Filter to find the first description that is in english.
	//
	// It is presumed here that all the descriptions are in version order.
	// As such selecting the first description will select the latest version.
	flavorText := ""
	for _, flavor := range species.FlavorTextEntries {
		if flavor.Language.Name == "en" {
			flavorText = flavor.FlavorText
			break
		}
	}

	description, err := a.GetShakespeareText(ctx, flavorText)
	if err != nil {
		return "", err
	}

	return description, nil
}
