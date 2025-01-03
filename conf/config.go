package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server
	DataBase
	RedisDB
}

type Server struct {
	Host string
	Port int
}

type DataBase struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type RedisDB struct {
	Addr     string
	Password string
	DB       int
	Protocol int
}

func GetConfig() *Config {
	cnf := viper.New()
	cwd, err := os.Getwd()
	cnf.SetConfigFile(cwd + "/config.json")

	if err != nil {

		panic(fmt.Errorf("Error getting current working directory: %s", err))
	}

	if err := cnf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return &Config{
		Server: Server{
			Host: cnf.GetString("server.host"),
			Port: cnf.GetInt("server.port"),
		},
		DataBase: DataBase{
			User: cnf.GetString("database.user"),
			Pass: cnf.GetString("database.pass"),
			Host: cnf.GetString("database.host"),
			Port: cnf.GetString("database.port"),
			Name: cnf.GetString("database.name"),
		},
		RedisDB: RedisDB{
			Addr:     cnf.GetString("redis.addr"),
			Password: cnf.GetString("redis.pass"),
			DB:       cnf.GetInt("redis.db"),
			Protocol: cnf.GetInt("redis.protocol"),
		},
	}
}
