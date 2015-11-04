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
	// We expect 3 results because we have 3 services: web, image and video
	c := make(chan m.Result, 3)

	// Create a pool of replica. Each service can search the web, images and videos
	webServers := CreateServers("web", nbReplicas, latency)
	imageServers := CreateServers("image", nbReplicas, latency)
	videoServers := CreateServers("video", nbReplicas, latency)

	// Find the first result from all the web servers
	go func() { c <- First(query, webServers...) }()
	// Find the first result from all the image servers
	go func() { c <- First(query, imageServers...) }()
	// Find the first result from all the video servers
	go func() { c <- First(query, videoServers...) }()

	// Define the timeout for a search query
	timeout := time.After(time.Duration(timeoutValue) * time.Millisecond)

	// Go find a result for each service, or
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		// Exit if we've been waiting for too long to have a search result
		case <-timeout:
			return
		}
	}

	return
}
