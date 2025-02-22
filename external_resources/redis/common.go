package epredis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

type RedisKeyValue struct {
	Key            string
	Value          interface{}
	ExpiryInMillis int
}

func closeConnection(ctx context.Context, redisConn redis.Conn) {
	redisConn.Close()
}

func stringSliceToInterfaceSlice(strSlice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(strSlice))
	for i, v := range strSlice {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
