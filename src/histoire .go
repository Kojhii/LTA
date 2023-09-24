package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func history(p *gameEngine.Player) {
	logo()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |Welcome to the dark world that is Little Theft Auto |\n                                                         |                    are you ready ?                 |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n   ")
	var imput string
	fmt.Scan(&imput)
	if imput == "yes" || imput == "Yes" {
		fmt.Print("\033[H\033[2J")
		logo()
		choicegang(p)
	}
	if imput == "no" {
		fmt.Print("\033[H\033[2J")
		logo()
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                          ____________________________________________________\n                                                         |          Little bitch , you don'y have choice     |\n                                                         |               Press yes or i kill you             |\n                                                         |___________________________________________________|")
		var secondimput string
		fmt.Scan(&secondimput)
		if secondimput == "yes" || secondimput == "Yes" {
			fmt.Print("\033[H\033[2J")
			logo()
			fmt.Println("                                                                       [THIS IS THE GOOD SPIRIT, keep up]")
			choicegang(p)
		}
	}
}
