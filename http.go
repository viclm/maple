package maple

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var rjson = regexp.MustCompile(`\bapplication/json\b`)

type Options struct {
	Url      string
	Method   string
	headers  map[string]string
	timeout  int
	datatype string
	data     map[string]interface{}
}

type HTTPError struct {
	Code    int
	Message string
}

func (err *HTTPError) Error() string {
	return err.Message
}

func GetJSON(options Options, result interface{}) error {
	client := http.Client{}

	request, err := http.NewRequest(options.Method, options.Url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	status := response.StatusCode

	if status >= 200 && status < 300 || status == 304 {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if rjson.MatchString(strings.Join(response.Header["Content-Type"], ";")) {
			if err := json.Unmarshal(body, result); err != nil {
				return err
			} else {
				return nil
			}
		}

		return &HTTPError{status, "response isn't json"}

	}

	return &HTTPError{status, response.Status}
}
