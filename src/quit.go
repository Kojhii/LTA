package main

import 
	"fmt"

func quit() {
	fmt.Print("\033[H\033[2J")
	logo3()
	fmt.Println("\nYou were not strong enough to be a real gangster...\n\n\n ")
}
