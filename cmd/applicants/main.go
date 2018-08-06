package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Select action: ")
	fmt.Println("0) Quit")
	fmt.Println("1) View all applicants")
	fmt.Println("2) View applicant by ID")
	fmt.Println("3) Accept/Reject an applicant")
	fmt.Println("4) Purge an applicant")
	fmt.Printf("> ")

	var s string
	fmt.Scanln(&s)

	i, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("Invalid entry")
		return
	}

	switch i {
	case 0:
		return
	case 1:
		fmt.Println("case 1")
		break
	case 2:
		fmt.Println("case 2")
		break
	case 3:
		fmt.Println("case 3")
		break
	case 4:
		fmt.Println("case 4")
		break
	default:
		fmt.Println("invalid")
		break
	}
}
