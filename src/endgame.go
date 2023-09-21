package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func logo3() {
	myFigure := figure.NewFigure("THANKS FOR PLAYING ", "", true)
	myFigure.Print()
}
func endgame() {
	fmt.Println("Wow...you really did it bro...\n you are now the king of all the gang of the USA...")
	logo3()
	fmt.Println("-Zeljko Mitrovic / Ryan Avaro")
}
