package repository

import (
	domain "backend-test/domain/users"
	"database/sql"
)

type UsersRepository struct {
	db *sql.DB
}

type Data struct {
	UserId          int64
	UserFirstname   string
	UserLastName    string
	UserName        string
	UserEmail       string
	UserAddress     string
	UserCity        string
	UserPhone       string
	UserGender      string
	UserDateOfBirth string
	UserPassword    string
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (mysql *UsersRepository) Create(data *domain.Users) error {
	exec, err := mysql.db.Query("INSERT INTO users (id, firstname, lastname, username, email, address, city, photo, gender, dateofbirth, password) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.GetUserId(),
		data.GetUserFirstName(),
		data.GetUserLastName(),
		data.GetUserName(),
		data.GetUserEmail(),
		data.GetUserAddress(),
		data.GetUserCity(),
		data.GetUserPhone(),
		data.GetUserGender(),
		data.GetUserDateOfBirth(),
		data.GetUserPassword(),
	)

	if err != nil {
		return err
	}

	defer exec.Close()

	return nil
}

func (mysql *UsersRepository) Read(username string) (*domain.Users, error) {
	var data Data

	err := mysql.db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(
		&data.UserId,
		&data.UserFirstname,
		&data.UserLastName,
		&data.UserName,
		&data.UserEmail,
		&data.UserAddress,
		&data.UserCity,
		&data.UserPhone,
		&data.UserGender,
		&data.UserDateOfBirth,
		&data.UserPassword)

	if err != nil {
		return nil, err
	}

	users := domain.BuilderUsers()
	users.SetUid(data.UserId)
	users.SetUserFirstName(data.UserFirstname)
	users.SetUserLastName(data.UserLastName)
	users.SetUserName(data.UserName)
	users.SetUserEmail(data.UserEmail)
	users.SetUserAddress(data.UserAddress)
	users.SetUserCity(data.UserCity)
	users.SetUserPhone(data.UserPhone)
	users.SetUserGender(data.UserGender)
	users.SetUserDateOfBirth(data.UserDateOfBirth)
	users.SetUserPassword(data.UserPassword)
	build := users.BuildUsers()

	return build, nil
}

func (mysql *UsersRepository) Profile(uid int64) (*domain.Users, error) {
	var data Data

	err := mysql.db.QueryRow("SELECT * FROM users WHERE id = ?", uid).Scan(
		&data.UserId,
		&data.UserFirstname,
		&data.UserLastName,
		&data.UserName,
		&data.UserEmail,
		&data.UserAddress,
		&data.UserCity,
		&data.UserPhone,
		&data.UserGender,
		&data.UserDateOfBirth,
		&data.UserPassword)

	if err != nil {
		return nil, err
	}

	users := domain.BuilderUsers()
	users.SetUid(data.UserId)
	users.SetUserFirstName(data.UserFirstname)
	users.SetUserLastName(data.UserLastName)
	users.SetUserName(data.UserName)
	users.SetUserEmail(data.UserEmail)
	users.SetUserAddress(data.UserAddress)
	users.SetUserCity(data.UserCity)
	users.SetUserPhone(data.UserPhone)
	users.SetUserGender(data.UserGender)
	users.SetUserDateOfBirth(data.UserDateOfBirth)
	users.SetUserPassword(data.UserPassword)
	build := users.BuildUsers()

	return build, nil
}

func (mysql *UsersRepository) Update() error {
	return nil
}

func (mysql *UsersRepository) Delete() error {
	return nil
}
