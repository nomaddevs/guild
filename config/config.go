package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	Key              string
	Secret           string
	RedirectURL      string
	Addr             string
	CertFile         string
	KeyFile          string
	DBUsername       string
	DBPassword       string
	DBUnixsocketpath string
	DBHost           string
	DBPort           string
	DBName           string
	DBCharset        string
)

type cfg struct {
	key              string
	secret           string
	redirectURL      string
	addr             string
	certFile         string
	keyFile          string
	dbuname          string
	dbpasswd         string
	dbunixsocketpath string
	dbhost           string
	dbport           string
	dbname           string
	dbcharset        string
}

func Write(filename string) error {
	c := &cfg{
		key:              Key,
		secret:           Secret,
		redirectURL:      RedirectURL,
		addr:             Addr,
		certFile:         CertFile,
		keyFile:          KeyFile,
		dbuname:          DBUsername,
		dbpasswd:         DBPassword,
		dbunixsocketpath: DBUnixsocketpath,
		dbhost:           DBHost,
		dbport:           DBPort,
		dbname:           DBName,
		dbcharset:        DBCharset,
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
	RedirectURL = c.redirectURL
	Addr = c.addr
	CertFile = c.certFile
	KeyFile = c.keyFile
	DBUsername = c.dbuname
	DBPassword = c.dbpasswd
	DBUnixsocketpath = c.dbunixsocketpath
	DBHost = c.dbhost
	DBPort = c.dbport
	DBName = c.dbname
	DBCharset = c.dbcharset
}
