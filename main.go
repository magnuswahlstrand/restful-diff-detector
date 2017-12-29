package main

import (
	"github.com/op/go-logging"
	"github.com/sergi/go-diff/diffmatchpatch"

	"os"
	"time"
)

var logger = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} - %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var logDebug bool = true

func stringAreEqual(text1, text2 string) bool {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)

	if logDebug == true {
		logger.Debug("Result of comparision: ", dmp.DiffPrettyText(diffs))
	}

	return dmp.DiffLevenshtein(diffs) == 0
}

func main() {
	// Set up logging
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(backendFormatter)

	logger.Info("Starting server")
	srv := startHTTPServer()

	logger.Info("Sleeping for 5 seconds")
	time.Sleep(5 * time.Second)

	logger.Warning("Shutting down server")
	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}

	logger.Warning("Server shut down")
}
