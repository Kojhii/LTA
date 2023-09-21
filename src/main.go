package main

import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/common-nighthawk/go-figure"
)

func logo() {
	myFigure := figure.NewFigure("Little Theft Auto ", "", true)
	myFigure.Print()
}

func logo2() {
	myFigure := figure.NewFigure("Virginie ", "", true)
	myFigure.Print()
}

func main() {
	var p gameEngine.Player
	p.Init("sam", "zero", 0, 0, 0, 0, 0, map[string]int{})
	logo()
	fmt.Println("                                               Start ?                 ")
	fmt.Println("                                               Press Yes...           ")
	var first string
	fmt.Scan(&first)

	if first == "Yes" || first == "yes" {
		fmt.Print("\033[H\033[2J")
		history(&p)
	}

	if first == "lukas" {
		fmt.Print("\033[H\033[2J")
		logo2()
	}
}

func startmenu(p *gameEngine.Player) {
	levelmax := 30
	if p.Level < levelmax {
		fmt.Println("character Inventory Quit ")
		fmt.Println("(Press 1 ) (Press 2) (Press 3) ")
		var imput string
		fmt.Scan(&imput)
		switch imput {
		case "2":
			fmt.Print("\033[H\033[2J")
			Inventory(p)

		case "1":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p)
		case "3":
			quit()
		}
	} else {
		endgame()
	}
}
