package main 

import "fmt"

type Person struct {
	Name              string
	Gang              string
	Level             int
	MaxHP         int
	CurrentLifePoints int
	damage int
}	

func choicegang() {
	crips := "| THE CRIPS |\n|This organization often|\n|engages in robbery illegal|\n|drug trafficking and|\n|other criminal activities.|\n|This gang is in deep conflict with the Bloods|\n|maxHP : 10|\n|damage : 7|"
	fmt.Printf("It's all very well to want to play the bandit, but why don't you start by choosing a gang? ")
	fmt.Printf("      \n                                  press Yes \n")
	var choice string
	fmt.Scan(&choice)
	if choice == "yes" || choice == "Yes" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("_____________")
		fmt.Println(crips)
		fmt.Printf("What is your choice ? \n")
		var sndchoice string 
		fmt.Scan(&sndchoice)
		if sndchoice == "CRIPS" || sndchoice == "crips" || sndchoice == "Crips" {
			fmt.Print("\033[H\033[2J")
			println("Good luck on your adventure bro      ")
			startmenu()
		

		}
	}
}




