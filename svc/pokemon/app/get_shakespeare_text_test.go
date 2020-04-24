package app

import (
	"context"
	"testing"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare/test"

	"github.com/stretchr/testify/assert"
)

func TestGetShakespeareTextPreCached(t *testing.T) {
	ctx := context.Background()
	a := NewTestApp(false, false)
	const text = "test"

	// Add our key to the lru cache.
	a.(*app).shakespeareLRU.Add(shakespeareKey(text), test.TranslatedText)

	v, err := a.GetShakespeareText(ctx, text)

	assert.Nil(t, err)
	assert.Equal(t, test.TranslatedText, v)
}

func TestGetShakespeareTextNoCache(t *testing.T) {
	ctx := context.Background()
	a := NewTestApp(false, true)
	const text = "test"

	v, err := a.GetShakespeareText(ctx, text)

	assert.Nil(t, err)
	assert.Equal(t, test.TranslatedText, v)
	assert.True(t, a.(*app).shakespeareLRU.Contains(shakespeareKey(text)))
}

func TestGetShakespeareTextError(t *testing.T) {
	ctx := context.Background()
	a := NewTestApp(false, false)
	const text = "test"

	v, err := a.GetShakespeareText(ctx, text)

	assert.Empty(t, v)
	assert.Equal(t, test.ErrExpected, err)
	assert.False(t, a.(*app).shakespeareLRU.Contains(shakespeareKey(text)))
}
