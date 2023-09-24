package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func equiporuse(p *gameEngine.Player, already bool, alcoolic bool,equiped bool) {
	fmt.Print("\033[H\033[2J")
	logoinventory()
	for key, value := range p.Inventory {
		fmt.Println("\n\n                                                                ", key, ":", value)
	}
	fmt.Println("\n                                                            [PUT THE NAME OF THE OBJECT YOU WANT TO USE] \n\n                                                                         ___________________\n                                                                        |   Back to Menu   |\n                                                                        |    (Press 1)     |   \n                                                                        |__________________| \n\n\n   ")
	if !already {
		fmt.Println("                                                             [ OBJECT ALREADY EQUIPED /Can't be used]")
	}
	if !alcoolic {
		fmt.Println("                                                             [ Are you  fucking alcoholic bro ? ]")
	}
	if !equiped {
		fmt.Println("                                                              [        Object equiped /used     ]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "1":
		startmenu(p)

	case "vest":
		if p.Armorequiped == "vest" {
			equiporuse(p, false, true,true)
		} else {
			p.Armorequiped = "vest"
			p.Armor = 10
		}
	case "soda":
		if p.Inventory["soda(give 10 health)"] >0 {
			p.Hp += 10
		if p.Hp > p.MaxHP {
			p.Hp = p.MaxHP
		}
		p.Inventory["soda(give 10 health)"] -= 1
		equiporuse(p, true, true,false)
		} else {
			equiporuse(p,false,true,true)
		}

	case "vodka":
		equiporuse(p, true, false,true)
	case "cotton":
		equiporuse(p, false ,true,true)
	case "molotov":
		equiporuse(p,false,true,true)
	case "knife":
		p.Weapon = "knife(+10)"
		equiporuse(p,true,true,false)
	case "glock":
		p.Weapon = "glock(+20)"
		equiporuse(p,true,true,false)
	}
}
