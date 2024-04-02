package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel           string               `json:"LogLevel"`
	MysqlDsn           string               `json:"MysqlDsn"`
	GameDataConfigPath string               `json:"GameDataConfigPath"`
	OuterIp            string               `json:"OuterIp"`
	AppList            map[string]AppList   `json:"AppList"`
	NetConf            map[string]string    `json:"NetConf"`
	RedisConf          map[string]RedisConf `json:"RedisConf"`
}
type AppList struct {
	App map[string]App `json:"app"`
}
type App struct {
	Port string `json:"port"`
}
type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type NetConf struct {
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(confName string) error {
	if _, err := os.Stat("./conf"); os.IsNotExist(err) {
		os.MkdirAll("./conf", 0644)
	}
	filePath := "./conf/" + confName
	f, err := os.Open(filePath)
	if err != nil {
		return FileNotExist
	}
	defer func() {
		_ = f.Close()
	}()
	c := new(Config)
	d := json.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return err
	}
	CONF = c
	return nil
}

var DefaultConfig = &Config{
	LogLevel:           "Info",
	MysqlDsn:           "root:password@tcp(127.0.0.1:3306)/hkrpg-go?charset=utf8mb4&parseTime=True&loc=Local",
	GameDataConfigPath: "resources",
	OuterIp:            "127.0.0.1",
	AppList: map[string]AppList{
		"9001.1.1.1": {
			App: map[string]App{
				"port_player": {
					Port: "20041",
				},
			},
		},
		"9001.2.1.1": {
			App: map[string]App{
				"port_gt": {
					Port: "20071",
				},
			},
		},
		"9001.3.1.1": {
			App: map[string]App{
				"port_service": {
					Port: "20081",
				},
			},
		},
		"9001.4.1.1": {
			App: map[string]App{
				"port_http": {
					Port: "8080",
				},
			},
		},
		"9001.5.1.1": {
			App: map[string]App{
				"port_http": {
					Port: "20011",
				},
			},
		},
	},
	NetConf: map[string]string{
		"Node": "127.0.0.1:20081",
		"Game": "127.0.0.1:20081",
	},
	RedisConf: map[string]RedisConf{
		"player_login": {
			Addr:     "127.0.0.1:6379",
			Password: "password",
			DB:       1,
		},
	},
}
