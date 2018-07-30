package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Key    string
	Secret string
}

func WriteTOML(key, secret string) error {
	var inputs = Config{
		key,
		secret,
	}

	var buffer bytes.Buffer

	encoder := toml.NewEncoder(&buffer)

	err := encoder.Encode(inputs)

	if nil != err {
		return err
	}

	return ioutil.WriteFile("config.toml", buffer.Bytes(), 0644)
}

func ReadTOML(filename string) *Config {
	var config *Config
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		fmt.Println("Error reading %s:", filename)
		fmt.Println("%s", err.Error())
		os.Exit(1)
	}
	return config
}
