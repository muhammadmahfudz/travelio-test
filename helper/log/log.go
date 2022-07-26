package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type FormatLog struct {
	Event        string        `json:"event,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	Method       string        `json:"method,omitempty"`
	Header       string        `json:"header,omitempty"`
	Request      string        `json:"request,omitempty"`
	URL          string        `json:"url,omitempty"`
	Response     interface{}   `json:"response,omitempty"`
	ResponseTime time.Duration `json:"response_time,omitempty"`
	Error        interface{}   `json:"error,omitempty"`
}

func CreateLogResponse(data *FormatLog) error {

	logs, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		fmt.Println(err)
	}
	var log = logrus.New()

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	s := strings.Replace(string(logs), "\\", "", -1)

	log.WithFields(logrus.Fields{
		"log": strings.Replace(s, "\n", "", -1),
	}).Info()

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.Out = os.Stdout

	return nil
}
