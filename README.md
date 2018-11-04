[![Build Status](https://travis-ci.org/sha1n/gostash.svg?branch=master)](https://travis-ci.org/sha1n/gostash)

# gostash

A utility on top of [logrus](https://github.com/sirupsen/logrus) for structured logging and tracing in logstash json format. 

```go
import "github.com/sha1n/gostash/logging"

logging.SetLogConfig(
    logging.Config{
        FileName: "./test.baseEntry",
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

logging.
    NewTrace("traceAction", baseEntry).
    StartSegment("segment-1").
    EndWithErrorIf(errors.New("fake error"))
```