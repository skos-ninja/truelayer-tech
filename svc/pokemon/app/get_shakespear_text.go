package app

import (
	"context"
)

// shakespeareKey is a custom string type to ensure no collisions
type shakespeareKey string

func (a *app) GetShakespearText(ctx context.Context, text string) (string, error) {
	key := shakespeareKey(text)
	translation, ok := a.shakespeareLRU.Get(key)
	if !ok {
		var err error
		translation, err = a.shakespeare.ConvertText(ctx, text)
		if err != nil {
			return "", err
		}

		a.shakespeareLRU.Add(key, translation)
	}

	return translation.(string), nil
}
