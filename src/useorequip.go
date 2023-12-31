package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func equiporuse(p *gameEngine.Player, already bool, alcoolic bool, equiped bool, s bool) {
	fmt.Print("\033[H\033[2J")
	logoinventory()
	//print de 'linventaire
	for key, value := range p.Inventory {
		fmt.Println("\n\n                                                                ", key, ":", value)
	}
	fmt.Println("\n                                                            [PUT THE NAME OF THE OBJECT YOU WANT TO USE] \n\n                                                                         ___________________\n                                                                        |   Back to Menu   |\n                                                                        |    (Press 1)     |   \n                                                                        |__________________| \n\n\n   ")
	//si deja equipé
	if !already {
		fmt.Println("                                                             [ OBJECT ALREADY EQUIPED /Can't be used]")
	}
	//si on utilise la vodka
	if !alcoolic {
		fmt.Println("                                                             [ Are you  fucking alcoholic bro ? ]")
	}
	//si on l'equipe
	if !equiped {
		fmt.Println("                                                              [        Object equiped /used     ]")
	}
	// si mauvais imput
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "1":
		startmenu(p, false)

	case "vest":
		if p.Armorequiped == "vest" {
			equiporuse(p, false, true, true, false)
		} else {
			p.Armorequiped = "vest"
			p.Armor = 10
		}
	case "soda":
		if p.Inventory["soda(give 10 health)"] > 0 {
			p.Hp += 10
			if p.Hp > p.MaxHP {
				p.Hp = p.MaxHP
			}
			p.Inventory["soda(give 10 health)"] -= 1
			equiporuse(p, true, true, false, false)
		} else {
			equiporuse(p, false, true, true, false)
		}

	case "vodka":
		equiporuse(p, true, false, true, false)
	case "cotton":
		equiporuse(p, false, true, true, false)
	case "molotov":
		equiporuse(p, false, true, true, false)
	case "knife":
		p.Weapon = "knife(+10)"
		p.WeaponDamage = 10
		equiporuse(p, true, true, false, false)
	case "glock":
		p.Weapon = "glock(+20)"
		p.WeaponDamage = 20
		equiporuse(p, true, true, false, false)
	default:
		equiporuse(p, true, true, true, true)
	}
}
