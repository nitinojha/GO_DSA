package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"time"
)

func main(){
	f,err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	l := pprof.StartCPUProfile(f)
	fmt.Println(l)
	pprof.StopCPUProfile()
	num := 10
	//Some intensive CPU usagew work
	for i:=0; i < num; i++ {
		_ = math.Pow10(i)
		_ = i*i
	}
	//Allow time for all CPU Usage to be Captured
	time.Sleep( 2 * time.Second)
}