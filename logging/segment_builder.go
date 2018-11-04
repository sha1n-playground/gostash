package logging

import (
	"github.com/sirupsen/logrus"
	"time"
)

type SegmentBuilder interface {
	WithField(name string, value interface{}) SegmentBuilder
	WithFields(fields map[string]interface{}) SegmentBuilder
	WithErrorMarkersOnly() SegmentBuilder
	Start(segmentName string, args ...interface{}) Segment
}

type segmentBuilder struct {
	parent           *trace
	errorMarkersOnly bool
	logger           *logrus.Entry
}

func (builder *segmentBuilder) WithField(name string, value interface{}) SegmentBuilder {
	builder.logger = builder.logger.WithField(name, value)

	return builder
}
func (builder *segmentBuilder) WithFields(fields map[string]interface{}) SegmentBuilder {
	builder.logger = builder.logger.WithFields(fields)

	return builder
}
func (builder *segmentBuilder) WithErrorMarkersOnly() SegmentBuilder {
	builder.errorMarkersOnly = true
	return builder
}
func (builder *segmentBuilder) Start(segmentName string, args ...interface{}) Segment {
	start := time.Now()
	baseEntry := builder.logger.
		WithFields(
			logrus.Fields{
				FieldNameTraceId: builder.parent.id,
				FieldNameAction:  builder.parent.name,
				FieldNameSegment: segmentName,
			})

	var s Segment = &segment{
		logger:    baseEntry,
		parent:    builder.parent,
		name:      segmentName,
		startTime: start,
	}

	if builder.errorMarkersOnly {
		s = &errorMarkersOnlySegment{
			delegate: s,
		}
	}

	s.start(args...)

	return s
}
