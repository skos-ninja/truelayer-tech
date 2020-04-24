package rpc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	hTest "github.com/skos-ninja/truelayer-tech/lib/http/test"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/app/test"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonNoID(t *testing.T) {
	app := test.New(false, false)
	rpc := New(app)

	r := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(r)

	rpc.GetPokemon(ctx)

	resp := r.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetPokemonSuccess(t *testing.T) {
	app := test.New(false, false)
	rpc := New(app)

	r := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(r)

	// Set our route params so the id is test
	ctx.Params = gin.Params{
		gin.Param{
			Key:   "id",
			Value: "test",
		},
	}

	rpc.GetPokemon(ctx)

	resp := r.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))

	expectedResponse := getPokemonResponse{
		Name:        "test",
		Description: test.Text,
	}
	hTest.AssertJSONMatches(t, resp, expectedResponse)
}

func TestGetPokemonNotFound(t *testing.T) {
	app := test.New(false, true)
	rpc := New(app)

	r := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(r)

	// Set our route params so the id is test
	ctx.Params = gin.Params{
		gin.Param{
			Key:   "id",
			Value: "test",
		},
	}

	rpc.GetPokemon(ctx)

	resp := r.Result()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetPokemonInternalError(t *testing.T) {
	app := test.New(true, false)
	rpc := New(app)

	r := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(r)

	// Set our route params so the id is test
	ctx.Params = gin.Params{
		gin.Param{
			Key:   "id",
			Value: "test",
		},
	}

	rpc.GetPokemon(ctx)

	resp := r.Result()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
