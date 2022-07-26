package usercase

import (
	domain "backend-test/domain/users"
	log "backend-test/helper/log"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type UsersUsecase struct {
	repository domain.UsersRepositoryService
}

func NewUsersUsecase(repository domain.UsersRepositoryService) *UsersUsecase {
	return &UsersUsecase{repository: repository}
}

func (user *UsersUsecase) Register(firstname string,
	lastname string,
	username string,
	email string,
	address string,
	city string,
	phone string,
	gender string,
	dateOfbirth string,
	password string) (*domain.Users, error) {

	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))

	users := domain.BuilderUsers()
	users.SetUserFirstName(firstname)
	users.SetUserLastName(lastname)
	users.SetUserName(username)
	users.SetUserEmail(email)
	users.SetUserAddress(address)
	users.SetUserCity(city)
	users.SetUserPhone(phone)
	users.SetUserGender(gender)
	users.SetUserDateOfBirth(dateOfbirth)
	users.SetUserPassword(pass)
	build := users.BuildUsers()

	err := user.repository.Create(build)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "usecase|mysql|register|user",
			Error: err.Error(),
		})

		return nil, err
	}
	return build, nil
}

func (user *UsersUsecase) Login(username string, password string) (*domain.Users, error) {

	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))

	users, err := user.repository.Read(username)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "usecase|mysql|login|user",
			Error: err.Error(),
		})

		return nil, err
	}

	if users.GetUserPassword() != pass {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "usecase|validate|login|user",
			Error: "username or password wrong",
		})

		return nil, errors.New("username or password wrong")
	}

	return users, nil
}

func (user *UsersUsecase) Profile(uid int64) (*domain.Users, error) {

	users, err := user.repository.Profile(uid)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "usecase|mysql|profile|user",
			Error: err.Error(),
		})

		return nil, err
	}

	return users, nil
}
