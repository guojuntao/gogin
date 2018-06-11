// +build pprof

package main

import (
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		// /debug/pprof
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
}
