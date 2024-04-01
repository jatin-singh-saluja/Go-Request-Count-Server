package main

import (
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type redisConnectionInstance struct {
	redisConnect redis.Conn
}

var instance *redisConnectionInstance
var once sync.Once

func GetRedisInstance() *redisConnectionInstance {
	once.Do(func() {

		instance = &redisConnectionInstance{}

		// Establish a connection to the Redis server using Dial.
		// Dial is an application supplied function for creating and
		// configuring a connection.
		conn, err := redis.Dial("tcp", "redis-server:6379")

		if err != nil {
			log.Fatal(err)
		}
		instance.redisConnect = conn

		// creates visits key in redis to
		// keep count of total number of requests to cluster.
		_, err = conn.Do("SET", "visits", 0)

		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
