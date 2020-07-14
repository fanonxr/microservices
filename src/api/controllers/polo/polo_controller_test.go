package polo

import (
	"github.com/stretchr/testify/assert"
	"microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/Marco", nil)
	c := test_utils.GetMockedContext(request, response)

	Polo(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
