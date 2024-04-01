package main

import (
	"log"
	"net/http"
)

func main() {

	rdc := GetRedisInstance().redisConnect

	// use defer to ensure the redis connection is
	// closed before exiting the main() function.
	defer rdc.Close()

	// keeps track of the total number of
	// requests to the server, and the cluster
	counter := &requestCounter{
		val: 0,
	}

	http.HandleFunc("/", counter.requestCounterHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
