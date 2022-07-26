package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	usersHttp "backend-test/internal/delivery/http/users"
	usersRepo "backend-test/internal/repository/users"
	usersUsecase "backend-test/internal/usecase/users"

	bookHttp "backend-test/internal/delivery/http/book"
	bookInfra "backend-test/internal/delivery/infrastructure/book"
	bookUsecase "backend-test/internal/usecase/book"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: SSL},
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnections,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

var client = createHTTPClient()

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env can't be load")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")

	dbconfig := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName
	db, err := sql.Open("mysql", dbconfig)

	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxOpenConns(300)
	db.SetMaxIdleConns(300)
	db.SetConnMaxLifetime(5 * time.Minute)

	defer db.Close()

	r := mux.NewRouter()

	usersRepository := usersRepo.NewUsersRepository(db)
	usersUsecase := usersUsecase.NewUsersUsecase(usersRepository)
	usersHttp := usersHttp.NewUsersHttp(usersUsecase)

	bookInfra := bookInfra.NewBookService(client)
	bookUsecase := bookUsecase.NewBookUsecase(bookInfra)
	bookHttp := bookHttp.NewBookHttp(bookUsecase)

	r.HandleFunc("/api/user/register", usersHttp.Register).Methods("POST")
	r.HandleFunc("/api/user/login", usersHttp.Login).Methods("POST")
	r.HandleFunc("/api/user/profile", usersHttp.Profile).Methods("GET")
	r.HandleFunc("/api/user/logout", usersHttp.Logout).Methods("GET")
	r.HandleFunc("/api/book/list", bookHttp.Search).Methods("GET")

	http.ListenAndServe(":8090", r)

}
