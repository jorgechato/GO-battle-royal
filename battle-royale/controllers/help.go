package help

import (
	"encoding/json"
	"net/http"
)

func getHelpCtrl(w http.ResponseWriter, r *http.Request) {
	// Replace it with the logic of your controller
	users := []string{
		"ok",
	}

	json.NewEncoder(w).Encode(users)
}
