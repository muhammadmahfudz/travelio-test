package domain

type BookUsecaseService interface {
	Search(key string) (interface{}, error)
}

type BookInfraService interface {
	List(key string) (interface{}, error)
}
