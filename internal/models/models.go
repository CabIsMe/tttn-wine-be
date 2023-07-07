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
	Username string `json:"username"`
	Password string `gorm:"column:user_password" json:"password"`
	RoleId   int8   `json:"-"`
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
	Email       string    `json:"email"`
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
