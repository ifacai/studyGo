package main

import "flag"

func main() {
	//1010000
	var startId, endId int
	flag.IntVar(&startId, "start", 0, "startId")
	flag.IntVar(&endId, "end", 0, "endId")
	flag.Parse()
}
