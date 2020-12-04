// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package main

import (
	"fmt"
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func add(a, b int) int {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	return a + b
}

func sendEmail(sender, receiver string) {
	fmt.Println("sender:", sender)
	fmt.Println("receiver:", receiver)
}

func main() {

	// create redis connection pool
	redisPool := &redis.Pool{
		MaxIdle:     3,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)

	// register task
	cli.Register("worker.add", add)
	cli.Register("worker.sendEmail", sendEmail)

	// start workers (non-blocking call)
	cli.StartWorker()

	for {
	}
	// // wait for client request
	// time.Sleep(10 * time.Second)

	// // stop workers gracefully (blocking call)
	// cli.StopWorker()
}
