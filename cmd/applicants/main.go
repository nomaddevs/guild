package main

import (
	"fmt"
	"strconv"

	"github.com/munsy/guild/config"
	"github.com/munsy/guild/pkg/applicants"
)

func viewApplicant() (bool, int) {
	fmt.Printf("Enter applicant by ID: ")
	var s string
	fmt.Scanln(&s)

	id, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("Couldn't convert entry to ID number: " + err.Error())
		return false, -1
	}

	apps, err := applicants.View(id)

	if nil != err {
		fmt.Println("Couldn't convert entry to ID number: " + err.Error())
		return false, -1
	}

	if 0 == len(apps) {
		fmt.Println("No entries found.")
		return false, -1
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

	return true, id
}

func viewApplicants() {
	apps, err := applicants.ViewAll()

	if nil != err {
		fmt.Println("Couldn't get applicants: " + err.Error())
		return
	}

	fmt.Printf("%s%24s%29s%30s\n", "ID", "BattleTag", "Character", "Email")
	fmt.Printf("%105s\n", "-")

	for i := 0; i < len(apps); i++ {
		app := apps[i]
		fmt.Printf("%d%20s%40s%38s\n", app.BattleID, app.Battletag, app.Character, app.Email)
	}
}

func getBattleID() int {
	fmt.Printf("Enter BattleID: ")

	fmt.Printf("> ")
	var s string
	fmt.Scanln(&s)

	bid, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("Couldn't convert entry to Battle ID: " + err.Error())
		return -1
	}

	return bid
}

func applicantVerdict() {
	ok, id := viewApplicant()

	if ok {
		fmt.Println("1) Accept")
		fmt.Println("2) Reject")
		fmt.Println("3) Cancel")
	}

	fmt.Printf("> ")
	var s string
	fmt.Scanln(&s)

	ans, err := strconv.Atoi(s)

	if nil != err {
		fmt.Println("Couldn't convert entry to option: " + err.Error())
		return
	}

	switch ans {
	case 1:
		fmt.Println("OK. Accepting applicant.")
		err = applicants.Accept(id)

		if nil != err {
			fmt.Println("Failed to accept applicant: " + err.Error())
			return
		}
		fmt.Println("Applicant accepted.")
		return
	case 2:
		fmt.Println("OK. Rejecting applicant.")
		err = applicants.Reject(id)

		if nil != err {
			fmt.Println("Failed to reject applicant: " + err.Error())
			return
		}
		fmt.Println("Applicant rejected.")
		return
	case 3:
		fmt.Println("Aborting...")
		return
	default:
		fmt.Println("Invalid entry.")
		return
	}
}

func purgeApplicant() {
	id := getBattleID()

	if id == -1 {
		return
	}

	err := applicants.Purge(id)

	if nil != err {
		fmt.Println("Failed to purge applicant: " + err.Error())
		return
	}

}

func init() {
	config.Read()
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
		viewApplicants()
		break
	case 2:
		viewApplicant()
		break
	case 3:
		applicantVerdict()
		break
	case 4:
		purgeApplicant()
		break
	default:
		fmt.Println("invalid")
		break
	}
}
