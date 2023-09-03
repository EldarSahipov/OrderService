package service

import (
	"OrderService/internal/models"
	"OrderService/internal/pkg/repository"
)

type Orders interface {
	GetByUID(uid string) (models.Order, error)
	GetAll() ([]models.Order, error)
	Create(order *models.Order) (string, error)
}

type Service struct {
	Orders
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Orders: NewOrderService(repo),
	}
}
