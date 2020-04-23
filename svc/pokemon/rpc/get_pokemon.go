package rpc

import (
	"net/http"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/pokeapi"

	"github.com/gin-gonic/gin"
)

type getPokemonResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *rpc) GetPokemon(c *gin.Context) {
	pokemonID := c.Param("id")
	if pokemonID == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	description, err := r.app.GetShakespearPokemonDescription(c, pokemonID)
	if err != nil {
		if err == pokeapi.ErrSpeciesNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getPokemonResponse{
		Name:        pokemonID,
		Description: description,
	})
}
