package models

import (
	"math/rand"
	"time"
)

type Result struct {
	Result   string
	ServerId string
	Duration string
}

type Results []Result

func (r Result) String() string {
	return "{Search result: `" + r.Result + "` from: `" + r.ServerId + "` in " + r.Duration + "}"
}

type Search interface {
	Search(string) Result
}

type SearchServer struct {
	Id      string
	latency int
}

// Perform a search operation
func (s SearchServer) Search(query string) Result {

	start := time.Now()
	time.Sleep(time.Duration(rand.Int31n(int32(s.latency))) * time.Millisecond)
	elapsed := time.Since(start)

	return Result{
		Result:   "res for: " + query,
		ServerId: s.Id,
		Duration: elapsed.String(),
	}
}

// Create a new search server
func NewSearchServer(name string, latency int) SearchServer {
	return SearchServer{Id: name, latency: latency}
}
