package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func Statpoint(p *gameEngine.Player) {

	if p.Statpoint > 0 {
		fmt.Println("  Which stat do you want to put your point ? [HP|DAMAGE] \n", "remaining point")
		fmt.Println("                                             Press 1/2")
		var imput string
		fmt.Scan(&imput)

		switch imput {
		case "1":
			fmt.Print("\033[H\033[2J")
			p.Hp += 1
			p.MaxHP += 1
			p.Statpoint -= 1
			fmt.Println("Your new HP stat is :", p.Hp, "|", p.MaxHP)
			fmt.Println("                Go back to menu")
			fmt.Println("                    (Press 1)")
			var imput string
			fmt.Scan(&imput)
			if imput == "1" {

				fmt.Print("\033[H\033[2J")
				startmenu(p)
			}

		case "2":
			fmt.Print("\033[H\033[2J")
			p.Damage += 1
			p.Statpoint -= 1
			fmt.Println("Your new damage stat is:", p.Damage)
			fmt.Println("                Go back to menu")
			fmt.Println("                    (Press 1)")
			var imput string
			fmt.Scan(&imput)
			if imput == "1" {
				fmt.Print("\033[H\033[2J")
				startmenu(p)
			}
		}
	} else {
		fmt.Println("What are you doing bro ? you don't have any point stat \n go train little bitch")
		fmt.Println("                Go back to menu")
		fmt.Println("                    (Press 1)")
		var imput string
		fmt.Scan(&imput)
		switch imput {

		case "1":
			fmt.Print("\033[H\033[2J")
			startmenu(p)
		}
	}
}

func DisplayInfo(p *gameEngine.Player) {
	fmt.Println("STAT ", "\nName :", p.Name, "\nGang :", p.Gang, "\nLevel :", p.Level, "\nYour base damage :", p.Damage, "\nLife :", p.Hp, "|", p.MaxHP, "\nStatpoint", p.Statpoint)
	fmt.Println("                              put stat point                            Get back to menu")
	fmt.Println("                                 (Press 1)                                 (Press 2)")
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		fmt.Print("\033[H\033[2J")
		startmenu(p)

	case "1":
		fmt.Print("\033[H\033[2J")
		Statpoint(p)
	}
}
