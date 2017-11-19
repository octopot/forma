package server

import (
	"fmt"
	"net/http"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
)

// UUIDKey used as a context key to store a form UUID.
type UUIDKey struct{}

// FormAPI defines behavior of Form API.
type FormAPI interface {
	// GetV1 responses for `GET /api/v1/{UUID}` request handling.
	GetV1(http.ResponseWriter, *http.Request)
	// PostV1 responses for `POST /api/v1/{UUID}` request handling.
	PostV1(http.ResponseWriter, *http.Request)
}

// New returns new instance of Form API server.
func New() *server {
	return &server{}
}

type server struct{}

func (*server) GetV1(rw http.ResponseWriter, req *http.Request) {
	uuid := req.Context().Value(UUIDKey{}).(data.UUID)
	request := v1.GetRequest{UUID: uuid, Format: req.Header.Get("Accept")}
	fmt.Println(request)
	rw.Write([]byte("get form schema"))
}

func (*server) PostV1(rw http.ResponseWriter, req *http.Request) {
	uuid := req.Context().Value(UUIDKey{}).(data.UUID)
	var httpErr *Error
	if err := req.ParseForm(); err != nil {
		httpErr.InvalidFormData(err).MarshalTo(rw)
		return
	}
	request := v1.PostRequest{UUID: uuid, Data: req.PostForm}
	fmt.Println(request)
	rw.Write([]byte("send form data"))
}

/* TODO for GetV1
supported := map[string]string{
	"json": "application/json",
	//"toml": "application/toml",
	//"xml":  "application/xml",
	//"yaml": "application/yaml",
}
format := strings.ToLower(req.URL.Query().Get("format"))
if format == "" {
	format = "json"
}
ct, ok := supported[format]
if !ok {
	http.Error(rw, "Unsupported output format.", http.StatusUnsupportedMediaType)
	return
}
rw.Header().Set("Content-Type", ct)
fmt.Fprintf(rw, "get form schema in format %q with Content-Type %q", format, ct)
*/
