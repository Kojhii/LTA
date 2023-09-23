package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)


func main() {
	var p gameEngine.Player
	p.Init("sam", "zero", 0, 0, 0, 0, 0,0, map[string]int{})
	logo()
	fmt.Println("                                                                             Start ?                 ")
	fmt.Println("                                                                             Press Yes...           ")
	var first string
	fmt.Scan(&first)

	if first == "Yes" || first == "yes" {
		fmt.Print("\033[H\033[2J")
		history(&p)
	}

	if first == "lukas" {
		fmt.Print("\033[H\033[2J")
		logo2()
	}
}

func startmenu(p *gameEngine.Player) {
	levelmax := 10
	if p.Level < levelmax {
		
		logo()
		fmt.Println("\n\n\n\n\n\n ")
		choicelogo(p)
		fmt.Println("\n\n  \n                                                       _______________________________________________\n                                                       |character (Press 1)                           |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              | \n                                                       |Inventory (Press 2)                           |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |Quit    (Press 3)                             |\n                                                       |______________________________________________| \n\n ")
		choicelogolevel(p)
		var imput string
		fmt.Scan(&imput)
		
		switch imput {
		
		case "2":
			fmt.Print("\033[H\033[2J")
			Inventory(p)

		
		case "1":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p)
		
		case "3":
			quit()
		}
	} else {
		endgame()
	}
}
