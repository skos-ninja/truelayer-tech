package app

import (
	"context"
	"testing"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare/test"

	"github.com/stretchr/testify/assert"
)

func TestGetShakespeareTextPreCached(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, false)
	const text = "test"

	// Add our key to the lru cache.
	a.shakespeareLRU.Add(shakespeareKey(text), test.TranslatedText)

	v, err := a.GetShakespeareText(ctx, text)

	assert.Nil(t, err)
	assert.Equal(t, test.TranslatedText, v)
}

func TestGetShakespeareTextNoCache(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, true)
	const text = "test"

	v, err := a.GetShakespeareText(ctx, text)

	assert.Nil(t, err)
	assert.Equal(t, test.TranslatedText, v)
	assert.True(t, a.shakespeareLRU.Contains(shakespeareKey(text)))
}

func TestGetShakespeareTextError(t *testing.T) {
	ctx := context.Background()
	a := newTestApp(false, false)
	const text = "test"

	v, err := a.GetShakespeareText(ctx, text)

	assert.Empty(t, v)
	assert.Equal(t, test.ErrExpected, err)
	assert.False(t, a.shakespeareLRU.Contains(shakespeareKey(text)))
}
