package util

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	log *logrus.Logger
)

const LogFile = "logs/application.log"

func init() {
	f, _ := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	log = logrus.New()

	log.Out = f

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func Fatal(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

//Info ...
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

var (

	// FatalError ...
	FatalError = "%v type=fatal.error"

	// ConfigError ...
	ConfigError = "%v type=util.error"

	// HTTPError ...
	HTTPError = "%v type=http.error"

	// HTTPWarn ...
	HTTPWarn = "%v type=http.warn"

	// Informational ...
	Informational = "%v type=info"
)
