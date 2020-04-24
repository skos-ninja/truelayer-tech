package shakespeare

import (
	"context"
	"net/http"
	"testing"

	"github.com/skos-ninja/truelayer-tech/lib/http/test"

	"github.com/stretchr/testify/assert"
)

const (
	okResponse = `{
		"success": {
			"total": 1
		},
		"contents": {
			"translated": "translated",
			"text": "text",
			"translation": "shakespeare"
		}
	}`

	notFoundResponse = `{
	
	}`
)

func TestConvertTextValidResponse(t *testing.T) {
	ctx := context.Background()
	client, close := test.CreateTestHTTPClient(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})
	defer close()

	svc := New()
	svc.(*service).client = client

	translation, err := svc.ConvertText(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, "translated", translation)
}

func TestConvertTextInvalidResponse(t *testing.T) {
	ctx := context.Background()
	client, close := test.CreateTestHTTPClient(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	defer close()

	svc := New()
	svc.(*service).client = client

	translation, err := svc.ConvertText(ctx, "test")
	assert.Empty(t, translation)
	assert.NotNil(t, err)
}
