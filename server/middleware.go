package server

import (
	"net/http"

	"github.com/kamilsk/form-api/dao"
)

// UUIDKey used as a context key to store form UUID.
type UUIDKey struct{}

// UUID validated form UUID stored in request context.
func UUID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		uuid, ok := req.Context().Value(UUIDKey{}).(string)
		if ok && dao.UUID(uuid).IsValid() {
			next.ServeHTTP(rw, req)
			return
		}
		var err *ErrorMessage
		if !ok {
			err.NotProvidedUUID().MarshalTo(rw)
			return
		}
		err.InvalidUUID().MarshalTo(rw)
	})
}
