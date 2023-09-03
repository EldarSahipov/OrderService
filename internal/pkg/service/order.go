package service

import (
	"OrderService/internal/cache"
	"OrderService/internal/models"
	"OrderService/internal/pkg/repository"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	repo  repository.Orders
	cache *cache.Cache
}

func NewOrderService(repo repository.Orders) *OrderService {
	c := cache.NewCache()

	orders, err := repo.GetAll()
	if err != nil {
		logrus.Fatalf("an error occurred while restoring the cache from the database: %s", err.Error())
	}
	for _, order := range orders {
		c.SetOrderFromCache(order.OrderUid, order)
	}

	return &OrderService{
		repo:  repo,
		cache: c,
	}
}

func (s *OrderService) Create(order *models.Order) (string, error) {
	uid, err := s.repo.Create(order)
	if err != nil {
		return "", err
	} else {
		s.cache.SetOrderFromCache(uid, *order)
		return uid, err
	}
}

func (s *OrderService) GetByUID(uid string) (models.Order, error) {
	order, found := s.cache.GetOrderFromCacheById(uid)
	if found {
		return order, nil
	}

	order, err := s.repo.GetByUID(uid)
	if err != nil {
		return models.Order{}, err
	}

	s.cache.SetOrderFromCache(uid, order)
	return order, nil
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.repo.GetAll()
}
