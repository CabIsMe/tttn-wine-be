package models

import "time"

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
	CategoryId   string `json:"category_id"`
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
	FullName    string    `json:"full_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number" validate:"max=11,min=10"`
	Email       string    `json:"email" validate:"required"`
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

type Product struct {
	ProductId       string    `json:"product_id"`
	ProductName     string    `json:"product_name"`
	Cost            float32   `json:"cost"`
	ProductImg      string    `json:"product_img"`
	Description     string    `json:"description"`
	InventoryNumber int       `json:"inventory_number"`
	Status          string    `json:"status"`
	BrandId         string    `json:"brand_id"`
	CategoryId      string    `json:"category_id"`
	IsNew           int8      `json:"is_new"`
	BrandInfo       *Brand    `json:"brand_info" gorm:"references:BrandId;foreignKey:BrandId"`
	CategoryInfo    *Category `json:"category_info" gorm:"references:CategoryId;foreignKey:CategoryId"`
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

type CustomerOrder struct {
	CustomerOrderId string    `json:"customer_order_id"`
	TCreate         time.Time `json:"t_create"`
	FullName        string    `json:"full_name"`
	Address         string    `json:"address"`
	PhoneNumber     string    `json:"phone_number"`
	TDelivery       time.Time `json:"t_delivery"`
	Status          int       `json:"status"`
	EmployeeId      string    `json:"employee_id"`
	DelivererId     string    `json:"deliverer_id"`
	CustomerId      string    `json:"customer_id"`
}

func (c *CustomerOrder) TableName() string {
	return "customer_order"
}

type CustomerOrderDetail struct {
	CustomerOrderDetailId string  `json:"customer_order_detail_id"`
	ProductId             string  `json:"product_id"`
	CustomerOrderId       string  `json:"customer_order_id"`
	Amount                int     `json:"amount"`
	Cost                  float32 `json:"cost"`
}

type Bill struct {
	BillId          string    `json:"bill_id"`
	TCreate         time.Time `json:"t_create"`
	TaxId           string    `json:"tax_id"`
	TaxName         string    `json:"tax_name"`
	CustomerOrderId string    `json:"customer_order_id"`
	EmployeeId      string    `json:"employee_id"`
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
	EmployeeId    string    `json:"employee_id"`
	EmployeeInfo  Employee  `json:"employee_info" gorm:"references:EmployeeId;foreignKey:EmployeeId"`
}

func (Promotion) TableName() string {
	return "promotion"
}

type PromotionDetail struct {
	ProductID          string  `json:"product_id" gorm:"primaryKey"`
	PromotionID        string  `json:"promotion_id" gorm:"primaryKey"`
	DiscountPercentage float32 `json:"discount_percentage"`
}

func (PromotionDetail) TableName() string {
	return "promotion_detail"
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
