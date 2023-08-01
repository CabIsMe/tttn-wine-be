package repositories

type Repos interface {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
	CustomerOrderRepository
	BillRepository
}

type repos struct {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
	CustomerOrderRepository
	BillRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
		NewAuthRepository(),
		NewEmployeeRepository(),
		NewPromotionRepository(),
		NewCustomerRepository(),
		NewCustomerOrderRepository(),
		NewBillRepository(),
	}
}
