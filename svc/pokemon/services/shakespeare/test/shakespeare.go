package test

import (
	"context"
	"errors"

	"github.com/skos-ninja/truelayer-tech/svc/pokemon/services/shakespeare"
)

var ErrExpected = errors.New("expected error")

type testService struct {
	success bool
}

func (t *testService) ConvertText(ctx context.Context, text string) (string, error) {
	if !t.success {
		return "", ErrExpected
	}

	return "translated", nil
}

func New(success bool) shakespeare.Service {
	return &testService{success}
}