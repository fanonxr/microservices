package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMEssage string `json:"message"`
	AError   string `json:"error, omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMEssage
}

func (e *apiError) Error() string {
	return e.AError
}

// Public functions to interact with package

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		AStatus:  statusCode,
		AMEssage: message,
		AError:   "",
	}
}

func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json bodys")
	}
	return &result, nil
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMEssage: message,
		AError:   "",
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMEssage: message,
		AError:   "",
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMEssage: message,
		AError:   "",
	}
}
