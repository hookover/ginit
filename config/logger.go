package config

import (
	"github.com/hookover/logging"
	"github.com/rs/zerolog"
)

func SetUpLogger() {
	conf := &logging.Conf{DefaultLogFile: "./storage/logs/app.log",Channels: []*logging.Channel{
		{
			Name:    "sql",
			LogFile: "./storage/logs/sql.log",
		},
		{
			Name:    "console",
			Days:    3,
			Level:   zerolog.DebugLevel,
			LogFile: "./storage/logs/console.log",
		},
		{
			Name:    "gin",
			Days:    3,
			Level:   zerolog.DebugLevel,
			LogFile: "./storage/logs/gin.log",
		},
		{
			Name:    "rpc",
			Days:    3,
			Level:   zerolog.DebugLevel,
			LogFile: "./storage/logs/rpc.log",
		},
	}}

	if _, err := logging.Initialization(conf); err != nil {
		panic(err)
	}
}
