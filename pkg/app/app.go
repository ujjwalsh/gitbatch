package app

import (
	"github.com/isacikgoz/gitbatch/pkg/gui"
	log "github.com/sirupsen/logrus"
)

// The App struct is responsible to hold app-wide related entities. Currently
// it has only the gui.Gui pointer for interface entity.
type App struct {
	Gui *gui.Gui
}

// Setup will handle pre-required operations. It is designed to be a wrapper for
// main method right now.
func Setup(directory, repoPattern, logLevel string) (*App, error) {
	// initiate the app and give it initial values
	app := &App{}
	setLogLevel(logLevel)
	var err error
	directories := generateDirectories(directory, repoPattern)

	// create a gui.Gui struct and set it as App's gui
	app.Gui, err = gui.NewGui(directories)
	if err != nil {
		// the error types and handling is not considered yer
		log.Error(err)
		return app, err
	}
	// hopefull everything went smooth as butter
	log.Trace("App configuration completed")
	return app, nil
}

// Close function will handle if any cleanup is required. e.g. closing streams
// or cleaning temproray files so on and so forth
func (app *App) Close() error {
	return nil
}

// set the level of logging it is fatal by default
func setLogLevel(logLevel string) {
	switch logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.FatalLevel)
	}
	log.WithFields(log.Fields{
		"level": logLevel,
	}).Trace("logging level has been set")
}