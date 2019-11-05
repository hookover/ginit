package config

import (
	"gin_api/database/redis"
	"github.com/hookover/logging"
)

func SetupRedis() {
	conf := &redis.Conf{Channels: []*redis.Channel{
		{Name: "local", Password: "", Addr: "localhost:6379", DB: 0},
	}}

	if err := redis.Initialization(conf); err != nil {
		logging.Error().Msg(err.Error())
		panic(err.Error())
	}
}
