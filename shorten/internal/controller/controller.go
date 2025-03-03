package controller

import (
	"context"
)

type shortenRepository interface {
	CreateShortUrl(ctx context.Context, longURL string) (string, error)
}
type Controller struct {
	repo shortenRepository
}

const maxRetries = 5

func New(repo shortenRepository) *Controller {
	return &Controller{repo}
}

func (ctrl *Controller) CreateShortUrl(ctx context.Context, longURL string) (string, error) {
	short_url, err := ctrl.repo.CreateShortUrl(ctx, longURL)
	if err != nil {
		return "", err
	}
	return short_url, err
}
