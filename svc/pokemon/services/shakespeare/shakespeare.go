package shakespeare

import (
	"context"
	"net/http"
)

const baseURL = "https://api.funtranslations.com"

type Service interface {
	ConvertText(ctx context.Context, text string) (string, error)
}

type service struct {
	client *http.Client
}

func New() Service {
	return &service{
		client: http.DefaultClient,
	}
}
