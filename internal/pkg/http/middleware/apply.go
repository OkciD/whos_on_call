package middleware

import "net/http"

func ApplyMiddlewares(basicHandler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	patchedHandler := basicHandler

	for _, middleware := range middlewares {
		patchedHandler = middleware(patchedHandler)
	}

	return patchedHandler
}
