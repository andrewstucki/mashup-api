package globals

import (
	"github.com/coopernurse/gorp"
	"github.com/garyburd/redigo/redis"
)

var PostgresConnection *gorp.DbMap
var RedisPool *redis.Pool
var ExternalQueue bool