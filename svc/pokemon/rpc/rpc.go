package rpc

import (
	"github.com/skos-ninja/truelayer-tech/svc/pokemon/app"

	"github.com/gin-gonic/gin"
)

type RPC interface {
	GetPokemon(c *gin.Context)
}

type rpc struct {
	app app.App
}

func New(app app.App) RPC {
	return &rpc{
		app: app,
	}
}