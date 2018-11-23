package compute

import (
	"fmt"
	"net/http"
	"strconv"
)

// Handler compute
func Handler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		n := 100000

		if q := req.URL.Query().Get("n"); q != "" {
			n, _ = strconv.Atoi(q)
		}

		fmt.Fprintln(res, fibonacci(n))
	})
}
