package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetTopSellingProducts() ([]models.ProductAndFrequency, error)
	GetNewReleaseProducts() ([]models.Product, error)
	GetProduct(productId string) (*models.Product, error)
	GetProductsByName(productName string) ([]models.Product, error)
	GetPromotionalProducts() ([]models.Product, error)
	GetProductsByTypeAndBrand(productId string) ([]models.Product, error)
	AddNewProduct(product models.Product) error
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

func (r *product_repos) GetTopSellingProducts() ([]models.ProductAndFrequency, error) {
	var products []models.ProductAndFrequency
	// var model models.Product
	err := internal.Db.Debug().
		Table("customer_order_detail").
		Preload("PromotionDetailInfo").
		Preload("BrandInfo").
		Preload("CategoryInfo").
		Select("product.*, SUM(customer_order_detail.amount) AS frequency").
		Joins("INNER JOIN product ON customer_order_detail.product_id = product.product_id").
		Where("customer_order_id IN (SELECT customer_order_id FROM customer_order WHERE t_delivery BETWEEN ? AND ? AND customer_order.status = 3 AND payment_status = 2)", time.Now().AddDate(0, 0, -30), time.Now()).
		Group("customer_order_detail.product_id").
		Order("frequency DESC").
		Limit(5).
		Find(&products).Error
	if err != nil {
		internal.Log.Error("GetTopSellingProducts", zap.Any("Error", err))
	}
	return products, err
}

func (r *product_repos) GetNewReleaseProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().Select("*").Table(model.TableName()).
		Preload("BrandInfo").
		Preload("CategoryInfo").
		Where(fmt.Sprintf("%s = ?", model.ColumnIsNew()), "1").
		Find(&products)
	return products, err.Error
}

func (r *product_repos) GetProduct(productId string) (*models.Product, error) {
	model := &models.Product{}

	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnProductId()), productId).
		Preload("PromotionDetailInfo").
		Preload("BrandInfo").
		Preload("CategoryInfo").
		First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetProduct", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
func (r *product_repos) GetProductsByName(productName string) ([]models.Product, error) {
	listModels := []models.Product{}
	model := models.Product{}
	result := internal.Db.Where(fmt.Sprintf("%s LIKE ?", model.ColumnProductName()), "%"+productName+"%").
		Preload("BrandInfo").
		Preload("CategoryInfo").
		Find(&listModels).Error
	return listModels, result
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
func (r *product_repos) GetProductsByTypeAndBrand(productId string) ([]models.Product, error) {
	var products []models.Product
	var product models.Product
	subquery := internal.Db.Debug().Table(product.TableName()).
		Select("brand_id, category_id").
		Where(fmt.Sprintf("%s = ?", product.ColumnProductId()), productId)

	err := internal.Db.Debug().Model(&product).
		Where("(brand_id, category_id) IN (?)", subquery).Find(&products).Error
	if err != nil {
		internal.Log.Error("GetProductsByTypeAndBrand", zap.Any("Error", err))
		return nil, err
	}
	return products, nil
}
func (r *product_repos) AddNewProduct(product models.Product) error {
	return internal.Db.Debug().Create(&product).Error
}
