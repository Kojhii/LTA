package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)


func DisplayInfo(p *gameEngine.Player) {
	logonstat()
	fmt.Println("\n\n                                                        _______________________________________________\n                                                       |              ","   Name :" ,p.Name,"                 |\n                                                       |              ","   Level :" ,p.Level, "                  |\n                                                       |                  HP :" ,p.Hp,"|",p.MaxHP,"               | \n                                                       |                  Base  Damage :", p.Damage, "          |\n                                                       |                  Armor :",p.Armor ,"                 |\n                                                       |                  Statpoint :",p.Statpoint,"              |\n                                                       |                  Weapon : ",p.Weapon,"       |\n                                                       |______________________________________________| \n\n ")
	fmt.Println("\n\n                                           _____________________                              ___________________\n                                           |   Put Statpoint    |                             |   Back to Menu   |\n                                           |     (Press 1)      |                             |    (Press 2)     |   \n                                           |____________________|                             |__________________| \n\n\n\n\n\n   ")
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
