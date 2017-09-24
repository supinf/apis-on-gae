// Package controllers defines application's routes.
package controllers

import (
	"net/http"
	"strings"

	"github.com/supinf/apis-on-gae/api/google/appengine"
	"github.com/supinf/apis-on-gae/api/google/stackdriver"
	"github.com/supinf/apis-on-gae/api/restapi/operations"
	"github.com/supinf/apis-on-gae/api/restapi/operations/services"
)

// Routes set API handlers
func Routes(api *operations.DemoApisAPI) {
	api.ServicesGetVersionHandler = services.GetVersionHandlerFunc(serviceGetVerion)
	api.ServicesDeleteVersionHandler = services.DeleteVersionHandlerFunc(serviceDeleteVerion)
}

// Wrap wraps original HTTP handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case eqauls(r, "/health"):
			w.WriteHeader(http.StatusOK)
		case eqauls(r, "/_ah/health"):
			w.WriteHeader(http.StatusOK)
		default:
			handler.ServeHTTP(w, r)
			stackdriver.LogInfo(r, struct{ Method, Path, Address string }{
				r.Method, r.URL.Path, ipAddress(r),
			})
		}
	})
}

func eqauls(r *http.Request, url string) bool {
	return url == r.URL.Path
}

func ipAddress(request *http.Request) string {
	if appengine.ProjectID != "" {
		if forwarded := request.Header.Get("X-Forwarded-For"); forwarded != "" {
			parts := strings.Split(forwarded, ",")
			for i, p := range parts {
				parts[i] = strings.TrimSpace(p)
			}
			return parts[0]
		}
	}
	addr := request.RemoteAddr
	idx := strings.LastIndex(addr, ":")
	if idx == -1 {
		return addr
	}
	return addr[:idx]
}
