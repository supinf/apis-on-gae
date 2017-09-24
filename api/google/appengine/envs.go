package appengine

import (
	"os"
)

// ProjectID of Google Cloud
var ProjectID string

// ServiceName of Google App Engine
var ServiceName string

// Version of the service
var Version string

// InstanceID of Google Compute Engine
var InstanceID string

func init() {
	ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	ServiceName = os.Getenv("GAE_SEVICE")
	Version = os.Getenv("GAE_VERSION")
	InstanceID = os.Getenv("GAE_INSTANCE")
}
