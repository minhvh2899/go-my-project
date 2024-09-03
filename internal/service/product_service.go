package service

import (
	"my-project/internal/models"
	"my-project/internal/repository"
)

type ProductService interface {
    CreateProduct(title, desc string) (*models.Product, error)
    GetProductByID(id uint) (*models.Product, error)
    UpdateProduct(id uint, title, desc string) (*models.Product, error)
    DeleteProduct(id uint) error
    ListProducts() ([]models.Product, error)
}

type productService struct {
    repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
    return &productService{repo: repo}
}

func (s *productService) CreateProduct(title, desc string) (*models.Product, error) {
    product := &models.Product{Title: title, Desc: desc}
    return s.repo.Create(product)
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
    return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(id uint, title, desc string) (*models.Product, error) {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    product.Title = title
    product.Desc = desc
    return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
    return s.repo.Delete(id)
}

func (s *productService) ListProducts() ([]models.Product, error) {
    return s.repo.FindAll()
}