package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func Statpoint(p *gameEngine.Player) {
	logostatpoint()
	if p.Statpoint > 0 {
		fmt.Println("\n\n\n\n                                                         _____________________________________________________\n                                                         |   In which stat do you want to put a statpoint ?   |\n                                                         |                    [HP | DAMAGE]                   |\n                                                         |____________________________________________________|\n\n\n\n\n   ")
		fmt.Println("                                                _____________________                              ___________________\n                                                |     Put in HP      |                             |   Put in Damage  |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

		var imput string
		fmt.Scan(&imput)

		switch imput {
		case "1":
			fmt.Print("\033[H\033[2J")
			logostatpoint()
			p.Hp += 1
			p.MaxHP += 1
			p.Statpoint -= 1
			fmt.Println("\n\n\n\n\n                                                         ____________________________________________________\n                                                         |               You put a stat in HP                |\n                                                         |          Your new HP stat is : ", p.Hp, "|", p.MaxHP, "          |\n                                                         |___________________________________________________|")
			fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |  Back to character |                             |   Back to Menu   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

			fmt.Scan(&imput)
			switch imput {
			case "1":
				fmt.Print("\033[H\033[2J")
				DisplayInfo(p)
			case "2":
				fmt.Print("\033[H\033[2J")
				startmenu(p)
			}

		case "2":
			fmt.Print("\033[H\033[2J")
			logostatpoint()
			p.Damage += 1
			p.Statpoint -= 1
			fmt.Println("\n\n\n\n\n                                                          ____________________________________________________\n                                                         |          You put a stat in Damage                 |\n                                                         |          Your new damage stat is : ", p.Damage, "           |\n                                                         |___________________________________________________|")
			fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |  Back to character |                             |   Back to Menu   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

			fmt.Scan(&imput)
			switch imput {
			case "1":
				fmt.Print("\033[H\033[2J")
				DisplayInfo(p)
			case "2":
				fmt.Print("\033[H\033[2J")
				startmenu(p)
			}

		}
	} else {
		fmt.Println("\n\n\n\n\n                                                          ____________________________________________________\n                                                         |   Little bitch , you don'y have any statpoint     |\n                                                         |        Go make some crime and come back           |\n                                                         |___________________________________________________|")
		fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |  Back to character |                             |   Back to Menu   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")
		var imput string
		fmt.Scan(&imput)
		switch imput {
		case "1":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p)
		case "2":
			fmt.Print("\033[H\033[2J")
			startmenu(p)
		}

	}
}
