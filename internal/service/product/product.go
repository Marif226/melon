package product

import (
	"context"
	"github.com/Marif226/melon/internal/model"
	"github.com/Marif226/melon/internal/repository"
)

type productService struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) *productService {
	return &productService {
		productRepo: productRepo,
	}
}

func (s *productService) Create(ctx context.Context, request model.Product) (*model.Product, error) {
	response, err := s.productRepo.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *productService) List(ctx context.Context, request model.ProductListRequest) ([]*model.Product, error) {
	response, err := s.productRepo.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}