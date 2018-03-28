package chi

import (
	"net/http"

	"github.com/go-chi/chi"
)

type realPacker func(string, http.ResponseWriter, *http.Request, http.Handler)

func ctxPacker(packer realPacker, param string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			packer(chi.URLParam(req, param), rw, req, next)
		})
	}
}
