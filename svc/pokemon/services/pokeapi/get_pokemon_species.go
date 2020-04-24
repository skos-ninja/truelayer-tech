package pokeapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi/models"
)

var ErrSpeciesNotFound = errors.New("pokemon species not found")

// GetPokemonSpecies implements a request wrapper for https://pokeapi.co/docs/v2.html/#pokemon-species
func (s *service) GetPokemonSpecies(ctx context.Context, pokemon string) (*models.PokemonSpecies, error) {
	// Ensure our request can't just hang forever
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprint(baseURL, "/api/v2/pokemon-species/", pokemon), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrSpeciesNotFound
		}

		return nil, fmt.Errorf("PokeAPI returned: %v", resp.Status)
	}

	species := &models.PokemonSpecies{}
	err = json.NewDecoder(resp.Body).Decode(species)
	if err != nil {
		return nil, err
	}

	return species, nil
}
