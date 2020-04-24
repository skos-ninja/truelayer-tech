package app

import (
	"context"
	"testing"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/test"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonSpeciesPreCached(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, false)
	const pokemon = "test"

	// Add our key to the lru cache.
	a.pokeLRU.Add(pokeKey(pokemon), test.ExpectedModel)

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Nil(t, err)
	assert.Equal(t, test.ExpectedModel, v)
}

func TestGetPokemonSpeciesNoCache(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(true, false)
	const pokemon = "test"

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Nil(t, err)
	assert.Equal(t, test.ExpectedModel, v)
	assert.True(t, a.pokeLRU.Contains(pokeKey(pokemon)))
}

func TestGetPokemonSpeciesError(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, false)
	const pokemon = "test"

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Empty(t, v)
	assert.Equal(t, pokeapi.ErrSpeciesNotFound, err)
	assert.False(t, a.pokeLRU.Contains(pokeKey(pokemon)))
}
