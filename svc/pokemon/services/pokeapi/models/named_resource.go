package models

// NamedAPIResource follows the naming convention provided here: https://pokeapi.co/docs/v2.html/#namedapiresource
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
