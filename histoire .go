package main

import "fmt"

func histoire() {
	fmt.Println("                                                     Welcome to the dark world that is Little Theft Auto \n                                                     a world where you must fight to become the \n                                                     ultimate gang leader in the ruthless city of San Andreas.")
	fmt.Println("                                                                  are you ready ?")
	fmt.Println("                                                                  press   yes/no")
	var imput string
	fmt.Scan(&imput)
	if imput == "yes" || imput == "Yes"{
		fmt.Print("\033[H\033[2J")
		choicegang()
	}
	if imput == "no" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("                                                  Little bitch , you don't have choice")
		fmt.Println("                                                          press Yes or i kill you")
		var secondimput string
		fmt.Scan(&secondimput)
		if secondimput == "yes" || secondimput == "Yes" {
			fmt.Print("\033[H\033[2J")
			fmt.Println("THIS IS THE GOOD SPIRIT")
			choicegang()
		}
	}
}