package main

import "fmt"

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
			name = " Sam"
			gang = "Crips"
			maxHP = 10
			hp = 10
			damage = 7
			level = 1
			pointstat = 1
			startmenu()
		}
	}
}

func Statpoint() {
	if pointstat > 0 {
		fmt.Println("  Which stat do you want to put your point ? [HP|DAMAGE] \n", "remaining point", pointstat)
	var imput string
	fmt.Scan(&imput)

	if imput == "HP" || imput == "Hp" || imput == "hp" {
		maxHP = maxHP + 1
		hp = hp + 1
		pointstat = pointstat - 1
		fmt.Print("\033[H\033[2J")
		fmt.Println("you put a point in HP, your new maxHP stat is", maxHP)
		fmt.Println("                Go back to menu")
		fmt.Println("                    (Press M)")
		var imput string
		fmt.Scan(&imput)
		if imput == "M" {
			fmt.Print("\033[H\033[2J")
			startmenu()
		}
	}
	if imput == "DAMAGE" || imput == "Damage" || imput == "damage" {
		damage = damage + 1
		pointstat = pointstat - 1
		fmt.Print("\033[H\033[2J")
		fmt.Println("you put a point in damage, your new damage stat is", damage)
		fmt.Println("                Go back to menu")
		fmt.Println("                    (Press M)")
		var imput string
		fmt.Scan(&imput)
		if imput == "M" {
			fmt.Print("\033[H\033[2J")
			startmenu()
			}
		}
	} else {
		fmt.Println("What are you doing bro ? you don't have any point stat \n go train little bitch")
		fmt.Println("                Go back to menu")
		fmt.Println("                    (Press M)")
		var imput string
		fmt.Scan(&imput)
		if imput == "M" {
			fmt.Print("\033[H\033[2J")
			startmenu()
			}
	}
}
func DisplayInfo() {
	fmt.Println(name, "\n", gang, "\n", "Level :", level, "\n HP :", hp, "|", maxHP, "\n Your damage : ", damage, " \n Your Statpoint", pointstat)
	fmt.Println("                              put stat point                            Get back to menu")
	fmt.Println("                                 (Press P)                                 (Press M)")
	var imput string
	fmt.Scan(&imput)
	if imput == "M" {
		fmt.Print("\033[H\033[2J")
		startmenu()
	}
	if imput == "P" {
		fmt.Print("\033[H\033[2J")
		Statpoint()
	}
}