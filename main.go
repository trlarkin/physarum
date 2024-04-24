package main

import (
	"math/rand"
	"myproject/physarum"
	_ "net/http/pprof"
)

func main() {

	//Set a static seed so we get the same results everytime
	rand.Seed(42)

	physarum.Run()
	// physarum.Tristan()

}
