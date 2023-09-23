package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func Inventory(p *gameEngine.Player) {
	logoinventory()
	for key, value := range p.Inventory {
		fmt.Println("\n\n                                                                ",key,":",value)
	}
	fmt.Println("\n\n                                                                         ____________________                              \n                                                                         |   Back to Menu   |\n                                                                         |    (Press 1)     |   \n                                                                         |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scanln(&imput)
	switch imput {
	case "1":
		fmt.Print("\033[H\033[2J")
		startmenu(p)
	}
}
