package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

type Loader struct {
	key    string
	secret string
}

func (l *Loader) WriteTOML(key, token, region, locale string) error {
	var inputs = Config{
		key,
		region,
		locale,
	}

	var buffer bytes.Buffer

	encoder := toml.NewEncoder(&buffer)

	err := encoder.Encode(inputs)

	if nil != err {
		return err
	}

	return ioutil.WriteFile("config.toml", buffer.Bytes(), 0644)
}

func (l *Loader) ReadTOML(filename string) {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
	}
}
