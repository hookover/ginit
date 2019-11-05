package redis

import (
	"errors"
	"github.com/go-redis/redis"
)

var (
	redisInstance       *Redis
	defaultMinIdleConns int = 20
)

type Conf struct {
	Channels []*Channel
}

type Channel struct {
	Name         string
	Addr         string
	Password     string
	DB           int
	MinIdleConns int
}

type Redis struct {
	Clients map[string]*redis.Client
}

func Initialization(conf *Conf) error {
	if len(conf.Channels) == 0 {
		return errors.New("配置channels 不能为空")
	}

	redisInstance = &Redis{Clients: make(map[string]*redis.Client)}
	for _, cf := range conf.Channels {
		if cf.Name == "" {
			return errors.New("配置名称不能为空")
		}
		if cf.MinIdleConns == 0 {
			cf.MinIdleConns = defaultMinIdleConns
		}
		redisInstance.Clients[cf.Name] = redis.NewClient(&redis.Options{
			Addr:         cf.Addr,
			Password:     cf.Password,
			DB:           cf.DB,
			MinIdleConns: cf.MinIdleConns,
		})
	}

	return nil
}

func Chan(channel string) *redis.Client {
	if client, ok := redisInstance.Clients[channel]; ok {
		return client
	}
	for _, client := range redisInstance.Clients {
		return client
	}

	return nil
}
