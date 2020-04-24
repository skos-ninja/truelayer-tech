package app

import (
	"context"
	"testing"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare/test"

	"github.com/stretchr/testify/assert"
)

func TestGetShakespearePokemonDescriptionNotFound(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, false)
	const pokemon = "test"

	v, err := a.GetShakespearePokemonDescription(ctx, pokemon)

	assert.Empty(t, v)
	assert.Equal(t, pokeapi.ErrSpeciesNotFound, err)
}

func TestGetShakespearePokemonDescriptionFailedTranslation(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(true, false)
	const pokemon = "test"

	v, err := a.GetShakespearePokemonDescription(ctx, pokemon)

	assert.Empty(t, v)
	assert.Equal(t, test.ErrExpected, err)
}

func TestGetShakespearePokemonDescriptionSuccess(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(true, true)
	const pokemon = "test"

	v, err := a.GetShakespearePokemonDescription(ctx, pokemon)

	assert.Nil(t, err)
	assert.Equal(t, test.TranslatedText, v)
}
