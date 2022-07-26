package http

import (
	domain "backend-test/domain/users"
	log "backend-test/helper/log"
	response "backend-test/helper/response"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
)

type UsersFormat struct {
	Uid         int64  `json:"uid,omitempty"`
	Reason      string `json:"reason,omitempty"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	UserName    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Gender      string `json:"gender,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Password    string `json:"password,omitempty"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var SESSION_ID = "access"

type UsersHttp struct {
	usecase domain.UsersUsecaseService
}

func NewUsersHttp(usecase domain.UsersUsecaseService) *UsersHttp {
	return &UsersHttp{usecase: usecase}
}

func (user *UsersHttp) Register(w http.ResponseWriter, r *http.Request) {
	var req UsersFormat
	t := time.Now()

	header, _ := json.Marshal(r.Header)
	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&req); err != nil {
		response.BuildResponse(w, false, http.StatusBadGateway, nil, err.Error())
		return
	}

	format, _ := json.Marshal(req)

	log.CreateLogResponse(&log.FormatLog{
		Event:   "http|request|register|user",
		Method:  r.Method,
		Header:  string(header),
		Request: string(format),
	})

	register, err := user.usecase.Register(req.FirstName, req.LastName,
		req.UserName,
		req.Email,
		req.Address,
		req.City,
		req.Phone,
		req.Gender,
		req.DateOfBirth,
		req.Password,
	)

	if err != nil {
		response.BuildResponse(w, false, http.StatusBadGateway, nil, err.Error())
		return
	}

	defer log.CreateLogResponse(&log.FormatLog{
		Event:        "http|response|register|user",
		Method:       r.Method,
		Header:       string(header),
		Request:      string(format),
		URL:          r.URL.Host,
		Response:     register,
		ResponseTime: time.Since(t),
	})

	response.BuildResponse(w, true, http.StatusCreated, &UsersFormat{
		Uid:    register.GetUserId(),
		Reason: "register has been successfully",
	}, "")
}

func (user *UsersHttp) Login(w http.ResponseWriter, r *http.Request) {
	var req UsersFormat
	t := time.Now()

	header, _ := json.Marshal(r.Header)
	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&req); err != nil {
		response.BuildResponse(w, false, http.StatusBadGateway, nil, err.Error())
		return
	}

	format, _ := json.Marshal(req)

	log.CreateLogResponse(&log.FormatLog{
		Event:   "http|request|login|user",
		Method:  r.Method,
		Header:  string(header),
		Request: string(format),
	})

	login, err := user.usecase.Login(req.UserName, req.Password)

	if err != nil {
		response.BuildResponse(w, false, http.StatusBadGateway, nil, err.Error())
		return
	}

	session, _ := store.Get(r, SESSION_ID)
	session.Values["firstname"] = login.GetUserFirstName()
	session.Values["lastname"] = login.GetUserLastName()
	session.Values["username"] = login.GetUserName()
	session.Values["uid"] = login.GetUid()

	if err := session.Save(r, w); err != nil {
		response.BuildResponse(w, false, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	fmt.Println(login.GetUid())

	defer log.CreateLogResponse(&log.FormatLog{
		Event:        "http|response|login|user",
		Method:       r.Method,
		Header:       string(header),
		Request:      string(format),
		URL:          r.URL.Host,
		Response:     login,
		ResponseTime: time.Since(t),
	})

	response.BuildResponse(w, true, http.StatusOK, &UsersFormat{
		Reason: "login successfully",
	}, "")

}

func (user *UsersHttp) Profile(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	header, _ := json.Marshal(r.Header)
	session, _ := store.Get(r, SESSION_ID)

	if len(session.Values) == 0 {
		response.BuildResponse(w, false, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	uid := session.Values["uid"].(string)
	i, _ := strconv.ParseInt(uid, 10, 64)

	profile, err := user.usecase.Profile(i)

	if err != nil {
		response.BuildResponse(w, false, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	defer log.CreateLogResponse(&log.FormatLog{
		Event:        "http|response|profile|user",
		Method:       r.Method,
		Header:       string(header),
		URL:          r.URL.Host,
		Response:     profile,
		ResponseTime: time.Since(t),
	})

	response.BuildResponse(w, true, http.StatusOK, &UsersFormat{
		Uid:         profile.GetUid(),
		FirstName:   profile.GetUserFirstName(),
		LastName:    profile.GetUserLastName(),
		UserName:    profile.GetUserName(),
		Email:       profile.GetUserEmail(),
		Address:     profile.GetUserAddress(),
		City:        profile.GetUserCity(),
		Phone:       profile.GetUserPhone(),
		Gender:      profile.GetUserGender(),
		DateOfBirth: profile.GetUserDateOfBirth(),
		Reason:      "login successfully",
	}, "")
}

func (user *UsersHttp) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SESSION_ID)

	if len(session.Values) == 0 {
		response.BuildResponse(w, false, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	session.Options.MaxAge = -1
	session.Save(r, w)

	response.BuildResponse(w, true, http.StatusCreated, &UsersFormat{
		Reason: "logout has been successfully",
	}, "")
}
