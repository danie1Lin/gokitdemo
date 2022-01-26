package handlers

// gomock...
type Repository interface {
	CreateEvent(event *Event) error
}
