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

	obj := NewJsonLogFileHook(logFileName, logrus.DebugLevel, LogProperties{})

	testNewJsonLogHook(obj, t)
}

func Test_NewJsonLogHook(t *testing.T) {
	obj := NewJsonLogHook(logrus.DebugLevel, LogProperties{}, new(bytes.Buffer))
	testNewJsonLogHook(obj, t)
}

func Test_JsonLogFireProperties(t *testing.T) {
	expect := assert.New(t)

	var expectedProperties = &LogProperties{
		DcName:       randomStr(),
		AppName:      randomStr(),
		PodName:      randomStr(),
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
		DcName:  randomStr(),
		AppName: randomStr(),
	}

	jsonMap := fireAndInterceptAsMapWith(expectedProperties)

	expect.Equal(expectedProperties.DcName, jsonMap[expectedFieldNameDc])
	expect.Empty(jsonMap[expectedFieldNameServiceName])
}

func fireAndInterceptAsMapWith(expectedProperties *LogProperties) map[string]string {
	buffer := new(bytes.Buffer)
	obj := NewJsonLogHook(logrus.DebugLevel, *expectedProperties, buffer)

	entry := newLogEntry(logrus.New(), expectedProperties)
	entry.Level = logrus.InfoLevel
	obj.Fire(entry)

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
