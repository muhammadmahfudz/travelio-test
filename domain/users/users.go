package domain

import "time"

type Users struct {
	uid             int64
	userId          int64
	userFirstname   string
	userLastName    string
	userName        string
	userEmail       string
	userAddress     string
	userCity        string
	userPhone       string
	userGender      string
	userDateOfBirth string
	userPassword    string
}

func BuilderUsers() *Users {
	return &Users{}
}

func (u *Users) BuildUsers() *Users {
	return &Users{
		uid:             u.uid,
		userId:          time.Now().UnixNano() / (1 << 22),
		userFirstname:   u.userFirstname,
		userLastName:    u.userLastName,
		userName:        u.userName,
		userEmail:       u.userEmail,
		userAddress:     u.userAddress,
		userCity:        u.userCity,
		userPhone:       u.userPhone,
		userGender:      u.userGender,
		userDateOfBirth: u.userDateOfBirth,
		userPassword:    u.userPassword,
	}
}

func (u *Users) SetUid(uid int64) {
	u.uid = uid
}

func (u *Users) SetUserFirstName(userFirstName string) {
	u.userFirstname = userFirstName
}

func (u *Users) SetUserLastName(userLastName string) {
	u.userLastName = userLastName
}

func (u *Users) SetUserName(userName string) {
	u.userName = userName
}

func (u *Users) SetUserEmail(userEmail string) {
	u.userEmail = userEmail
}

func (u *Users) SetUserAddress(userAddress string) {
	u.userAddress = userAddress
}

func (u *Users) SetUserCity(userCity string) {
	u.userCity = userCity
}

func (u *Users) SetUserPhone(userPhone string) {
	u.userPhone = userPhone
}

func (u *Users) SetUserGender(userGender string) {

	switch userGender {
	case "laki-laki":
		u.userGender = "l"
	case "perempuan":
		u.userGender = "p"
	}
}

func (u *Users) SetUserDateOfBirth(userDateOfBirth string) {
	u.userDateOfBirth = userDateOfBirth
}

func (u *Users) SetUserPassword(userPassword string) {
	u.userPassword = userPassword
}

func (u *Users) GetUid() int64 {
	return u.uid
}

func (u *Users) GetUserId() int64 {
	return u.userId
}

func (u *Users) GetUserFirstName() string {
	return u.userFirstname
}

func (u *Users) GetUserLastName() string {
	return u.userLastName
}

func (u *Users) GetUserName() string {
	return u.userName
}

func (u *Users) GetUserEmail() string {
	return u.userEmail
}

func (u *Users) GetUserAddress() string {
	return u.userAddress
}

func (u *Users) GetUserCity() string {
	return u.userCity
}

func (u *Users) GetUserPhone() string {
	return u.userPhone
}

func (u *Users) GetUserGender() string {
	return u.userGender
}

func (u *Users) GetUserDateOfBirth() string {
	return u.userDateOfBirth
}

func (u *Users) GetUserPassword() string {
	return u.userPassword
}
