package http

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	log "backend-test/helper/log"

	domain "backend-test/domain/book"

	response "backend-test/helper/response"

	"github.com/gorilla/sessions"
)

type BookFormat struct {
	Data   interface{} `json:"data,omitempty"`
	Reason string      `json:"reason,omitempty"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var SESSION_ID = "access"

type BookHttp struct {
	Book domain.BookUsecaseService
}

func NewBookHttp(Book domain.BookUsecaseService) *BookHttp {
	return &BookHttp{Book}
}

func (book *BookHttp) Search(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	header, _ := json.Marshal(r.Header)
	req := r.URL.Query().Encode()

	log.CreateLogResponse(&log.FormatLog{
		Event:   "http|request|search|book",
		Method:  r.Method,
		Header:  string(header),
		Request: req,
	})

	usecase, _ := book.Book.Search(r.URL.Query().Get("s"))

	defer log.CreateLogResponse(&log.FormatLog{
		Event:        "http|response|search|book",
		Method:       r.Method,
		Header:       string(header),
		Request:      req,
		URL:          r.URL.Host,
		ResponseTime: time.Since(t),
	})

	response.BuildResponse(w, true, http.StatusCreated, &BookFormat{
		Data:   usecase,
		Reason: "register has been successfully",
	}, "")

}
