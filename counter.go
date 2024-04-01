package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gomodule/redigo/redis"
)

type requestCounter struct {
	// keeps count of total number of requests to the server.
	val int64
}

func (c *requestCounter) incrementVal(d int) {
	atomic.AddInt64(&c.val, int64(d))
}

func (c *requestCounter) getVal() int {
	return int(atomic.LoadInt64(&c.val))
}

// keeps track of total number of
// requests to the server, and the cluster.
func (c *requestCounter) requestCounterHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		return
	}

	fmt.Fprintf(w, "You are talking to instance %s on port %s\n",
		os.Getenv("HOSTNAME"), os.Getenv("PORT"))

	c.incrementVal(1)
	fmt.Fprintf(w, "This is %d th request to this instance\n", c.getVal())

	// retrieves connection to the redis server
	conn := GetRedisInstance().redisConnect

	redisKey := "visits"

	// using Do(), retrieves the value associated with "visits" key in Redis.
	// clusterRequestCount stores total requests to the cluster.
	clusterRequestCount, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		fmt.Printf("%s does not exist\n", redisKey)
	} else if err != nil {
		log.Fatal(err)
	} else {
		clusterRequestCount = clusterRequestCount + 1
		fmt.Fprintf(w, "This is %d th request to this cluster\n", clusterRequestCount)
		_, err := conn.Do("SET", "visits", clusterRequestCount)
		if err != nil {
			log.Fatal(err)
		}
	}

}
