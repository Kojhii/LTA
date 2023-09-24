package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)


func main() {
	var p gameEngine.Player
	p.Init("sam", "zero", 0, 0, 0, 0, 0,0, map[string]int{},0,"punch")
	logo()
	fmt.Println("                                                                             Start ?                 ")
	fmt.Println("                                                                             Press yes...           ")
	
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
	fmt.Print("\033[H\033[2J")
	levelmax := 10
	
	if p.Level < levelmax {
		
		logo()
		fmt.Println("\n\n\n\n\n\n ")
		choicelogo(p)
		fmt.Println("\n\n  \n                                                       _______________________________________________\n                                                       |Map (Press 1)                                 |\n                                                       |                                              |\n                                                       |Character (Press 2)                           |\n                                                       |                                              | \n                                                       |Inventory (Press 3)                           |\n                                                       |                                              |\n                                                       |Quit    (Press 4)                             |\n                                                       |______________________________________________| \n\n ")
		choicelogolevel(p)
		var imput string
		fmt.Scan(&imput)
		
		switch imput {
		
		case "3":
			fmt.Print("\033[H\033[2J")
			Inventory(p)

		
		case "2":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p)
		
		case "4":
			quit()
		case "1":
			fmt.Print("\033[H\033[2J")
			fmt.Println("                         ")
			Map(p)
		}
	} else {
		endgame()
	}
}
