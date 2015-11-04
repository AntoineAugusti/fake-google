package main

import (
	"flag"
	"fmt"

	"github.com/AntoineAugusti/fake-google/core"
)

func main() {
	nbReplicas := flag.Int("replicas", 10, "The number of replicas to create")
	timeout := flag.Int("timeout", 50, "Maximum timeout allowed for a search query (in ms)")
	latency := flag.Int("latency", 150, "Maximum latency for finding the answer for a query (in ms)")

	flag.Parse()

	fmt.Println(core.Google("test query", *nbReplicas, *timeout, *latency))
}
