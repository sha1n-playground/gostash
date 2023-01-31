[![CI](https://github.com/sha1n/gostash/actions/workflows/go.yml/badge.svg)](https://github.com/sha1n/gostash/actions/workflows/go.yml)

# gostash

A utility on top of [logrus](https://github.com/sirupsen/logrus) for structured logging and tracing in logstash json format. 


## Usage Example: (see <repo>/example.go)
```go
package main

import (
	"github.com/sha1n/gostash/logging"
)

func main() {
	logging.SetLogConfig(
		logging.Config{
			FileName: "./test.log",
			Level:    "info",
			Properties: logging.LogProperties{
				DcName:       "east",
				ServiceName:  "my-service",
				InstanceName: "my-instance",
			},
		})

	baseEntry := logging.
		NewEntry("test").
		WithField("user", "me")

	// start a trace segment
	traceSegment := logging.
		NewTrace("traceAction", baseEntry).
		StartSegment("segment-1")

	// your business logic here...


	// end a trace segment with error
	traceSegment.End()
}
```

### Segment start document:
```json
{
	"@timestamp": "2018-11-04T15:35:38.041+02:00",
	"action": "traceAction",
	"dc": "east",
	"instance": "my-instance",
	"level": "info",
	"marker": "start",
	"message": "",
	"name": "test",
	"segment": "segment-1",
	"service": "my-service",
	"trace_id": "81775219-3a91-443c-b77d-b8996337952f",
	"user": "me"
}
```

### Segment end document:
```json
{
	"@timestamp": "2018-11-04T15:35:38.041+02:00",
	"action": "traceAction",
	"dc": "east",
	"duration_sec": 0.000062782,
	"instance": "my-instance",
	"level": "info",
	"marker": "end",
	"message": "",
	"name": "test",
	"segment": "segment-1",
	"service": "my-service",
	"trace_id": "81775219-3a91-443c-b77d-b8996337952f",
	"user": "me"
}
```
