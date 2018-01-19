package errors

// ApplicationError defines the behavior of application errors.
type ApplicationError interface {
	error

	// Cause returns the underlying cause of the error.
	// It is friendly to `github.com/pkg/errors.Cause` method.
	Cause() error
	// Message returns an error message intended to a user.
	Message() string

	// IsClientError returns true and specific error if the error on the client side.
	// Otherwise, it returns false and nil.
	IsClientError() (ClientError, bool)
	// IsServerError returns true and specific error if the error on the server side.
	// Otherwise, it returns false and nil.
	IsServerError() (ServerError, bool)
}

// ClientError defines the behavior of application errors related to a user error.
type ClientError interface {
	// IsInvalidInput returns true if the error related to invalid data provided by a user.
	IsInvalidInput() bool
	// IsResourceNotFound returns true if the error related to an empty search result.
	IsResourceNotFound() bool
}

// ServerError defines the behavior of application errors related to a server error.
type ServerError interface {
	// IsDatabaseFail returns true if the error related to database problems.
	IsDatabaseFail() bool
	// IsSerializationFail returns true if the error related to serialization problems.
	IsSerializationFail() bool
}

type withCode struct {
	code  int
	msg   string
	cause error
}

type empty string

func (err empty) Error() string {
	return "<nil>"
}

func (err withCode) Error() string {
	return err.Message() + ": " + err.Cause().Error()
}

func (err withCode) Cause() error {
	if err.cause == nil {
		return empty("")
	}
	return err.cause
}

func (err withCode) Message() string {
	if err.msg == "" {
		if err.code < ServerErrorCode {
			return ClientErrorMessage
		}
		return ServerErrorMessage
	}
	return err.msg
}

func (err withCode) IsClientError() (ClientError, bool) {
	if err.code < ServerErrorCode {
		return err, true
	}
	return nil, false
}

func (err withCode) IsServerError() (ServerError, bool) {
	if err.code >= ServerErrorCode {
		return err, true
	}
	return nil, false
}

func (err withCode) IsInvalidInput() bool { return err.code == InvalidInputCode }

func (err withCode) IsResourceNotFound() bool { return err.code == ResourceNotFoundCode }

func (err withCode) IsDatabaseFail() bool { return err.code == DatabaseFailCode }

func (err withCode) IsSerializationFail() bool { return err.code == SerializationFailCode }
