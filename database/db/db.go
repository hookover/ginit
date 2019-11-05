package db

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hookover/logging"
	"github.com/rs/zerolog"
	"xorm.io/core"
)

type Channel struct {
	Name    string
	DSN     string
	ShowSQL bool
}

type Config struct {
	Channels []*Channel
}

type DB struct {
	Clients map[string]*xorm.Engine
}

var (
	db  *DB
	err error
)

func Initialization(conf *Config) error{
	db = &DB{Clients: make(map[string]*xorm.Engine)}
	for _, cf := range conf.Channels {
		if cf.Name == "" {
			return errors.New("配置名称不能为空")
		}

		db.Clients[cf.Name], err = xorm.NewEngine("mysql", cf.DSN)
		if err != nil {
			return err
		}
		db.Clients[cf.Name].SetMapper(core.SnakeMapper{})
		db.Clients[cf.Name].SetLogger(SqlLogger{Chan: "sql"})
		db.Clients[cf.Name].ShowSQL(cf.ShowSQL)
	}
	return nil
}

func Chan(name string) *xorm.Engine {
	if engine, ok := db.Clients[name]; ok {
		return engine
	}

	for _, engine := range db.Clients {
		return engine
	}

	return nil
}

type SqlLogger struct {
	Chan string
}

func (l SqlLogger) Debug(v ...interface{}) {
	logMap := make(map[string]interface{})
	for idx, n := range v {
		logMap[string(idx)] = n
	}

	logging.Chan(l.Chan).Debug().Fields(logMap)
}
func (l SqlLogger) Debugf(format string, v ...interface{}) {
	logging.Chan(l.Chan).Debug().Msg(fmt.Sprintf(format, v...))
}
func (l SqlLogger) Error(v ...interface{}) {
	logMap := make(map[string]interface{})
	for idx, n := range v {
		logMap[string(idx)] = n
	}

	logging.Chan(l.Chan).Error().Fields(logMap)
}
func (l SqlLogger) Errorf(format string, v ...interface{}) {
	logging.Chan(l.Chan).Error().Msg(fmt.Sprintf(format, v...))
}
func (l SqlLogger) Info(v ...interface{}) {
	logMap := make(map[string]interface{})
	for idx, n := range v {
		logMap[string(idx)] = n
	}

	logging.Chan(l.Chan).Info().Fields(logMap)
}
func (l SqlLogger) Infof(format string, v ...interface{}) {
	logging.Chan(l.Chan).Info().Msg(fmt.Sprintf(format, v...))
}
func (l SqlLogger) Warn(v ...interface{}) {
	logMap := make(map[string]interface{})
	for idx, n := range v {
		logMap[string(idx)] = n
	}

	logging.Chan(l.Chan).Warn().Fields(logMap)
}
func (l SqlLogger) Warnf(format string, v ...interface{}) {
	logging.Chan(l.Chan).Warn().Msg(fmt.Sprintf(format, v...))
}

func (l SqlLogger) Level() core.LogLevel {
	return core.LogLevel(logging.Chan(l.Chan).GetLevel())
}
func (l SqlLogger) SetLevel(c core.LogLevel) {
	logging.Chan(l.Chan).Level(zerolog.Level(c))
}

func (l SqlLogger) ShowSQL(show ...bool) {

}
func (l SqlLogger) IsShowSQL() bool {
	return true
}
