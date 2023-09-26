package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Workshopmenu(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")
	filePath := "4.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 35}
	flags.Colored = true
	flags.Braille = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |       Enter        |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p, false)

	case "1":
		Workshop(p, false)
	default:
		Workshopmenu(p, true)
	}
}

func Workshop(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")
	filePath := "2.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 35}
	flags.Colored = true
	flags.Braille = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |    Craft item      |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p, false)

	case "1":
		Craftitem(p, true, false)
	default:
		Workshop(p, true)
	}
}

func Craftitem(p *gameEngine.Player, donthavematerial bool, s bool) {
	fmt.Print("\033[H\033[2J")
	logocraft()
	fmt.Println("\n                                                          [PRESS THE NAME OF WHAT YOU WANT TO CRAFT]\n                                                        _______________________________________________\n                                                       |    Objet :       Material :        STOCK :   |\n                                                       |    molotov       vodka+cotton         ", p.Inventory["molotov"], "    |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              | \n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |______________________________________________| \n\n ")
	fmt.Println("\n\n                                                                     ___________________\n                                                                    |   Back to Menu   |\n                                                                    |    (Press 1)     |   \n                                                                    |__________________| \n\n\n   ")
	if !donthavematerial {
		fmt.Println("                                                         [You don't have the material required bro]")
	}
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "molotov":
		if p.Inventory["cotton"] > 0 || p.Inventory["vodka"] > 0 {
			p.Inventory["molotov"] += 1
			p.Inventory["cotton"] -= 1
			p.Inventory["vodka"] -= 1
			Craftitem(p, true, false)
		} else {
			Craftitem(p, false, false)
		}
	case "1":
		Map(p, false)
	default:
		Craftitem(p, true, true)
	}
}
