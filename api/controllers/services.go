package controllers

import (
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/supinf/apis-on-gae/api/models"
	"github.com/supinf/apis-on-gae/api/restapi/operations"
	"github.com/supinf/apis-on-gae/api/restapi/operations/services"
)

func serviceAPI(api *operations.DemoApisAPI) {
	c := cServices{}
	api.ServicesGetVersionHandler = services.GetVersionHandlerFunc(c.getVerion)
}

type cServices struct{}

func (c cServices) getVerion(params services.GetVersionParams) middleware.Responder {
	version := os.Getenv("APP_VERSION")
	if len(version) == 0 {
		version = "edge"
	}
	return services.NewGetVersionOK().WithPayload(&models.Version{
		Version: swag.String(version),
	})
}
