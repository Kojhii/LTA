package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
)

func choicegang(p *gameEngine.Player) {

	crips := "|              THE CRIPS                       |\n|This organization often                       |\n|engages in robbery illegal                    |\n|drug trafficking and                          |\n|other criminal activities.                    |                                           [Press 1 to be a fucking Crips bro] \n|This gang is in deep conflict with the Bloods | \n|HP : 15                                       |\n|damage : 12                                   |\n|______________________________________________|"

	bloods := "|              The BLOODS                      |\n|Street gang                                   |\n|involved in drugs, theft, and murder          |\n|other criminal activities                     |\n|This gang is in deep conflict with the Crips  |                                           [Press 2 to be a Bloods , the rival of the crips]\n|HP : 32                                       |\n|damage : 10                                   |\n|______________________________________________|"

	latinos := "|              THE LATINOS                     |\n|A gang made up mainly of                      |\n|Latinos from South America                    |\n|who play on their influence and dangerousness |\n|Looking for a family who can                  |\n|protect you as you kill?                      |                                           [Press 3 to be a LATINOS] \n|drug trafficking and                          |\n|other criminal activities.                    |\n|This gang is for you.                         |\n|HP : 28                                       |\n|damage : 14                                   |\n|______________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n  "

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                       ____________________________________________________________________________________________ \n                                       |It's all very well to want to play the bandit, but why don't you start by choosing a gang?|\n                                       |                                      Press Yes                                           |")
	fmt.Println("\n                                       |__________________________________________________________________________________________|\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n  ")

	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "yes":
		fmt.Print("\033[H\033[2J")

		fmt.Println("_______________________________________________")           
		fmt.Println(crips)
		fmt.Println("_______________________________________________")  

		fmt.Println(bloods)
		fmt.Println("_______________________________________________")
		fmt.Println(latinos)

		var sndchoice string
		fmt.Scan(&sndchoice)

		switch sndchoice {
		case "1":
			fmt.Print("\033[H\033[2J")
			println("                                        [Good luck on your adventure bro  , the future of the crips's gang is in your hands]   ")
			p.Init("Sam", "Crips", 1, 30, 30, 12, 1,10, map[string]int{"Bulletproof vest (+10 armor)" : 1},40,"punch(+00)")
			startmenu(p)

		case "2":
			fmt.Print("\033[H\033[2J")
			println("                                        [Good luck on your adventure bro ,the future of the bloods's gang is in your hands]     ")
			p.Init("Sam", "Bloods", 1, 32,32, 10, 1,10, map[string]int{ "Bulletproof vest (+10 armor) " : 1},40,"punch(+00)")
			startmenu(p)

		case "3":
			fmt.Print("\033[H\033[2J")
			println("                                        [Good luck on your adventure bro , the future of the latinos's gang is in your hands]     ")
			p.Init("Sam", "Latinos", 1, 28, 28, 14, 1,10, map[string]int{"Bulletproof vest (+10 armor) " : 1},40,"punch(+00)")
			startmenu(p)

		}
	}
}
