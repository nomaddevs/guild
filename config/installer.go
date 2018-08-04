package config

import (
	"fmt"
	"strings"

	"github.com/munsy/guild/database"
)

func Install() bool {
	// Read in values to see if any exist
	err := Read()
	if nil != err {
		fmt.Printf("Error trying to read %s.\n", TOMLFile)
		return false
	}

	if "" == Key {
		fmt.Printf("Key is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set Key: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for Key: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set Key: %s\n", err.Error())
				return false
			}
			Key = answer
		}
	}
	if "" == Secret {
		fmt.Printf("Secret is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set Secret: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for Secret: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set Secret: %s\n", err.Error())
				return false
			}
			Secret = answer
		}
	}
	if "" == RedirectURL {
		fmt.Printf("RedirectURL is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set RedirectURL: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for RedirectURL: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set RedirectURL: %s\n", err.Error())
				return false
			}
			RedirectURL = answer
		}
	}
	if "" == Addr {
		fmt.Printf("Addr is not set. Would you like to set it now (y/n, default n)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set Addr: %s\n", err.Error())
			return false
		}

		if strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for Addr: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set Addr: %s\n", err.Error())
				return false
			}
			Addr = answer
		}
	}
	if "" == CertFile {
		fmt.Printf("CertFile is not set. Would you like to set it now (y/n, default n)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set CertFile: %s\n", err.Error())
			return false
		}

		if strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for CertFile: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set CertFile: %s\n", err.Error())
				return false
			}
			CertFile = answer
		}
	}
	if "" == KeyFile {
		fmt.Printf("KeyFile is not set. Would you like to set it now (y/n, default n)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set KeyFile: %s\n", err.Error())
			return false
		}

		if strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for KeyFile: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set KeyFile: %s\n", err.Error())
				return false
			}
			KeyFile = answer
		}
	}
	if "" == DBUsername {
		fmt.Printf("DBUsername is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBUsername: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBUsername: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBUsername: %s\n", err.Error())
				return false
			}
			DBUsername = answer
		}
	}
	if "" == DBPassword {
		fmt.Printf("DBPassword is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBPassword: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBPassword: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBPassword: %s\n", err.Error())
				return false
			}
			DBPassword = answer
		}
	}
	if "" == DBUnixsocketpath {
		fmt.Printf("DBUnixsocketpath is not set. Would you like to set it now (y/n, default n)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBUnixsocketpath: %s\n", err.Error())
			return false
		}

		if strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBUnixsocketpath: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBUnixsocketpath: %s\n", err.Error())
				return false
			}
			DBUnixsocketpath = answer
		}
	}
	if "" == DBHost {
		fmt.Printf("DBHost is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBHost: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBHost: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBHost: %s\n", err.Error())
				return false
			}
			DBHost = answer
		}
	}
	if "" == DBPort {
		fmt.Printf("DBPort is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBPort: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBPort (default 3306): ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBPort: %s\n", err.Error())
				return false
			}
			if answer == "" {
				DBPort = "3306"
			}
		}
	}
	if "" == DBName {
		fmt.Printf("DBName is not set. Would you like to set it now (y/n, default y)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBName: %s\n", err.Error())
			return false
		}

		if answer == "" || strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBName: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBName: %s\n", err.Error())
				return false
			}
			DBName = answer
		}
	}
	if "" == DBCharset {
		fmt.Printf("DBCharset is not set. Would you like to set it now (y/n, default n)? ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if nil != err {
			fmt.Printf("Error trying to set DBCharset: %s\n", err.Error())
			return false
		}

		if strings.Contains(answer, "y") {
			answer = ""
			fmt.Printf("New value for DBCharset: ")
			_, err = fmt.Scanln(&answer)
			if nil != err {
				fmt.Printf("Error trying to set DBCharset: %s\n", err.Error())
				return false
			}
			DBCharset = answer
		}
	}

	fmt.Printf("Setting up %s...\n", TOMLFile)

	err = Write()

	if nil != err {
		fmt.Printf("Error creating %s: %s\n", TOMLFile, err.Error())
		return false
	}

	fmt.Printf("Succeeded writing file '%s'\n", TOMLFile)

	fmt.Printf("Creating database '%s'\n", DBName)

	err = installDB()

	if nil != err {
		fmt.Printf("Error creating %s: %s\n", DBName, err.Error())
		return false
	}

	fmt.Printf("Succeeded creating database '%s'\n", DBName)

	return true
}

func installDB() error {
	db := &database.MariaDB{
		Username:       DBUsername,
		Unixsocketpath: DBUnixsocketpath,
		Password:       DBPassword,
		Host:           DBHost,
		Port:           DBPort,
		Database:       DBName,
		Charset:        DBCharset,
	}

	return db.Install()
}
