package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)
func combat(p *gameEngine.Player, cop *gameEngine.Ennemy){
	fmt.Print("\033[H\033[2J")
	logofight()
	fmt.Println("\n\n\n                                                           ", p.Name, "                                     ",cop.Name,"\n                                                           ", p.Level, "                                     ",cop.Level,"\n                                                           ", p.Hp,"|",p.MaxHP, "                                 ",cop.Hp,"|",cop.MaxHP)
	fmt.Println("\n\n\n\n\n\n\n                                                _____________________                              ___________________\n                                                |      attack        |                             |   use a object   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

}