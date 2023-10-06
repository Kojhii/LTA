package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func equipeornotknife(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")
	logo()
	// print du tableau avec la proposition de choix
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |                    Equip knife ?                   |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n   ")
	//si mauvais imput
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	//scan de l'imput 
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "yes":
		//nouvel arme et donc nouveau degat
		p.Weapon = "Knife(+10)"
		p.WeaponDamage = 10
		Shop(p, true, true, false)
	case "no":
		//on renvoie vers le shop
		Shop(p, true, true, false)
	default:
		equipeornotknife(p, true)
	}
}
func equipeornotglock(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")
	logo()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |                    Equip glock ?                   |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n   ")
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "yes":
		//nouvel arme et donc nouveau degat
		p.Weapon = "Glock(+20)"
		p.WeaponDamage = 20
		Shop(p, true, true, false)
	case "no":
		// on renvoie vers le shop
		Shop(p, true, true, false)
	default:
		equipeornotglock(p, true)
	}
}
