package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization
	//Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "We are not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not Found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Abdillah", user.FirstName)
	assert.EqualValues(t, "Abu Musa", user.LastName)
	assert.EqualValues(t, "email@email.com", user.Email)

}
