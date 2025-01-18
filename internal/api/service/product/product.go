package product

import (
	"errors"

	"github.com/bigxxby/digital-travel-test/internal/api/repo/product"
	"github.com/bigxxby/digital-travel-test/internal/api/repo/user"
	"github.com/bigxxby/digital-travel-test/internal/models"
	"github.com/google/uuid"
)

type ProductService struct {
	ProductRepository product.IProductRepo
	UserRepo          user.IUserRepo
}

type IProductService interface {
	GetProductById(productId *uuid.UUID) (*models.Product, int, error)
	CreateProduct(userId *uuid.UUID, product models.Product) (*models.Product, int, error)
	UpdateProduct(userId *uuid.UUID, product models.Product) (*models.Product, int, error)
	DeleteProduct(userId *uuid.UUID, productId *uuid.UUID) (int, error)
	GetAllProducts() ([]models.Product, int, error)
}

func NewProductService(productRepository product.IProductRepo, userRepo user.IUserRepo) IProductService {
	return &ProductService{
		ProductRepository: productRepository,
		UserRepo:          userRepo,
	}
}

func (ps ProductService) CreateProduct(userId *uuid.UUID, product models.Product) (*models.Product, int, error) {
	//check permission
	user, err := ps.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, 403, err
	}
	if user.Role != "admin" {
		return nil, 403, errors.New("permission denied")
	}
	if product.Name == "" {
		return nil, 400, errors.New("name is required")
	}
	if product.Price < 0 {
		return nil, 400, errors.New("price is required")
	}
	if product.Quantity < 0 {
		return nil, 400, errors.New("quantity is required")
	}

	createdProduct, err := ps.ProductRepository.Create(product)
	if err != nil {
		return nil, 500, err
	}
	return createdProduct, 200, nil
}
func (ps ProductService) UpdateProduct(userId *uuid.UUID, product models.Product) (*models.Product, int, error) {
	if product.Name == "" {
		return nil, 400, errors.New("name is required")
	}
	if product.Price < 0 {
		return nil, 400, errors.New("price is required")
	}
	if product.Quantity < 0 {
		return nil, 400, errors.New("quantity is required")
	}

	updatedProduct, err := ps.ProductRepository.Update(product)
	if err != nil {
		return nil, 500, err
	}
	return updatedProduct, 200, nil
}
func (ps ProductService) DeleteProduct(userId *uuid.UUID, productId *uuid.UUID) (int, error) {

	err := ps.ProductRepository.Delete(productId)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
func (ps ProductService) GetProductById(productId *uuid.UUID) (*models.Product, int, error) {
	product, err := ps.ProductRepository.GetProductById(productId)
	if err != nil {
		return nil, 404, err
	}
	return product, 200, nil
}
func (ps ProductService) GetAllProducts() ([]models.Product, int, error) {
	products, err := ps.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, 500, err
	}
	return products, 200, nil
}
