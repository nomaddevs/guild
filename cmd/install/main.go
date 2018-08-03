package main

import (
	"fmt"

	"github.com/munsy/guild/config"
)

func main() {
	fmt.Println("Installing...")

	if config.Install() {
		fmt.Println("Install successful.")
		return
	}

	fmt.Println("Install FAILED.")
	fmt.Println("Check the error log for more information.")
}
