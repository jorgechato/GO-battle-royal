package battleroyale

import (
	"io"
	"net/http"
)

func homeCtrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
