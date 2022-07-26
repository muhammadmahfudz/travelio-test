package service

import (
	log "backend-test/helper/log"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Books struct {
	Items []struct {
		VolumeInfo struct {
			Title      string   `json:"title"`
			Authors    []string `json:"authors"`
			ImageLinks struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo,omitempty"`
	} `json:"items"`
}

type BookService struct {
	con *http.Client
}

func NewBookService(con *http.Client) *BookService {
	return &BookService{con}
}

func (con *BookService) List(key string) (interface{}, error) {
	t := time.Now()
	urlStr := os.Getenv("SERVER")

	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "infrastructure|http-request|list|book",
			Error: err.Error(),
		})
	}

	q := req.URL.Query()
	q.Add("q", key)

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := con.con.Do(req)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "infrastructure|http-request|list|book",
			Error: err.Error(),
		})

		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "infrastructure|http-request|list|book",
			Error: err.Error(),
		})

		return nil, err
	}

	io.Copy(ioutil.Discard, resp.Body) // <= NOTE
	defer resp.Body.Close()            // <= NOTE

	header, _ := json.Marshal(req.Header)

	defer log.CreateLogResponse(&log.FormatLog{
		Event:        "infrastructure|http-request|list|book",
		StatusCode:   resp.StatusCode,
		Method:       req.Method,
		Header:       string(header),
		Request:      "key: " + key,
		URL:          req.URL.String(),
		ResponseTime: time.Since(t),
	})

	if resp.StatusCode == 200 {
		var lists Books

		err := json.Unmarshal([]byte(body), &lists)

		if err != nil {
			defer log.CreateLogResponse(&log.FormatLog{
				Event: "infrastructure|http-request|list|book",
				Error: err.Error(),
			})

			return nil, err
		}

		return lists, nil

	} else {
		return nil, errors.New("someting wrong, please confirm to admin for checking log")
	}
}
