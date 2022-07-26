package domain

type UsersUsecaseService interface {
	Register(
		firstname string,
		lastname string,
		username string,
		email string,
		address string,
		city string,
		phone string,
		gender string,
		dateOfbirth string,
		password string,
	) (*Users, error)
	Login(username string, password string) (*Users, error)
	Profile(uid int64) (*Users, error)
}

type UsersRepositoryService interface {
	Create(data *Users) error
	Read(username string) (*Users, error)
	Profile(uid int64) (*Users, error)
	Update() error
	Delete() error
}
