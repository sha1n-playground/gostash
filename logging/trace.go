package logging

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const FieldNameAction = "action"
const FieldNameTraceId = "trace_id"
const FieldNameSegment = "segment"
const FieldNameMarker = "marker"
const FieldNameDuration = "duration_sec"
const MarkerStart = "start"
const MarkerEnd = "end"

type Trace interface {
	StartSegment(segmentName string, args ...interface{}) Segment
	NewSegment() SegmentBuilder
	AddField(name string, value interface{}) Trace
	Entry() *logrus.Entry
	Id() string
}

type trace struct {
	entry *logrus.Entry
	name  string
	id    string
}

func NewTrace(action string, entry *logrus.Entry) Trace {
	id := uuid.New().String()

	return NewTraceWithId(id, action, entry)
}

func NewTraceWithId(id string, action string, entry *logrus.Entry) Trace {
	return &trace{
		entry: entry,
		name:  action,
		id:    id,
	}
}

func (t *trace) StartSegment(segmentName string, args ...interface{}) Segment {
	return t.NewSegment().Start(segmentName, args...)
}

func (t *trace) Entry() *logrus.Entry {
	return baseEntryForTrace(t)
}

func (t *trace) NewSegment() SegmentBuilder {
	return &segmentBuilder{
		parent: t,
		logger: t.entry,
	}
}

func (t *trace) AddField(name string, value interface{}) Trace {
	t.entry = t.entry.WithField(name, value)
	return t
}

func (t *trace) Id() string {
	return t.id
}

func baseEntryForTrace(trace *trace) *logrus.Entry {
	return trace.entry.WithFields(
		logrus.Fields{
			FieldNameTraceId: trace.id,
			FieldNameAction:  trace.name,
		})
}
