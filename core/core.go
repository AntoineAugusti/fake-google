package core

import (
	"strconv"
	"time"

	m "github.com/AntoineAugusti/fake-google/models"
)

// Find the first result coming from a pool of search servers
func First(query string, replicas ...m.Search) m.Result {
	c := make(chan m.Result, len(replicas))
	searchReplica := func(i int) {
		c <- replicas[i].Search(query)
	}

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

// Create nb numbers of servers of a given type
func CreateServers(serverType string, nb, latency int) []m.Search {
	var servers []m.Search

	for i := 1; i <= nb; i++ {
		id := strconv.Itoa(i)
		servers = append(servers, m.NewSearchServer(serverType+id, latency))
	}

	return servers
}

func Google(query string, nbReplicas, timeoutValue, latency int) (results m.Results) {

	services := []string{"web", "image", "video"}

	// We expect a response from each service we have
	resultChannel := make(chan m.Result, len(services))

	// Run the search query on multiple instances of each service
	for _, serviceName := range services {
		servers := CreateServers(serviceName, nbReplicas, latency)
		go func() {
			resultChannel <- First(query, servers...)
		}()
	}

	// Define the timeout for a search query
	timeout := time.After(time.Duration(timeoutValue) * time.Millisecond)

	// Go find a result for each service
	for i := 0; i < 3; i++ {
		select {
		case result := <-resultChannel:
			results = append(results, result)
		// Exit if we've been waiting for too long to have a search result
		case <-timeout:
			return
		}
	}

	return
}
