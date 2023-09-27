package main

import (
	"fmt"
	"time"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func Win(p *gameEngine.Player,cop * gameEngine.Ennemy) {
	fmt.Print("\033[H\033[2J")
	logoWIN()
	fmt.Println("                                                                         [     YOU JUST WIN     ]             ")
	fmt.Println("\n\n\n\n\n                                                         ____________________________________________________\n                                                         |                 +",cop.Money,"money                        |\n                                                         |                 + 15 xp                           |\n                                                         |           You are going back to the map           |\n                                                         |___________________________________________________|")
	time.Sleep(4 * time.Second)
	p.Money = p.Money + cop.Money
	preclevel := p.Level
	p.Levelbar +=15
	Levelbar(p)
	if preclevel <= p.Level{
		fmt.Print("\033[H\033[2J")
		logolevelup()
		fmt.Println("\n\n\n\n\n\n\n                                                                     [  YOU HAVE GAIN A LEVEL ]")
		fmt.Println("                                                                      [   NEW LEVEL IS :",p.Level,"   ]")
	}
	time.Sleep(2 * time.Second)
	Map(p,false)
}
func Death() {
	fmt.Print("\033[H\033[2J")
	logoDEATH()
	fmt.Println("YOU JUST DIE")
}
