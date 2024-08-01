package bootstrap

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func setupLogger() {
	// Create new Logrus logger
	logger := logrus.New()

	// Set the output
	logger.SetOutput(colorable.NewColorableStdout())

	// Use a custom log format with timestamp, colored level
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	// Set the log level
	logger.SetLevel(logrus.DebugLevel)

	// Set as the global logger
	Logger = logger
}

// InitLogger initializes
func InitLogger() *logrus.Logger {
	setupLogger()
	return Logger
}
