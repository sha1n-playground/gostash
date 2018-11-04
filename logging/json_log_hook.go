package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"strings"
)

type JsonLogHook struct {
	levels       []logrus.Level
	fileLogEntry *logrus.Entry
}

func NewJsonLogFileHook(fileName string, levelToSet logrus.Level, properties LogProperties) (retVal *JsonLogHook) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	return NewJsonLogHook(levelToSet, properties, file)
}

func NewJsonLogHook(levelToSet logrus.Level, properties LogProperties, writer io.Writer) (retVal *JsonLogHook) {
	logrusLogger := logrus.New()
	logrusLogger.Out = writer
	logrusLogger.Formatter = NewLogJsonFormatter()

	newFileLogEntry := newLogEntry(logrusLogger, &properties)

	levels := make([]logrus.Level, 0)
	for _, nextLevel := range logrus.AllLevels {
		levels = append(levels, nextLevel)
		if int32(nextLevel) >= int32(levelToSet) {
			break
		}
	}

	retVal = &JsonLogHook{
		levels:       levels,
		fileLogEntry: newFileLogEntry,
	}
	return retVal
}

// Hook implementation
func (hook *JsonLogHook) Fire(entry *logrus.Entry) error {
	fileEntry := hook.fileLogEntry.WithFields(entry.Data)
	methodName := strings.Title(entry.Level.String())

	logMethod := reflect.ValueOf(fileEntry).MethodByName(methodName)

	parameters := make([]reflect.Value, 1)
	parameters[0] = reflect.ValueOf(entry.Message)

	logMethod.Call(parameters)

	return nil
}

// Hook implementation
func (hook *JsonLogHook) Levels() []logrus.Level {
	return hook.levels
}

func newLogEntry(logger *logrus.Logger, logProperties *LogProperties) *logrus.Entry {
	return logrus.
		NewEntry(logger).
		WithField("dc", logProperties.DcName).
		WithField("service", logProperties.ServiceName).
		WithField("instance", logProperties.InstanceName)

}
