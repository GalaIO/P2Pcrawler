package misc

import (
	"net/http"
	"net/http/pprof"
)

func StartPProf() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/debug/pprof/", pprof.Index)
	serveMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	serveMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serveMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serveMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	server := &http.Server{
		Addr:    ":10080",
		Handler: serveMux,
	}
	go server.ListenAndServe()
}
