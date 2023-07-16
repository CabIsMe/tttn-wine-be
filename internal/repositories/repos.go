package repositories

type Repos interface {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
	CustomerOrderRepository
}

type repos struct {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
	CustomerOrderRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
		NewAuthRepository(),
		NewEmployeeRepository(),
		NewPromotionRepository(),
		NewCustomerRepository(),
		NewCustomerOrderRepository(),
	}
}
