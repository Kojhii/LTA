package main

import "fmt"

func quit() {
	//print de la fin de jeux (en quittant)
	fmt.Print("\033[H\033[2J")
	logo3()
	fmt.Println("\n[You were not strong enough to be a real gangster...]\n\n\n ")
}
