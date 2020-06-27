package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)
// steps for test cases
// 1) initialization
// 2) Execution
// 3) Validation


// defining a unit test for not finding a user
func TestGetUserNoUserFound(t *testing.T) {
	// when testing, as many returns there are,
	// will be the amount of test cases that need to be created

	user, err := GetUser(0)
	// asserting that the user is nil
	assert.Nil(t, user, "We were not expecting a user with id 0")
	// asserting that the error is not nil
	assert.NotNil(t, err)
	// asserting the correct status code is returned
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	//
	assert.EqualValues(t, "not_found", err.Code)
	//
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err) // expecting if the user exists
	assert.NotNil(t, user) // expexting the user to be not nill

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t,"Frank", user.FirstName)
	assert.EqualValues(t, "Rogers", user.LastName)
	assert.EqualValues(t, "frank@gmail.com", user.Email)
}
