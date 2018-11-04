package logging

import (
	"bytes"
	"encoding/json"
	"github.com/sha1n/go-playground/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

const expectedFieldNameDc = "dc"
const expectedFieldNameServiceName = "service"
const expectedFieldNameInstance = "instance"

func Test_NewJsonLogFileHook(t *testing.T) {
	logFileName := path.Join(os.TempDir(), randomStr()+"-file.log")
	defer os.Remove(logFileName)

	jsonLogHook := NewJsonLogFileHook(logFileName, logrus.TraceLevel, LogProperties{})

	testNewJsonLogHook(jsonLogHook, t)
}

func Test_NewJsonLogHook(t *testing.T) {
	jsonLogHook := NewJsonLogHook(logrus.TraceLevel, LogProperties{}, new(bytes.Buffer))
	testNewJsonLogHook(jsonLogHook, t)
}

func Test_JsonLogFireProperties(t *testing.T) {
	expect := assert.New(t)

	var expectedProperties = &LogProperties{
		DcName:       randomStr(),
		ServiceName:  randomStr(),
		InstanceName: randomStr(),
	}

	jsonMap := fireAndInterceptAsMapWith(expectedProperties)

	expect.Equal(expectedProperties.DcName, jsonMap[expectedFieldNameDc])
	expect.Equal(expectedProperties.ServiceName, jsonMap[expectedFieldNameServiceName])
	expect.Equal(expectedProperties.InstanceName, jsonMap[expectedFieldNameInstance])
}

func Test_JsonLogFirePartialProperties(t *testing.T) {
	expect := assert.New(t)

	var expectedProperties = &LogProperties{
		DcName: randomStr(),
	}

	jsonMap := fireAndInterceptAsMapWith(expectedProperties)

	expect.Equal(expectedProperties.DcName, jsonMap[expectedFieldNameDc])
	expect.Empty(jsonMap[expectedFieldNameServiceName])
}

func fireAndInterceptAsMapWith(expectedProperties *LogProperties) map[string]string {
	buffer := new(bytes.Buffer)
	jsonLogHook := NewJsonLogHook(logrus.DebugLevel, *expectedProperties, buffer)

	entry := newLogEntry(logrus.New(), expectedProperties)
	entry.Level = logrus.InfoLevel
	jsonLogHook.Fire(entry)

	jsonMap := make(map[string]string)
	json.Unmarshal([]byte(buffer.String()), &jsonMap)

	return jsonMap
}

func testNewJsonLogHook(hook *JsonLogHook, t *testing.T) {
	expect := assert.New(t)

	expect.NotNil(hook)
	expect.Equal(hook.levels, logrus.AllLevels)
}

func randomStr() string {
	return utils.RandomStr50()
}
