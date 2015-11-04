package models

import (
	"math"
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
	stdDev := math.Sqrt(float64(s.latency) / 2)
	// Normal law distribution for latency
	latency := rand.NormFloat64()*stdDev + float64(s.latency)

	// Count how much time it took to answer
	start := time.Now()
	time.Sleep(time.Duration(int32(latency)) * time.Millisecond)
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
