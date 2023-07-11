package repositories

type Repos interface {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
}

type repos struct {
	ProductRepository
	AuthRepository
	EmployeeRepository
	PromotionRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
		NewAuthRepository(),
		NewEmployeeRepository(),
		NewPromotionRepository(),
	}
}
