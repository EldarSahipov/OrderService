package repository

import (
	"OrderService/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Orders interface {
	GetByUID(uid string) (models.Order, error)
	GetAll() ([]models.Order, error)
	Create(order *models.Order) (string, error)
}

type Repository struct {
	Orders
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Orders: NewOrderPostgres(db),
	}
}
