package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func equipeornotknife(p *gameEngine.Player,) {
	fmt.Print("\033[H\033[2J")
	logo()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |                    Equip knife ?                   |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "yes":
		p.Weapon = "Knife(+10)"
		Shop(p,true,true)
	case "no":
		Shop(p,true,true)
	}
}
func equipeornotglock(p *gameEngine.Player,) {
	fmt.Print("\033[H\033[2J")
	logo()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |                    Equip glock ?                   |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "yes":
		p.Weapon = "Glock(+30)"
		Shop(p,true,true)
	case "no":
		Shop(p,true,true)
	}
}