package main

import (
	"math/rand"
	_ "net/http/pprof"
	"pkg/physarum"
	"time"
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	rand.Seed(time.Now().UTC().UnixNano())

	physarum.Run()
}
