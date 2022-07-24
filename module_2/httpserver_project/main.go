package main

import (
	"flag"
	"fmt"
	"net"
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

func clientIP(r *http.Request) string {
	var ip string
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func responseHandle(w http.ResponseWriter, r *http.Request) {
	var statusCode int

	defer func() {
		record(clientIP(r), statusCode)
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
		record(clientIP(r), statusCode)
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
