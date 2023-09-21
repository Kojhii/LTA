package main

import (
	"fmt"
)


func Statpoint() {
	pointstatt := 0
	if pointstatt > 0 {
		fmt.Println("  Which stat do you want to put your point ? [HP|DAMAGE] \n", "remaining point")
		fmt.Println("                                             Press 1/2")
		var imput string
		fmt.Scan(&imput)

		switch imput {
		case "1":
			fmt.Println("                Go back to menu")
			fmt.Println("                    (Press 1)")
			var imput string
			fmt.Scan(&imput)
			if imput == "1" {
				fmt.Print("\033[H\033[2J")
				startmenu()
			}

		case "2":
			fmt.Println("                Go back to menu")
			fmt.Println("                    (Press 1)")
			var imput string
			fmt.Scan(&imput)
			if imput == "1" {
				fmt.Print("\033[H\033[2J")
				startmenu()
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
			startmenu()
		}
	}
}
func DisplayInfo() {
	fmt.Println("                              put stat point                            Get back to menu")
	fmt.Println("                                 (Press 1)                                 (Press 2)")
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		fmt.Print("\033[H\033[2J")
		startmenu()

	case "1":
		fmt.Print("\033[H\033[2J")
		Statpoint()
	}
}
