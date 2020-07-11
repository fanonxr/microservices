package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks = make(map[string]*Mock)
)

type Response struct {
	response *http.Response
	error error
}

type Mock struct {
	Url string
	HttpMethod string
	Response *http.Response
	Err error

}

func StartMockUps() {
	enabledMocks = true
}

func FlushMockups(){
	mocks = make(map[string]*Mock)
}

func StopMockUps() {
	enabledMocks = false
}

func AddMockup(mock Mock) {
	mocks[GetMockId(mock.HttpMethod, mock.Url)] = &mock
}

func GetMockId(httpMethod string, url string) string{
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error){
	// for testing with a mock
	if enabledMocks {
		mock := mocks[GetMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mockup found")
		}
		return mock.Response, mock.Err
		// return local mock without calling any external resource
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)

}