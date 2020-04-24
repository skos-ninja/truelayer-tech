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
	a := NewTestApp(false, false)
	const pokemon = "test"

	// Add our key to the lru cache.
	a.(*app).pokeLRU.Add(pokeKey(pokemon), test.ExpectedModel)

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Nil(t, err)
	assert.Equal(t, test.ExpectedModel, v)
}

func TestGetPokemonSpeciesNoCache(t *testing.T) {
	ctx := context.Background()
	a := NewTestApp(true, false)
	const pokemon = "test"

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Nil(t, err)
	assert.Equal(t, test.ExpectedModel, v)
	assert.True(t, a.(*app).pokeLRU.Contains(pokeKey(pokemon)))
}

func TestGetPokemonSpeciesError(t *testing.T) {
	ctx := context.Background()
	a := NewTestApp(false, false)
	const pokemon = "test"

	v, err := a.GetPokemonSpecies(ctx, pokemon)

	assert.Empty(t, v)
	assert.Equal(t, pokeapi.ErrSpeciesNotFound, err)
	assert.False(t, a.(*app).pokeLRU.Contains(pokeKey(pokemon)))
}
