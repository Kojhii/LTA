package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func choicegang(p *gameEngine.Player) {

	crips := "| THE CRIPS |\n|This organization often|\n|engages in robbery illegal|\n|drug trafficking and|\n|other criminal activities.|\n|This gang is in deep conflict with the Bloods|\n|maxHP : 10|\n|damage : 7|"

	bloods := "|The BLOODS|\n|street gang|\n|involved in drugs, theft, and murder|\n| other criminal activities.|\n|This gang is in deep conflict with the Crips|\n|maxHP : 7|\n|damage : 10|"

	latinos := "|THE LATINOS |\n|A gang made up mainly of Latinos from South America, who play on their influence and dangerousness|\n| Looking for a family who can protect you as you kill?|\n|drug trafficking and|\n|other criminal activities.|\n|This gang is for you.|\n|maxHP : 8|\n|damage : 8|"

	fmt.Printf("It's all very well to want to play the bandit, but why don't you start by choosing a gang? ")
	fmt.Printf("      \n                                  press Yes \n")

	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "yes":
		fmt.Print("\033[H\033[2J")
		fmt.Println("_____________")

		fmt.Println(crips)
		fmt.Println("_____________")

		fmt.Println(bloods)
		fmt.Println("_____________")
		fmt.Println(latinos)

		fmt.Println("What is your choice ? \n", "1/2/3")
		var sndchoice string
		fmt.Scan(&sndchoice)

		switch sndchoice {
		case "1":
			fmt.Print("\033[H\033[2J")
			println("Good luck on your adventure bro  , the future of the crips's gang is in your hands   ")
			p.Init("Sam", "Crips", 1, 10, 10, 7, 1, map[string]int{"Knife": 1})
			startmenu(p)

		case "2":
			fmt.Print("\033[H\033[2J")
			println("Good luck on your adventure bro ,the future of the bloods's gang is in your hands     ")
			p.Init("Sam", "Bloods", 1, 7, 7, 10, 1, map[string]int{"Knife": 1})
			startmenu(p)

		case "3":
			fmt.Print("\033[H\033[2J")
			println("Good luck on your adventure bro , the future of the latinos's gang is in your hands     ")
			p.Init("Sam", "Latinos", 1, 8, 8, 8, 1, map[string]int{"Knife": 1})
			startmenu(p)

		}
	}
}
