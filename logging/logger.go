package logging

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	FileName   string
	Level      string
	Colors     bool
	Properties LogProperties
}

type LogProperties struct {
	DcName       string
	AppName      string
	PodName      string
	ServiceName  string
	InstanceName string
}

var loggerEntry *logrus.Entry

func SetLogConfig(config Config) {
	configureLogger(config)
}

func NewEntryFor(obj string) (retVal *logrus.Entry) {
	retVal = loggerEntry.WithField("obj", obj)
	return retVal
}

func init() {
	configureLogger(Config{
		Level:      "info",
		Colors:     false,
		Properties: LogProperties{},
	})
}

func configureLogger(config Config) {
	levelStr := config.Level
	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}

	if config.FileName != "" {
		logFileName := config.FileName
		fileHook := NewJsonLogFileHook(logFileName, logLevel, config.Properties)
		logger.Hooks.Add(fileHook)
	}

	loggerEntry = logrus.NewEntry(logger)
	NewEntryFor("logging").Info("Logging module configured successfully with", config)
}
