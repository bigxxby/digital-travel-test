package product

import (
	"github.com/bigxxby/digital-travel-test/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductRepo interface {
	Create(product models.Product) (*models.Product, error)
	Update(product models.Product) (*models.Product, error)
	Delete(productId *uuid.UUID) error
	GetProductById(productId *uuid.UUID) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}

type ProductRepo struct {
	Db *gorm.DB
}

// Create implements IProductRepo.
func (p *ProductRepo) Create(product models.Product) (*models.Product, error) {
	err := p.Db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func NewProductRepo(db *gorm.DB) IProductRepo {
	return &ProductRepo{
		Db: db,
	}
}
func (p *ProductRepo) Update(product models.Product) (*models.Product, error) {
	err := p.Db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (p *ProductRepo) Delete(productId *uuid.UUID) error {
	err := p.Db.Where("id = ?", productId.String()).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepo) GetProductById(productId *uuid.UUID) (*models.Product, error) {
	var product models.Product
	result := p.Db.Where("id = ?", productId.String()).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
func (p *ProductRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := p.Db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
