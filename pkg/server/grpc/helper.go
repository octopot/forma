package grpc

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/pkg/errors"
)

// Time converts a google.protobuf.Timestamp proto to a time.Time.
// It panics if the passed Timestamp is invalid.
func Time(ts *timestamp.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	tp, err := ptypes.Timestamp(ts)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into time.Time", *ts))
	}
	return &tp
}

// Timestamp converts a time.Time to a google.protobuf.Timestamp proto.
// It panics if the resulting Timestamp is invalid.
func Timestamp(tp *time.Time) *timestamp.Timestamp {
	if tp == nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(*tp)
	if err != nil {
		panic(errors.Wrapf(err, "converting %#v into google.protobuf.Timestamp", *tp))
	}
	return ts
}

func ptrToID(id string) *domain.ID {
	if id == "" {
		return nil
	}
	ptr := new(domain.ID)
	*ptr = domain.ID(id)
	return ptr
}
