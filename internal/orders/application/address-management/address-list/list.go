package addressList

import "goOnGo/internal/orders/model/domain/address"

type Repository interface {
	FindAll() ([]*address.Address, error)
}

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (handler *Handler) Handle() ([]*address.Address, error) {
	return handler.repository.FindAll()
}
