package main

import (
	"net/http"

	"goframe/app/handler"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	routeJson         = `/json`
	routeDb           = `/db`
	routeQueries      = `/queries`
	routeCachedWorlds = `/cached-worlds`
	routeFortunes     = `/fortunes`
	routeUpdates      = `/updates`
	routePlaintext    = `/plaintext`
)

func main() {
	// Init http server and handler.
	s := g.Server()
	s.SetHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "GoFrame")
		switch r.URL.Path {
		case routePlaintext:
			handler.Plaintext(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(http.StatusText(http.StatusNotFound)))
		}
	})
	s.SetPort(8082)
	s.Run()
}
