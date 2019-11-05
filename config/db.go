package config

import (
	"gin_api/database/db"
	"github.com/gobuffalo/envy"
)

func SetupMysql() {
	conf := &db.Config{Channels: []*db.Channel{
		{Name: "local", DSN: envy.Get("MYSQL_DSN", ""), ShowSQL: true},
	}}

	if err := db.Initialization(conf); err != nil {
		panic(err.Error())
	}
}
