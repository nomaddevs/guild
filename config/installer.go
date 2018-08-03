package config

type Installer struct {
	configFilename string
}

func (i *Installer) precheck() bool {
	var c cfg

	if _, err := toml.DecodeFile(i.configFilename, &c); err != nil {
		fmt.Println("Error reading %s:", configFilename)
		fmt.Println("%s", err.Error())
		os.Exit(1)
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

	return true
}

func NewInstaller() *Installer {
	return &Installer{
		ConfigFilename: "config.toml",
	}
}
