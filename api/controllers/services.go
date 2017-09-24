package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/supinf/apis-on-gae/api/google/appengine"
	"github.com/supinf/apis-on-gae/api/google/stackdriver"
	"github.com/supinf/apis-on-gae/api/models"
	"github.com/supinf/apis-on-gae/api/restapi/operations/services"
)

func serviceGetVerion(params services.GetVersionParams) middleware.Responder {
	version := appengine.Version
	if len(version) == 0 {
		version = "edge"
	}
	return services.NewGetVersionOK().WithPayload(&models.Version{
		Version: swag.String(version),
	})
}

func serviceDeleteVerion(params services.DeleteVersionParams) middleware.Responder {
	code := http.StatusForbidden
	stackdriver.LogError(params.HTTPRequest, struct{ Message string }{
		http.StatusText(code),
	})
	return services.NewDeleteVersionDefault(code).WithPayload(&models.Error{
		Code:    swag.String(fmt.Sprintf("%d", code)),
		Message: swag.String(http.StatusText(code)),
	})
}
