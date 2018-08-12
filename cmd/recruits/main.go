package main

import (
	"fmt"
	"strconv"

	"github.com/munsy/guild/config"
	//"github.com/munsy/guild/pkg/recruits"
)

func init() {
	config.Read()
}

func setClassRecruitCount() {
	fmt.Println("Classes: ")
	fmt.Println("1) DeathKnight")
	fmt.Println("2) DemonHunter")
	fmt.Println("3) Druid")
	fmt.Println("4) Hunter")
	fmt.Println("5) Mage")
	fmt.Println("6) Monk")
	fmt.Println("7) Paladin")
	fmt.Println("8) Priest")
	fmt.Println("9) Rogue")
	fmt.Println("10) Shaman")
	fmt.Println("11) Warlock")
	fmt.Println("12) Warrior")
	fmt.Printf("Enter class by number: ")

	var s string
	fmt.Scanln(&s)

	i, err := strconv.Atoi(s)

	if nil != err || i > 12 || i < 1 {
		fmt.Println("Invalid entry")
		return
	}

	fmt.Printf("Enter new recruit count: ")

	s = ""
	fmt.Scanln(&s)

	j, err := strconv.Atoi(s)

	if nil != err || j < 0 {
		fmt.Println("Invalid entry")
		return
	}

}

func main() {
	fmt.Println("Options:")
	fmt.Println("0) Quit")
	fmt.Println("1) Set the number of desired recruits for a class")
	fmt.Printf("> ")

	var s string
	fmt.Scanln(&s)

	i, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("invalid")
		return
	}

	switch i {
	case 0:
		return
	case 1:
		setClassRecruitCount()
		return
	default:
		fmt.Println("invalid")
		return
	}
}
