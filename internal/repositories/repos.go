package repositories

type Repos interface {
	ProductRepository
}

type repos struct {
	ProductRepository
}

func NewRepos() Repos {
	return &repos{
		NewProductRepository(),
	}
}
