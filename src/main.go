package main

import (
	"fmt"
	"github.com/Kojhii/LTA/gameEngine"
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
	var g gameEngine.Player
	g.Init()

	logo()
	fmt.Println("                                               Start ?                 ")
	fmt.Println("                                               Press Yes...           ")
	var first string
	fmt.Scan(&first)

	if first == "Yes" || first == "yes" {
		fmt.Print("\033[H\033[2J")
		history()
	}

	if first == "lukas" {
		fmt.Print("\033[H\033[2J")
		logo2()
	}
}

func startmenu() {
	levelmax := 30
	for i := 0; i <= levelmax; {
		fmt.Println("character Inventory Quit ")
		fmt.Println("(Press 1 ) (Press 2) (Press 3) ")
		var imput string
		fmt.Scan(&imput)
		switch imput {
		case "2":
			fmt.Print("\033[H\033[2J")
			Inventory()

		case "1":
			fmt.Print("\033[H\033[2J")
			DisplayInfo()
		}
	}
	endgame()
}
