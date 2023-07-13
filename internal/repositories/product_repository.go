package repositories

import (
	"errors"
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetTheTopSellingProducts() ([]models.Product, error)
	GetNewReleaseProducts() ([]models.Product, error)
	GetProduct(productId string) (*models.Product, error)
	GetPromotionalProducts() ([]models.Product, error)
}

type product_repos struct {
}

func NewProductRepository() ProductRepository {
	return &product_repos{}
}

func (r *product_repos) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().
		Select("*").
		Table(model.TableName()).
		Preload("BrandInfo").
		Preload("CategoryInfo").
		Find(&products).Error
	return products, err
}

func (r *product_repos) GetTheTopSellingProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().
		Select("product.*, brand.brand_name").
		Table(model.TableName()).
		Joins("JOIN brand ON product.brand_id = brand.brand_id").
		Find(&products)
	return products, err.Error
}

func (r *product_repos) GetNewReleaseProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().Table(model.TableName()).
		Where(fmt.Sprintf("%s = ?", model.ColumnIsNew()), "1").
		Find(&products)
	return products, err.Error
}

func (r *product_repos) GetProduct(productId string) (*models.Product, error) {
	model := &models.Product{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnProductId()), productId).
		Preload("BrandInfo").
		Preload("CategoryInfo").
		First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetProduct", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
func (r *product_repos) GetPromotionalProducts() ([]models.Product, error) {
	var products []models.Product
	var product models.Product
	var promotion models.Promotion
	var promotionDetail models.PromotionDetail
	where := fmt.Sprintf("%s in (select %s from %s as pd where pd.%s = (select %s from %s as p where now() between p.%s and p.%s))",
		product.ColumnProductId(), product.ColumnProductId(), promotionDetail.TableName(), promotionDetail.ColumnPromotionId(), promotion.ColumnPromotionId(),
		promotion.TableName(), promotion.ColumnDateStart(), promotion.ColumnDateEnd())
	err := internal.Db.Debug().
		Select("*").
		Table(product.TableName()).
		Where(where).
		Preload("BrandInfo").
		Preload("CategoryInfo").
		Preload("PromotionDetailInfo").
		Find(&products).Error
	return products, err
}
