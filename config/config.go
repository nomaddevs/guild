package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	DB  *MariaDBConfig `json:"dbconfig"`
	TLS *TLSconfig     `json:"tlsconfig"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open("./config.json")
	if nil != err {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if nil != err {
		return nil, err
	}

	var cfg *Config

	json.Unmarshal(data, &cfg)

	return cfg, nil
}

func (cfg *Config) Save(filename string) error {
	data, err := json.Marshal(cfg)
	if nil != err {
		return err
	}

	ioutil.WriteFile("./config.json", data, 0644)
	return nil

}
