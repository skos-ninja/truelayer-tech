package test

import (
	"context"
	"errors"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare"
)

const (
	TranslatedText = "translated"
)

var ErrExpected = errors.New("expected error")

type testService struct {
	success bool
}

func (t *testService) ConvertText(ctx context.Context, text string) (string, error) {
	if !t.success {
		return "", ErrExpected
	}

	return TranslatedText, nil
}

func New(success bool) shakespeare.Service {
	return &testService{success}
}