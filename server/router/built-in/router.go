package built_in

import (
	"net/http"

	"github.com/kamilsk/form-api/server"
)

// NewRouter returns configured `net/http` router.
func NewRouter(api server.FormAPI, withProfiler bool) http.Handler {
	mux := http.NewServeMux()

	if withProfiler {
	}

	return mux
}
