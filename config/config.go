package config

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var (
	TOMLFile = "../../config.toml"

	Debug            bool
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
}

func Write() error {
	c := &cfg{
		Key:              Key,
		Secret:           Secret,
		RedirectURL:      RedirectURL,
		Addr:             Addr,
		CertFile:         CertFile,
		KeyFile:          KeyFile,
		DBUsername:       DBUsername,
		DBPassword:       DBPassword,
		DBUnixsocketpath: DBUnixsocketpath,
		DBHost:           DBHost,
		DBPort:           DBPort,
		DBName:           DBName,
		DBCharset:        DBCharset,
	}

	var b bytes.Buffer

	e := toml.NewEncoder(&b)

	err := e.Encode(c)

	if nil != err {
		return err
	}

	return ioutil.WriteFile(TOMLFile, b.Bytes(), 0644)
}

func Read() error {
	var c cfg

	if _, err := toml.DecodeFile(TOMLFile, &c); err != nil {
		return err
	}

	Key = c.Key
	Secret = c.Secret
	RedirectURL = c.RedirectURL
	Addr = c.Addr
	CertFile = c.CertFile
	KeyFile = c.KeyFile
	DBUsername = c.DBUsername
	DBPassword = c.DBPassword
	DBUnixsocketpath = c.DBUnixsocketpath
	DBHost = c.DBHost
	DBPort = c.DBPort
	DBName = c.DBName
	DBCharset = c.DBCharset

	return nil
}

func Dump() {
	fmt.Printf("Key: %s\n", Key)
	fmt.Printf("Secret: %s\n", Secret)
	fmt.Printf("RedirectURL: %s\n", RedirectURL)
	fmt.Printf("Addr: %s\n", Addr)
	fmt.Printf("CertFile: %s\n", CertFile)
	fmt.Printf("KeyFile: %s\n", KeyFile)
	fmt.Printf("DBUsername: %s\n", DBUsername)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBUnixsocketpath: %s\n", DBUnixsocketpath)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBCharset: %s\n", DBCharset)
}
