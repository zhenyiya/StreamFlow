package utils

import (
	"github.com/zhenyiya/constants"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"net/http/pprof"
)

func AdaptRouterToDebugMode(router *mux.Router) *mux.Router {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
	return router
}

func AdaptHTTPWithHeader(w http.ResponseWriter,
	header constants.Header) http.ResponseWriter {
	w.Header().Add(header.Key, header.Value)
	return w
}

// AdaptLimiter adapts the HTTP handler f to a limited handler that allows events up to rate lr and permits bursts of at most lb tokens.
func AdaptLimiter(lim *rate.Limiter, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !lim.Allow() {
			AdaptHTTPWithHeader(w, constants.Header422ExceedLimit)
			io.WriteString(w, "Job request exceeds the maximum handling limit, please try later.\n")
			return
		}
		f(w, r)
		return
	}
}
