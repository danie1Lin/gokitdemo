package repository

import (
	"context"
	"trussdemo/liveevent-service/handlers"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateEvent(ctx context.Context, e *handlers.Event) error {
	return nil
}
