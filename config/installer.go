package config

import (
	"errors"
	"fmt"
)

func Install() bool {
	fmt.Printf("Creating %s...\n")

	err := Write()

	if nil != err {
		fmt.Printf("Error creating %s: %s\n", TOMLFile, err.Error())
		return false
	}

	fmt.Printf("Write succeeded.\n")

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
	return errors.New("database installer logic not implemented")
}
