package hkrpg_go_pe

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	LogLevel           string      `json:"LogLevel"`
	GameDataConfigPath string      `json:"GameDataConfigPath"`
	SqlPath            string      `json:"SqlPath"`
	Dispatch           *Dispatch   `json:"Dispatch"`
	GameServer         *GameServer `json:"GameServer"`
}

type Dispatch struct {
	AutoCreate   bool           `json:"AutoCreate"`
	Addr         string         `json:"Addr"`
	Port         string         `json:"Port"`
	OuterAddr    string         `json:"OuterAddr"`
	DispatchList []DispatchList `json:"DispatchList"`
}
type DispatchList struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	DispatchUrl string `json:"dispatchUrl"`
}

type GameServer struct {
	MaxPlayer     int32  `json:"MaxPlayer"`
	InnerAddr     string `json:"InnerAddr"`
	Port          string `json:"Port"`
	OuterAddr     string `json:"OuterAddr"`
	IsJumpMission bool   `json:"IsJumpMission"`
}

var CONF *Config = nil

func GetConfig() *Config {
	return CONF
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(confName string) error {
	filePath := "./" + confName
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
	GameDataConfigPath: "resources",
	SqlPath:            "hkrpg-go-pe.db",
	Dispatch: &Dispatch{
		AutoCreate: true,
		Addr:       "0.0.0.0",
		Port:       "8080",
		OuterAddr:  "http://127.0.0.1:8080",
		DispatchList: []DispatchList{
			{
				Name:        "hkrpg-go",
				Title:       "os_usa",
				Type:        "2",
				DispatchUrl: "/query_gateway",
			},
			{
				Name:        "hkrpg-official_os",
				Title:       "os_usa",
				Type:        "2",
				DispatchUrl: "/query_gateway_capture",
			},
			{
				Name:        "hkrpg-official_cn",
				Title:       "os_usa",
				Type:        "2",
				DispatchUrl: "/query_gateway_capture_cn",
			},
		},
	},
	GameServer: &GameServer{
		MaxPlayer:     -1,
		InnerAddr:     "0.0.0.0",
		Port:          "20041",
		OuterAddr:     "127.0.0.1",
		IsJumpMission: true,
	},
}
