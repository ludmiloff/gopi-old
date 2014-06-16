package middleware

import (
	"github.com/pilu/fresh/runner/runnerutils"
	"net/http"
)

func RunFresh(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if runnerutils.HasErrors() {
			//http.Error(w, http.StatusText(500), 500)
			runnerutils.RenderError(w)
			return
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
