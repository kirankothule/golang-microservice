package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initilisation

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "we were not expeting user with id 0")
	assert.NotNil(t, err, "Error should not be nil")
	assert.EqualValues(t, err.Message, "User 0 not found")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "not_found")
}

func TestGetUserUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, user, "User should not be nil")
	assert.EqualValues(t, user.ID, 123)
	assert.EqualValues(t, user.FirstName, "Lee")
	assert.EqualValues(t, user.LastName, "temp")
	assert.EqualValues(t, user.Email, "temp@gmail.com")
}
