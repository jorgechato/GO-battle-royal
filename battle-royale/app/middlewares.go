package app

import (
	"net/http"
)

func middlewares(next http.Handler) http.Handler {
	// next = middlewares.SwaggerMiddleware(next)
	//ADD custom middlewares
	return next
}
