package main

import (
	"fmt"
	"strconv"

	"github.com/munsy/guild/pkg/applicants"
)

func viewApplicant() {
	fmt.Printf("Enter applicant by ID: ")
	var s string
	fmt.Scanln(s)

	id, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("Couldn't convert entry to ID number: " + err.Error())
		return
	}

	apps, err := applicants.View(id)

	if nil != err {
		fmt.Println("Couldn't convert entry to ID number: " + err.Error())
		return
	}

	for i := 0; i < len(apps); i++ {
		app := apps[i]
		fmt.Println("BattleID: " + strconv.Itoa(app.BattleID))
		fmt.Println("Battletag: " + app.Battletag)
		fmt.Println("Character: " + app.Character)
		fmt.Println("Email: " + app.Email)
		fmt.Println("RealName: " + app.RealName)
		fmt.Println("Location: " + app.Location)
		fmt.Println("Age: " + app.Age)
		fmt.Println("Gender: " + app.Gender)
		fmt.Println("ComputerSpecs: " + app.ComputerSpecs)
		fmt.Println("PreviousGuilds: " + app.PreviousGuilds)
		fmt.Println("ReasonsLeavingGuilds: " + app.ReasonsLeavingGuilds)
		fmt.Println("WhyJoinThisGuild: " + app.WhyJoinThisGuild)
		fmt.Println("References: " + app.References)
	}
}

func viewApplicants() {
	apps, err := applicants.ViewAll()

	if nil != err {
		fmt.Println("Couldn't get applicants: " + err.Error())
		return
	}

	fmt.Printf("%20s%20s%20s%20s\n", "ID", "BattleTag", "Character", "Email")
	fmt.Println("----------------------------------------------------------")

	for i := 0; i < len(apps); i++ {
		app := apps[i]
		fmt.Printf("%20s%20s%20s%20s\n", app.BattleID, app.Battletag, app.Character, app.Email)
	}
}

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
		viewApplicants()
		break
	case 2:
		fmt.Println("case 2")
		viewApplicant()
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
