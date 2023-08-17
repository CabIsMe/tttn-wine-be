package models

import (
	"encoding/json"
	"time"
)

var Cos = InitCustomerOrderStatusObject()

type Brand struct {
	BrandId   string `json:"brand_id" gorm:"primaryKey"`
	BrandName string `json:"brand_name"`
	BrandImg  string `json:"brand_img"`
	BrandDesc string `json:"brand_desc"`
}

func (Brand) TableName() string {
	return "brand"
}

type Category struct {
	CategoryId   string `json:"category_id" gorm:"primaryKey"`
	CategoryName string `json:"category_name"`
}

func (Category) TableName() string {
	return "category"
}

type Provider struct {
	ProviderId   string `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
}
type Account struct {
	Username     string `json:"username" validate:"required,email"`
	UserPassword string `gorm:"user_password" json:"password" validate:"required"`
	RoleId       int8   `json:"-"`
	RoleInfo     Role   `gorm:"references:RoleId;foreignKey:RoleId" json:"role_info"`
}

type AccountAndFullName struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserId   string `json:"-"`
	RoleId   int8   `json:"-"`
	Name     string `json:"name" validate:"required"`
}
type AccountInfo struct {
	Username     string    `json:"username" validate:"required,email"`
	RoleId       int8      `json:"-"`
	RoleInfo     Role      `gorm:"references:RoleId;foreignKey:RoleId" json:"role_info"`
	CustomerInfo *Customer `json:"customer_info" gorm:"references:Username;foreignKey:Email"`
	EmployeeInfo *Employee `json:"employee_info" gorm:"references:Username;foreignKey:Email"`
}

func (Account) TableName() string {
	return "accounts"
}

func (Account) ColumnUsername() string {
	return "username"
}
func (Account) ColumnPassword() string {
	return "user_password"
}

func (Account) ColumnRoleId() string {
	return "role_id"
}

type Role struct {
	RoleId   int8   `json:"role_id"`
	RoleName string `json:"role_name"`
}
type Customer struct {
	CustomerId  string    `json:"customer_id"`
	FullName    string    `json:"full_name" validate:"required"`
	Gender      int8      `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"max=11,min=10"`
	Email       string    `json:"email" validate:"required"`
}

func (Customer) TableName() string {
	return "customer"
}
func (Customer) ColumnCustomerId() string {
	return "customer_id"
}
func (Customer) ColumnEmail() string {
	return "email"
}
func (d *Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer
	return json.Marshal(&struct {
		*Alias
		DateOfBirth string `json:"date_of_birth"`
	}{
		Alias:       (*Alias)(d),
		DateOfBirth: d.DateOfBirth.Format("2006-01-02"),
	})
}

type Employee struct {
	EmployeeId  string    `json:"employee_id"`
	FullName    string    `json:"full_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number" validate:"max=11,min=10"`
	Email       string    `json:"email"`
}

func (Employee) TableName() string {
	return "employee"
}
func (Employee) ColumnEmployeeId() string {
	return "employee_id"
}
func (Employee) ColumnEmail() string {
	return "email"
}

type Deliverer struct {
	EmployeeId        string    `json:"employee_id"`
	FullName          string    `json:"full_name"`
	Gender            string    `json:"gender"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Address           string    `json:"address"`
	PhoneNumber       string    `json:"phone_number" validate:"max=11,min=10"`
	Email             string    `json:"email"`
	NumberOfDelivered int64     `json:"numberOfDelivered"`
}
type Product struct {
	ProductId           string           `json:"product_id"`
	ProductName         string           `json:"product_name" validate:"required"`
	Cost                float32          `json:"cost" validate:"required"`
	ProductImg          string           `json:"product_img"`
	Description         string           `json:"description" validate:"required"`
	InventoryNumber     int              `json:"inventory_number" validate:"required"`
	Status              string           `json:"status"`
	BrandId             string           `json:"brand_id" validate:"required"`
	CategoryId          string           `json:"category_id" validate:"required"`
	IsNew               int8             `json:"is_new" validate:"required"`
	BrandInfo           *Brand           `json:"brand_info" gorm:"references:BrandId;foreignKey:BrandId"`
	CategoryInfo        *Category        `json:"category_info" gorm:"references:CategoryId;foreignKey:CategoryId"`
	PromotionDetailInfo *PromotionDetail `gorm:"references:ProductId;foreignKey:ProductId" json:"promotion_detail_info"`
}
type ProductAndFrequency struct {
	ProductId           string           `json:"product_id"`
	ProductName         string           `json:"product_name"`
	Cost                float32          `json:"cost"`
	ProductImg          string           `json:"product_img"`
	Description         string           `json:"description"`
	InventoryNumber     int              `json:"inventory_number"`
	Status              string           `json:"status"`
	BrandId             string           `json:"-"`
	CategoryId          string           `json:"-"`
	IsNew               int8             `json:"is_new"`
	BrandInfo           *Brand           `json:"brand_info" gorm:"references:BrandId;foreignKey:BrandId"`
	CategoryInfo        *Category        `json:"category_info" gorm:"references:CategoryId;foreignKey:CategoryId"`
	PromotionDetailInfo *PromotionDetail `gorm:"references:ProductId;foreignKey:ProductId" json:"promotion_detail_info"`
	Frequency           int              `json:"frequency"`
}

func (Product) TableName() string {
	return "product"
}

func (Product) ColumnIsNew() string {
	return "is_new"
}
func (Product) ColumnProductId() string {
	return "product_id"
}
func (Product) ColumnProductName() string {
	return "product_name"
}
func (Product) ColumnInventoryNumber() string {
	return "inventory_number"
}

type CustomerOrder struct {
	CustomerOrderId         string                `json:"customer_order_id"`
	TCreate                 time.Time             `json:"t_create"`
	FullName                string                `json:"full_name"`
	Address                 string                `json:"address" validate:"required"`
	PhoneNumber             string                `json:"phone_number" validate:"max=11,min=10"`
	TDelivery               *time.Time            `json:"t_delivery"`
	Status                  int8                  `json:"status" validate:"required"`
	PaymentStatus           int8                  `json:"payment_status" validate:"required"`
	EmployeeId              *string               `json:"employee_id"`
	DelivererId             *string               `json:"deliverer_id"`
	CustomerId              string                `json:"customer_id"`
	CustomerOrderDetailInfo []CustomerOrderDetail `gorm:"foreignKey:CustomerOrderId;references:CustomerOrderId" json:"customer_order_detail_info"`
}

func (CustomerOrder) TableName() string {
	return "customer_order"
}
func (CustomerOrder) ColumnCustomerOrderId() string {
	return "customer_order_id"
}
func (CustomerOrder) ColumnStatus() string {
	return "status"
}
func (CustomerOrder) ColumnCustomerId() string {
	return "customer_id"
}

func (d *CustomerOrder) MarshalJSON() ([]byte, error) {
	formatNullableTime := func(t *time.Time) string {
		if t != nil {
			return t.Format("2006-01-02")
		}
		return "" // Return empty string for nil time
	}
	type Alias CustomerOrder
	return json.Marshal(&struct {
		*Alias
		TCreate   string `json:"t_create"`
		TDelivery string `json:"t_delivery"`
	}{
		Alias:     (*Alias)(d),
		TCreate:   d.TCreate.Format("2006-01-02"),
		TDelivery: formatNullableTime(d.TDelivery),
	})
}

type UpdatingCustomerOrder struct {
	CustomerOrderId string    `json:"customer_order_id"`
	TDelivery       time.Time `json:"t_delivery"`
	Status          int8      `json:"status"`
	DelivererId     string    `json:"deliverer_id"`
	EmployeeId      string    `json:"-"`
}

type RevenueByDate struct {
	Date        time.Time `json:"date"`
	TotalAmount int       `json:"total_amount"`
	Revenue     float32   `json:"revenue"`
}

func (d *RevenueByDate) MarshalJSON() ([]byte, error) {
	type Alias RevenueByDate
	return json.Marshal(&struct {
		*Alias
		Date string `json:"date"`
	}{
		Alias: (*Alias)(d),
		Date:  d.Date.Format("2006-01-02"),
	})
}

type CustomerOrderStatusObject struct {
	UNAPPROVED        CustomerOrderStatus
	ORDER_CONFIRM     CustomerOrderStatus
	CHECK_OUT_SUCCESS CustomerOrderStatus
	CHECK_OUT_FAIL    CustomerOrderStatus
}

func InitCustomerOrderStatusObject() *CustomerOrderStatusObject {
	return &CustomerOrderStatusObject{
		UNAPPROVED:        CustomerOrderStatus{1, "The order has not been approved."},
		ORDER_CONFIRM:     CustomerOrderStatus{2, "Appoint delivery personnel"},
		CHECK_OUT_SUCCESS: CustomerOrderStatus{3, "Check out success"},
		CHECK_OUT_FAIL:    CustomerOrderStatus{4, "Order has been canceled."},
	}
}

type CustomerOrderStatus struct {
	StatusCode int8
	StatusDesc string
}

type CustomerOrderDetail struct {
	CustomerOrderDetailId string  `json:"customer_order_detail_id"`
	ProductId             string  `json:"product_id" validate:"required"`
	CustomerOrderId       string  `json:"customer_order_id"`
	Amount                int     `json:"amount" validate:"required"`
	Cost                  float32 `json:"cost" validate:"required"`
	ProductInfo           Product `json:"product_info" gorm:"references:ProductId;foreignKey:ProductId"`
}

func (CustomerOrderDetail) TableName() string {
	return "customer_order_detail"
}

type Bill struct {
	BillId          string    `json:"bill_id"`
	TCreate         time.Time `json:"t_create"`
	TaxId           string    `json:"tax_id"`
	TaxName         string    `json:"tax_name"`
	CustomerOrderId string    `json:"customer_order_id"`
	EmployeeId      string    `json:"employee_id"`
}

func (d *Bill) MarshalJSON() ([]byte, error) {
	type Alias Bill
	return json.Marshal(&struct {
		*Alias
		TCreate string `json:"t_create"`
	}{
		Alias:   (*Alias)(d),
		TCreate: d.TCreate.Format("2006/01/02 15:04"),
	})
}

func (Bill) TableName() string {
	return "bill"
}

type ReturnOrder struct {
	ReturnOrderId   string    `json:"return_order_id"`
	TCreate         time.Time `json:"t_create"`
	CustomerOrderId string    `json:"customer_order_id"`
	EmployeeId      string    `json:"employee_id"`
}
type ReturnOrderDetail struct {
	ReturnOrderId         string `json:"return_order_id"`
	CustomerOrderDetailId string `json:"customer_order_detail_id"`
	Amount                int    `json:"amount"`
}

type Promotion struct {
	PromotionId   string    `json:"promotion_id" gorm:"primaryKey"`
	PromotionName string    `json:"promotion_name"`
	DateStart     time.Time `json:"date_start"`
	DateEnd       time.Time `json:"date_end"`
	Description   string    `json:"description"`
	EmployeeId    string    `json:"employee_id" validate:"required"`
	// EmployeeInfo  Employee  `json:"employee_info" gorm:"references:EmployeeId;foreignKey:EmployeeId"`
}

func (Promotion) ColumnDateEnd() string {
	return "date_end"
}
func (Promotion) ColumnDateStart() string {
	return "date_start"
}
func (Promotion) ColumnPromotionId() string {
	return "promotion_id"
}

type PromotionInput struct {
	PromotionName string `json:"promotion_name"`
	DateStart     string `json:"date_start"`
	DateEnd       string `json:"date_end"`
	Description   string `json:"description"`
}
type PromotionAndPercent struct {
	PromotionId        string    `json:"promotion_id"`
	PromotionName      string    `json:"promotion_name"`
	DateStart          time.Time `json:"date_start"`
	DateEnd            time.Time `json:"date_end"`
	Description        string    `json:"description"`
	EmployeeId         string    `json:"employee_id"`
	DiscountPercentage float32   `json:"discount_percentage"`
}

func (Promotion) TableName() string {
	return "promotion"
}

type PromotionDetail struct {
	ProductId          string  `json:"product_id" gorm:"primaryKey" validate:"required"`
	PromotionId        string  `json:"promotion_id" gorm:"primaryKey" validate:"required"`
	DiscountPercentage float32 `json:"discount_percentage" validate:"required,gt=0,lt=1"`
}

func (PromotionDetail) TableName() string {
	return "promotion_detail"
}
func (PromotionDetail) ColumnProductId() string {
	return "product_id"
}
func (PromotionDetail) ColumnPromotionId() string {
	return "promotion_id"
}

type Order struct {
	OrderId    string    `json:"order_id" gorm:"primaryKey"`
	TCreate    time.Time `json:"t_create"`
	ProviderId string    `json:"provider_id"`
	EmployeeId string    `json:"employee_id"`
}

func (Order) TableName() string {
	return "order"
}

type OrderDetail struct {
	OrderId   string  `json:"order_id" gorm:"primaryKey"`
	ProductId string  `json:"product_id" gorm:"primaryKey"`
	Amount    int     `json:"amount"`
	Cost      float32 `json:"cost"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}

type ProvideProduct struct {
	ProviderId string `json:"provider_id" gorm:"primaryKey"`
	ProductId  string `json:"product_id" gorm:"primaryKey"`
}

func (ProvideProduct) TableName() string {
	return "provide_product"
}

type Receipt struct {
	ReceiptId  string    `json:"receipt_id" gorm:"primaryKey"`
	TCreate    time.Time `json:"t_create"`
	EmployeeId string    `json:"employee_id"`
	OrderId    string    `json:"order_id"`
}

func (Receipt) TableName() string {
	return "receipt"
}

type ReceiptDetail struct {
	ReceiptId string  `json:"receipt_id" gorm:"primaryKey"`
	ProductId string  `json:"product_id" gorm:"primaryKey"`
	Amount    int     `json:"amount"`
	Cost      float32 `json:"cost"`
}

func (ReceiptDetail) TableName() string {
	return "receipt_detail"
}

type Cart struct {
	CustomerId  string  `json:"customer_id"`
	ProductId   string  `json:"product_id"`
	Amount      int     `json:"amount"`
	Cost        float32 `json:"cost" gorm:"<-:false"`
	ProductInfo Product `json:"product_info" gorm:"references:ProductId;foreignKey:ProductId" validate:"omitempty"`
}

func (Cart) TableName() string {
	return "cart"
}
