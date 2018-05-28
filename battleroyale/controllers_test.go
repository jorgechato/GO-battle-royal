package battleroyale

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeCtrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeCtrl)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Status code should be equal")

	expected := `{"alive": true}`
	assert.Equal(t, rr.Body.String(), expected, "Body should be equal")
}
