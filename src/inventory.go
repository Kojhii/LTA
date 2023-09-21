package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func Inventory(p *gameEngine.Player) {
	fmt.Println("Your inventory :", p.Inventory)
	fmt.Println("              Get back to menu")
	fmt.Println("      		 (Press 1)")

	var imput string
	fmt.Scanln(&imput)
	switch imput {
	case "1":
		fmt.Print("\033[H\033[2J")
		startmenu(p)
	}
}
