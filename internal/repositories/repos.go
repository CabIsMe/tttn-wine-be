package repositories

type Repos interface {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
}

type repos struct {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
	CustomerRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
		NewAuthRepository(),
		NewEmployeeRepository(),
		NewPromotionRepository(),
		NewCustomerRepository(),
	}
}
