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