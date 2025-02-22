package epredis

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

/*
c.Send("SET", "foo", "bar")
c.Send("GET", "foo")
c.Flush()
c.Receive() // reply from SET
v, err = c.Receive() // reply from GET
*/

type CacheInterface interface {
	Set(ctx context.Context, keyValue *RedisKeyValue) error
	SetMultipleKeys(ctx context.Context, keyValues []*RedisKeyValue) error
	DeleteKey(ctx context.Context, key ...string) (interface{}, error)
	Get(ctx context.Context, key string) (interface{}, error)
	GetMultiple(ctx context.Context, keys ...interface{}) ([]string, error)
	GetAllKeys(ctx context.Context, patteren ...string) ([]interface{}, error)
	GetAllValuesFromRedis(ctx context.Context, keys ...interface{}) ([]interface{}, error)
	RemoveKeys(ctx context.Context, keys []string) (interface{}, error)
	Hexists(ctx context.Context, key string, field interface{}) (interface{}, error)
	Hset(ctx context.Context, key string, field interface{}, value interface{}) (interface{}, error)
	HsetMulti(ctx context.Context, key string, fields []interface{}) (interface{}, error)
	Hget(ctx context.Context, key string, field interface{}) (interface{}, error)
	HgetAll(ctx context.Context, key string) (map[string]string, error)
	Exists(ctx context.Context, key string) (interface{}, error)
	Sadd(ctx context.Context, key string, values []interface{}) (interface{}, error)
	Zadd(ctx context.Context, key string, values []interface{}) (interface{}, error)
	CopyKey(ctx context.Context, sourceKey, destinationKey string, replace bool) (interface{}, error)
}

type RedisClient struct {
	pool *redis.Pool
}

var redisClient *RedisClient

func GetRedisClient() CacheInterface {
	if redisClient == nil {
		redisClient = &RedisClient{pool: createRedisPool()}
	}
	return redisClient
}

func (client *RedisClient) Set(ctx context.Context, keyValue *RedisKeyValue) error {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	if keyValue.Key == "" || keyValue.Value == nil {
		return fmt.Errorf("key or Value not provided")
	}

	if keyValue.ExpiryInMillis <= 0 {
		_, err := redisConn.Do("SET", keyValue.Key, keyValue.Value)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := redisConn.Do("SET", keyValue.Key, keyValue.Value, "PX", keyValue.ExpiryInMillis)
	if err != nil {
		return err
	}

	return nil
}

func (client *RedisClient) SetMultipleKeys(ctx context.Context, keyValues []*RedisKeyValue) error {
	// defer gocommon.Timer(ctx, "SetMultipleKeys::")()

	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	if len(keyValues) < 1 {
		return nil
	}
	var keyValuesArgs []interface{}
	for _, kv := range keyValues {
		keyValuesArgs = append(keyValuesArgs, kv.Key, kv.Value)
	}

	_, err := redisConn.Do("MSET", keyValuesArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) DeleteKey(ctx context.Context, key ...string) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redis.String(redisConn.Do("DEL", key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (client *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (client *RedisClient) GetMultiple(ctx context.Context, keys ...interface{}) ([]string, error) {
	// defer gocommon.Timer(ctx, "GetMultiple::")()

	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	values, err := redis.Strings(redisConn.Do("MGET", keys...))
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (client *RedisClient) GetAllKeys(ctx context.Context, patternList ...string) ([]interface{}, error) {

	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	var pattern string
	if len(patternList) > 0 {
		pattern = patternList[0]
	} else {
		pattern = "*"
	}

	keys, err := redis.Values(redisConn.Do("KEYS", pattern))
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (client *RedisClient) GetAllValuesFromRedis(ctx context.Context, keys ...interface{}) ([]interface{}, error) {

	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	var err error
	if len(keys) == 0 {
		keys, err = client.GetAllKeys(ctx)
		if err != nil {
			return nil, err
		}
	}
	values, err := redis.Values(redisConn.Do("MGET", keys...))
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (client *RedisClient) RemoveKeys(ctx context.Context, keys []string) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)
	keysList := stringSliceToInterfaceSlice(keys)

	count, err := redisConn.Do("DEL", keysList...)
	if err != nil {
		return nil, err
	}
	return count, nil
}

func (client *RedisClient) Hexists(ctx context.Context, key string, field interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redisConn.Do("HEXISTS", key, field)
	if err != nil {
		return false, err
	}
	return value, nil
}

func (client *RedisClient) Hset(ctx context.Context, key string, field interface{}, value interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redisConn.Do("HSET", key, field, value)
	if err != nil {
		return false, err
	}
	return value, nil
}

func (client *RedisClient) HsetMulti(ctx context.Context, key string, fields []interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	// Append fields to args slice
	args := append([]interface{}{key}, fields...)
	// Call HSET command with args slice
	value, err := redisConn.Do("HSET", args...)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (client *RedisClient) Hget(ctx context.Context, key string, field interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redisConn.Do("HGET", key, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (client *RedisClient) HgetAll(ctx context.Context, key string) (map[string]string, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	values, err := redis.Values(redisConn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}

	hashData := map[string]string{}
	for i := 0; i < len(values); i += 2 {
		key := values[i]
		value := values[i+1]
		hashData[string(key.([]byte))] = string(value.([]byte))
	}
	return hashData, nil
}

func (client *RedisClient) Exists(ctx context.Context, key string) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	value, err := redisConn.Do("EXISTS", key)
	if err != nil {
		return false, err
	}
	return value, nil
}

// Sadd adds values to a set in Redis.
// Usage:
//
//	key := "mySet"
//	values := []interface{}{"two", "three"}
//
// Parameters:
//   - key: The key of the set in Redis.
//   - values: A list containing the members.
func (client *RedisClient) Sadd(ctx context.Context, key string, values []interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	// Append fields to args slice
	args := append([]interface{}{key}, values...)
	// Call HSET command with args slice
	value, err := redisConn.Do("SADD", args...)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// ZAdd adds values to a sorted set in Redis.
// Usage:
//
//	key := "mySortedSet"
//	values := []interface{}{2, "two", 3, "three"}
//
// Parameters:
//   - key: The key of the sorted set in Redis.
//   - values: A list containing the scores and the associated member.
func (client *RedisClient) Zadd(ctx context.Context, key string, values []interface{}) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	// Append fields to args slice
	args := append([]interface{}{key}, values...)
	// Call HSET command with args slice
	value, err := redisConn.Do("ZADD", args...)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (client *RedisClient) CopyKey(ctx context.Context, sourceKey, destinationKey string, replace bool) (interface{}, error) {
	redisConn := client.pool.Get()
	defer closeConnection(ctx, redisConn)

	args := []interface{}{sourceKey, destinationKey}
	// sending Replace will update the destinationKey if it's already present
	if replace {
		args = append(args, "REPLACE")
	}

	value, err := redisConn.Do("COPY", args...)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func createRedisPool() *redis.Pool {
	defer fmt.Println("createRedisPool Done")
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisPassword := viper.GetString("redis.password")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	fmt.Printf("redis address: %v |%v | %v", redisAddr, redisHost, redisPort)

	maxConnections := viper.GetInt("redis.maxConnections")
	pool := &redis.Pool{
		MaxIdle: maxConnections,
		Dial: func() (redis.Conn, error) {
			if redisPassword == "" {
				return redis.Dial("tcp", redisAddr)
			} else {
				return redis.Dial("tcp", redisAddr, redis.DialPassword(redisPassword))
			}
		},
	}
	resp, err := redis.String(pool.Get().Do("PING"))
	if err != nil || resp != "PONG" {
		return nil
	}
	return pool
}
