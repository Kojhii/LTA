package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func history(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")
	logo()
	//print du debut de l'histoire
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                         _____________________________________________________\n                                                         |Welcome to the dark world that is Little Theft Auto |\n                                                         |                    are you ready ?                 |\n                                                         |                    press yes/no                    |\n                                                         |____________________________________________________|\n\n\n   ")
	//si mauvais imput
	if s {
		fmt.Println("                                                                             [BAD IMPUT BRO]")
	}
	//scan de l'imput
	var imput string
	fmt.Scan(&imput)

	switch imput {
	//si imput = oui , lancement de la fonction choicegang
	case "yes":
		fmt.Print("\033[H\033[2J")
		logo()
		choicegang(p)

	case "no":
		//si imput = non , print d'un autre choix (obligé de dire oui)
		fmt.Print("\033[H\033[2J")
		logo()
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                          ____________________________________________________\n                                                         |          Little bitch , you don'y have choice     |\n                                                         |               Press yes or i kill you             |\n                                                         |___________________________________________________|")
		var secondimput string
		fmt.Scan(&secondimput)
		switch secondimput {
		default:
			fmt.Print("\033[H\033[2J")
			logo()
			fmt.Println("                                                                       [THIS IS THE GOOD SPIRIT, keep up]")
			choicegang(p)
		}

	default:
		history(p, true)
	}
}
