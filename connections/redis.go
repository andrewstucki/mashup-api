package connections

import (
	"github.com/mashup-cms/mashup-api/globals"
	"github.com/garyburd/redigo/redis"
)

var RedisServer = "localhost:6379"

func SetupRedis() {
	globals.RedisPool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", RedisServer)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
