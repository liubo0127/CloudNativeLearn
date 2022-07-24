package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuProfile = flag.String("cpuProfile", "/tmp/cpuprofile", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	log.Println(result)
}

// go tool pprof /tmp/cpuprofile
// top
