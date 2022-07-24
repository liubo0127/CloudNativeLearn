package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"sync"

	"github.com/golang/glog"
)

func getVersion() map[string]string {
	version := os.Getenv("VERSION")
	return map[string]string{
		"VERSION": version,
	}
}

func record(host string, statusCode int) {
	defer glog.Flush()
	glog.InfoDepth(1, fmt.Sprintf("Client Host: %s, statusCode: %d\n", host, statusCode))
	fmt.Printf("Client Host: %s, statusCode: %d\n", host, statusCode)
}

func responseHandle(w http.ResponseWriter, r *http.Request) {
	var statusCode int

	defer func() {
		record(r.Host, statusCode)
	}()

	defer func() {
		if err := recover(); err != nil {
			statusCode = 500
		}
		w.WriteHeader(statusCode)
	}()

	headers := r.Header
	version := getVersion()
	w.Header().Set("VERSION", version["VERSION"])
	for key, value := range headers {
		w.Header().Set(key, strings.Join(value, ","))
	}
	statusCode = 200
}

func healthz(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	defer func() {
		record(r.Host, statusCode)
	}()

	defer func() {
		if err := recover(); err != nil {
			statusCode = 500
		}
		w.WriteHeader(statusCode)
	}()

	headers := r.Header
	version := getVersion()
	w.Header().Set("VERSION", version["VERSION"])
	for key, value := range headers {
		w.Header().Set(key, strings.Join(value, ","))
	}
	statusCode = 200
}

func runHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()
	glog.InfoDepth(1, "Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/response", responseHandle)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		glog.FatalDepth(1, err)
		glog.Flush()
	}
}

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go runHttpServer(&wg)
	wg.Wait()
}
