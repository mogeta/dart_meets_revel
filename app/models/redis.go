package models

import (
	//	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

type Myredis struct{}

type Urls struct {
	Url   string
	Count int
}

func (r *Myredis) Connect() {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * 3, //time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", "127.0.0.1:6379")
			return
		},
	}
}

func (r *Myredis) Add() {
	con := redisPool.Get()
	defer con.Close()
}

func (r *Myredis) Zincr(m map[string]int) {
	con := redisPool.Get()
	defer con.Close()

	for k, v := range m {
		con.Send("zincrby", "prfm", v, k)
	}
	con.Flush()
	//c.Receive()           // reply from SET
	//v, err := c.Receive() // reply from GET
	//fmt.Printf("%s", v)
}

func (r *Myredis) Get() []Urls {

	result := []Urls{}
	con := redisPool.Get()
	defer con.Close()

	con.Send("zrevrange", "prfm", 0, -1, "WITHSCORES")
	con.Flush()

	reply, err := redis.MultiBulk(con.Receive())

	for len(reply) > 0 {
		var title string
		rate := -1 // initialize to illegal value to detect nil.
		reply, err = redis.Scan(reply, &title, &rate)
		if err != nil {
			panic(err)
		}

		result = append(result, Urls{title, rate})
		//fmt.Println(result)
	}

	//fmt.Printf("%s", bdata)
	return result
}
