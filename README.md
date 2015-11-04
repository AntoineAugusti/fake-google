[![Software License](http://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/AntoineAugusti/fake-google/LICENSE.md)
# Fake Google
This repository is for educational purpose only.

The goal is to find a behaviour that could be used at Google to handle a search query. We have got 3 services (web, images and videos) and we want to perform a search on each service according to the query. The goal is to respond as fast as possible.

## Architecture
We have got multiple instances of each service. We are going to send the search query **in parallel** to available instances of web servers, images servers and videos servers. For each server we will take **the first** returned search result, to meet our goal to respond as fast as possible.

## Metrics
We will assume that each server answers a query in a time that follows a normal distribution (referred to as `latency`). A search has also a timeout which represents the number of milliseconds we are willing to wait to have search results before exiting (it is possible that search results from all the services have not yet arrived). This is referred to as the `timeout` parameter. Finally, we can control how many instances of each service we have available. This is referred to as the `replicas` parameter.

## Getting started
You can grab this package with the following command:
```
go get github.com/antoineaugusti/fake-google
```

And then build it:
```
cd ${GOPATH%/}/src/github.com/antoineaugusti/fake-google
go build
```

## Running it
From the help manual
```
./fake-google -h
Usage of ./fake-google:
  -latency int
        Mean latency for finding the answer for a query (in ms) (default 55)
  -replicas int
        The number of replicas to create (default 10)
  -timeout int
        Maximum timeout allowed for a search query (in ms) (default 50)
```

## Execution samples
```
./fake-google -timeout 20 -replicas 200 -latency 28
[
  {Search result: `res for: test query` from: `image95` in 18.695281ms}
  {Search result: `res for: test query` from: `web129` in 17.11128ms}
  {Search result: `res for: test query` from: `video13` in 19.058285ms}
]

./fake-google -timeout 20 -replicas 100 -latency 28
[
  {Search result: `res for: test query` from: `web90` in 19.499019ms}
]

./fake-google -timeout 20 -replicas 10 -latency 25
[]

./fake-google -timeout 20 -replicas 100 -latency 20
[
  {Search result: `res for: test query` from: `web90` in 12.735776ms}
  {Search result: `res for: test query` from: `image63` in 12.727817ms}
  {Search result: `res for: test query` from: `video26` in 13.02499ms}
]
```
