package app

import (
	"context"
	"errors"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/app"
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"
)

const (
	Text = "translated description"
)

var (
	ErrInternal = errors.New("internal error")
)

type testApp struct {
	internal, notFound bool
}

func New(internal, notFound bool) app.App {
	return &testApp{internal, notFound}
}

func (t *testApp) GetShakespearePokemonDescription(ctx context.Context, pokemon string) (string, error) {
	if t.internal {
		return "", ErrInternal
	}

	if t.notFound {
		return "", pokeapi.ErrSpeciesNotFound
	}

	return Text, nil
}
