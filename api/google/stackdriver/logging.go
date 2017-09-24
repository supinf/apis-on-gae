package stackdriver

import (
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"github.com/supinf/apis-on-gae/api/google/appengine"
	"golang.org/x/net/context"
)

var logName string

func init() {
	// to be projects/<project-name>/logs/<logName> as `gce_instance`
	logName = os.Getenv("STACKDRIVER_LOGGING_NAME")
	if len(logName) == 0 {
		logName = "apis-on-gae"
	}
}

// LogDebug send log with debug level
func LogDebug(r *http.Request, payload interface{}) {
	client := logClient()
	if client == nil {
		log.Printf("%v", payload)
		return
	}
	defer client.Close()

	logger := client.Logger(logName)
	defer logger.Flush()

	logger.Log(logging.Entry{Severity: logging.Debug, Payload: payload})
}

// LogInfo send log with information level
func LogInfo(r *http.Request, payload interface{}) {
	client := logClient()
	if client == nil {
		log.Printf("%v", payload)
		return
	}
	defer client.Close()

	logger := client.Logger(logName)
	defer logger.Flush()

	logger.Log(logging.Entry{Severity: logging.Info, Payload: payload})
}

// LogWarning send log with warning level
func LogWarning(r *http.Request, payload interface{}) {
	client := logClient()
	if client == nil {
		log.Printf("%v", payload)
		return
	}
	defer client.Close()

	logger := client.Logger(logName)
	defer logger.Flush()

	logger.Log(logging.Entry{Severity: logging.Warning, Payload: payload})
}

// LogError send log with error level
func LogError(r *http.Request, payload interface{}) {
	errorReport(r, payload)

	client := logClient()
	if client == nil {
		log.Printf("%v", payload)
		return
	}
	defer client.Close()

	logger := client.Logger(logName)
	defer logger.Flush()

	logger.Log(logging.Entry{Severity: logging.Error, Payload: payload})
}

// LogCritical send log with critical level
func LogCritical(r *http.Request, payload interface{}) {
	errorReport(r, payload)

	client := logClient()
	if client == nil {
		log.Printf("%v", payload)
		return
	}
	defer client.Close()

	logger := client.Logger(logName)
	defer logger.Flush()

	logger.Log(logging.Entry{Severity: logging.Critical, Payload: payload})
}

func logClient() *logging.Client {
	if appengine.ProjectID == "" {
		return nil
	}
	client, err := logging.NewClient(context.Background(), appengine.ProjectID)
	if err != nil {
		return nil
	}
	return client
}

func errorReport(r *http.Request, payload interface{}) {
	if appengine.ProjectID == "" {
		return
	}
	ctx := context.Background()
	client, err := errorreporting.NewClient(
		ctx,
		appengine.ProjectID,
		appengine.ServiceName,
		appengine.Version,
		true)
	if err != nil {
		return
	}
	defer client.Close()
	defer client.Catch(ctx)

	client.Report(ctx, r, payload)
}
