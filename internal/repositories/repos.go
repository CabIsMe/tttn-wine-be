package repositories

type Repos interface {
	ProductRepository
	AuthRepository
}

type repos struct {
	ProductRepository
	AuthRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
		NewAuthRepository(),
	}
}
