package server

import (
	"context"
	"net/http"

	"github.com/kamilsk/form-api/data"
)

// ValidateUUID validates form UUID and makes a decision what to return to a client.
func ValidateUUID(formUUID string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	var err *Error
	uuid := data.UUID(formUUID)
	if uuid.IsEmpty() {
		err.NotProvidedUUID().MarshalTo(rw)
		return
	}
	if !uuid.IsValid() {
		err.InvalidUUID().MarshalTo(rw)
		return
	}
	next.ServeHTTP(rw,
		req.WithContext(
			context.WithValue(req.Context(), UUIDKey{}, uuid)))
}
