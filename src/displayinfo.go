package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func DisplayInfo(p *gameEngine.Player, s bool) {
	logonstat()
	//print du tableau de stat
	fmt.Println("                                                        _______________________________________________\n                                                       |              ", "   Name :", p.Name, "                 |\n                                                       |              ", "   Level :", p.Level, "                  |\n                                                       |                  HP :", p.Hp, "|", p.MaxHP, "               | \n                                                       |                  Base Damage :", p.Damage, "           |\n                                                       |                  Armor :", p.Armor, "                  |\n                                                       |                  Statpoint :", p.Statpoint, "              |\n                                                       |                  Weapon : ", p.Weapon, "       |\n                                                       |______________________________________________| \n\n ")
	fmt.Println("\n\n                                           _____________________                              ___________________\n                                           |   Put Statpoint    |                             |   Back to Menu   |\n                                           |     (Press 1)      |                             |    (Press 2)     |   \n                                           |____________________|                             |__________________| \n\n\n\n\n\n   ")
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	// scan de l'imput et redirection vers le choix proposé
	var imput string
	fmt.Scan(&imput)

	switch imput {

	case "2":
		fmt.Print("\033[H\033[2J")
		startmenu(p, false)

	case "1":
		fmt.Print("\033[H\033[2J")
		Statpoint(p, false)
	default:
		fmt.Print("\033[H\033[2J")
		DisplayInfo(p, true)
	}
}
