package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func Inventory(p *gameEngine.Player) {
	for key, value := range p.Inventory {
		fmt.Println(value,"*",key)

	}
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
