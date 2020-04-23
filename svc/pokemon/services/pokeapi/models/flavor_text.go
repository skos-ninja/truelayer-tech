package models

// FlavorText follows the naming convention provided in here: https://pokeapi.co/docs/v2.html/#flavortext
type FlavorText struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedAPIResource `json:"language"`
	Version    NamedAPIResource `json:"version"`
}
