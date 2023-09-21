package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func logo4() {
	myFigure := figure.NewFigure("THANKS FOR PLAYING ", "", true)
	myFigure.Print()
}
func quit() {
	fmt.Print("\033[H\033[2J")
	logo4()
	fmt.Println("You were not strong enough to be a real gangster...")
}
