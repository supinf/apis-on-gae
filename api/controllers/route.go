// Package controllers defines application's routes.
package controllers

import (
	"log"
	"net/http"

	"github.com/supinf/apis-on-gae/api/restapi/operations"
	"github.com/supinf/apis-on-gae/api/restapi/operations/services"
)

// Routes set API handlers
func Routes(api *operations.DemoApisAPI) {
	api.ServicesGetVersionHandler = services.GetVersionHandlerFunc(serviceGetVerion)
}

// Wrap wraps original HTTP handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		log.Printf("method: %s, path: %s, address: %s", r.Method, r.URL.Path, r.RemoteAddr)
	})
}
