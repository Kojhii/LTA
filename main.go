package main

import ("fmt"
		"github.com/common-nighthawk/go-figure")


func logo() {
	myFigure := figure.NewFigure("Little Theft Auto ", "", true)
	myFigure.Print()
  }



func main() {
	logo()
	fmt.Println("                                               Start ?                 ")
	fmt.Println("                                               Press Yes...           ")
	var first string
	fmt.Scan(&first)
	if first == "Yes" || first == "yes" {
		fmt.Print("\033[H\033[2J")
		histoire()
	}
}


func startmenu() {
	fmt.Println("character Inventory Quit ")
	fmt.Println("(Press P ) (Press C) (Press Q) ")
	var imput string
	fmt.Scan(&imput)
	if imput == "P" {
	}
	}
	
