package grpc_test

import (
	"math"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/server/grpc"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name      string
		timestamp *timestamp.Timestamp
		assert    func(assert.TestingT, *timestamp.Timestamp)
	}{
		{"nil pointer", nil, func(t assert.TestingT, tsp *timestamp.Timestamp) { assert.Nil(t, Time(tsp)) }},
		{"normal use", new(timestamp.Timestamp), func(t assert.TestingT, tsp *timestamp.Timestamp) {
			assert.NotNil(t, Time(tsp))
		}},
		{"invalid timestamp", func() *timestamp.Timestamp {
			tsp := timestamp.Timestamp{Seconds: -1, Nanos: -1}
			return &tsp
		}(), func(t assert.TestingT, ts *timestamp.Timestamp) { assert.Panics(t, func() { Time(ts) }) }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, tc.timestamp)
		})
	}
}

func TestTimestamp(t *testing.T) {
	tests := []struct {
		name   string
		time   *time.Time
		assert func(assert.TestingT, *time.Time)
	}{
		{"nil pointer", nil, func(t assert.TestingT, tp *time.Time) { assert.Nil(t, Timestamp(tp)) }},
		{"normal use", new(time.Time), func(t assert.TestingT, tp *time.Time) { assert.NotNil(t, Timestamp(tp)) }},
		{"invalid time", func() *time.Time {
			tp := time.Time{}.AddDate(-math.MaxInt32, -math.MaxInt32, -math.MaxInt32)
			return &tp
		}(), func(t assert.TestingT, tp *time.Time) { assert.Panics(t, func() { Timestamp(tp) }) }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, tc.time)
		})
	}
}
