package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	Key      string
	Secret   string
	Addr     string
	CertFile string
	KeyFile  string
)

type cfg struct {
	key      string
	secret   string
	addr     string
	certFile string
	keyFile  string
	dbuname  string
	dbpasswd string
}

func Write(filename string) error {
	c := &cfg{
		key:      Key,
		secret:   Secret,
		addr:     Addr,
		certFile: CertFile,
		keyFile:  KeyFile,
		dbuname:  DBUsername,
		dbpasswd: DBPassword,
	}

	var b bytes.Buffer

	e := toml.NewEncoder(&b)

	err := e.Encode(c)

	if nil != err {
		return err
	}

	return ioutil.WriteFile(filename, b.Bytes(), 0644)
}

func Read(filename string) {
	var c *cfg

	if _, err := toml.DecodeFile(filename, &c); err != nil {
		fmt.Println("Error reading %s:", filename)
		fmt.Println("%s", err.Error())
		os.Exit(1)
	}

	Key = c.key
	Secret = c.secret
	Addr = c.addr
	CertFile = c.certFile
	KeyFile = c.keyFile
	DBUsername = c.dbuname
	DBPassword = c.dbpasswd
}
