package handlers

import "context"

// gomock...
type Repository interface {
	CreateEvent(ctx context.Context, event *Event) error
}
