package app

import (
	"github.com/hashicorp/golang-lru/simplelru"
	pTest "github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/test"
	sTest "github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare/test"
)

func newTestApp(pokemon, translation bool) *app {
	pokeAPI := pTest.New(pokemon)
	pLRU, _ := simplelru.NewLRU(1, nil)

	shakespeare := sTest.New(translation)
	sLRU, _ := simplelru.NewLRU(1, nil)

	return &app{
		pokeAPI: pokeAPI,
		pokeLRU: pLRU,

		shakespeare:    shakespeare,
		shakespeareLRU: sLRU,
	}
}
