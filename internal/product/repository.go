package product

import (
	"fmt"
	"shop-near-u/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddProduct(product *models.FarmerProduct) error {
	// First verify that both Shop and CatalogProduct exist
	var shop models.Farmer
	if err := r.DB.First(&shop, product.ShopID).Error; err != nil {
		return fmt.Errorf("shop with ID %d not found: %w", product.ShopID, err)
	}

	var catalogProduct models.CatalogProduct
	if err := r.DB.First(&catalogProduct, product.CatalogID).Error; err != nil {
		return fmt.Errorf("catalog product with ID %d not found: %w", product.CatalogID, err)
	}

	return r.DB.Create(product).Error
}

func (r *Repository) GetProductsByShopID(shopID uint) ([]models.FarmerProduct, error) {
	var products []models.FarmerProduct
	result := r.DB.Preload("CatalogProduct").Where("shop_id = ?", shopID).Find(&products)
	return products, result.Error
}

func (r *Repository) GetProductByID(productID uint) (*models.FarmerProduct, error) {
	var product models.FarmerProduct
	result := r.DB.Preload("CatalogProduct").First(&product, productID)
	return &product, result.Error
}

func (r *Repository) UpdateProduct(product *models.FarmerProduct) error {
	// First verify that both Shop and CatalogProduct exist
	var shop models.Farmer
	if err := r.DB.First(&shop, product.ShopID).Error; err != nil {
		return fmt.Errorf("shop with ID %d not found: %w", product.ShopID, err)
	}

	var catalogProduct models.CatalogProduct
	if err := r.DB.First(&catalogProduct, product.CatalogID).Error; err != nil {
		return fmt.Errorf("catalog product with ID %d not found: %w", product.CatalogID, err)
	}

	return r.DB.Save(product).Error
}

func (r *Repository) DeleteProduct(productID uint) error {
	return r.DB.Delete(&models.FarmerProduct{}, productID).Error
}
