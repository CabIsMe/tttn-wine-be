package models

import "time"

type Brand struct {
	BrandId   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
	BrandImg  string `json:"brand_img"`
	BrandDesc string `json:"brand_desc"`
}
type Category struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
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
type Product struct {
	ProductId       string  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	Cost            float32 `json:"cost"`
	ProductImg      string  `json:"product_img"`
	Description     string  `json:"description"`
	InventoryNumber int     `json:"inventory_number"`
	Status          string  `json:"status"`
	BrandId         string  `json:"brand_id"`
	CategoryId      string  `json:"category_id"`
}

func (Product) TableName() string {
	return "product"
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
