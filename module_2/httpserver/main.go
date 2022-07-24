package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/golang/glog"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter root handler")
	user := r.URL.Query().Get("user")
	fmt.Printf("user is: %s\n", user)
	write, err := w.Write([]byte("enter root handler"))
	fmt.Println(write, err)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	//flag.Set("v", "4")
	glog.InfoDepth(2, "Starting http server...")
	//http.HandleFunc("/test", rootHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/test/", rootHandler)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
}
