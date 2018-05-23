package help

import (
	"encoding/json"
	"net/http"
)

func getHelpCtrl(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(http.StatusOK)
}
