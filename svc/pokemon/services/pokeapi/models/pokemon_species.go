package models

// PokemonSpecies follows the naming convention provided here: https://pokeapi.co/docs/v2.html/#pokemon-species
//
// For the purposes of this test we have only implemented the fields we require
type PokemonSpecies struct {
	Name              string       `json:"name"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}
