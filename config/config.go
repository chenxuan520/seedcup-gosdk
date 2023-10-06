package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MapSize int32  `json:"map_size"`
	Ip      string `json:"ip"`
	Port    uint32 `json:"port"`
}

func InitConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	return config, nil
}
