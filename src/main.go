package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func main() {
	var p gameEngine.Player
	p.Init("sam", "zero", 0, 0, 0, 0, 0, 0, map[string]int{}, 0, "punch", "nothing", 0,0)
	logo()
	fmt.Println("                                                                             Start ?                 ")
	fmt.Println("                                                                             Press yes...           ")

	var first string
	fmt.Scan(&first)

	switch first {
	case "lukas":
		fmt.Print("\033[H\033[2J")
		Lukas()
	default:
		fmt.Print("\033[H\033[2J")
		history(&p, false)
	}
}

func startmenu(p *gameEngine.Player, s bool) {

	fmt.Print("\033[H\033[2J")
	levelmax := 10

	if p.Level < levelmax {
		logo()
		fmt.Println("\n\n\n \n ")
		choicelogo(p)
		fmt.Println("\n                                                        _______________________________________________\n                                                       |Map (Press 1)                                 |\n                                                       |                                              |\n                                                       |Character (Press 2)                           |\n                                                       |                                              | \n                                                       |Inventory (Press 3)                           |\n                                                       |                                              |\n                                                       |Quit    (Press 4)                             |\n                                                       |______________________________________________| \n ")
		choicelogolevel(p)
		if s {
			fmt.Println("                                                                       [BAD IMPUT BRO]")
		}
		var imput string
		fmt.Scan(&imput)
		switch imput {

		case "3":
			fmt.Print("\033[H\033[2J")
			Inventoryentrée(p, false)

		case "2":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p, false)

		case "4":
			quit()
		case "1":
			fmt.Print("\033[H\033[2J")
			fmt.Println("                         ")
			Map(p, false)
		default:
			startmenu(p, true)
		}
	} else {
		endgame()
	}
}
