package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func commitacrimevalidation(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n ")

	filePath := "7.jpg"
	flags := aic_package.DefaultFlags()
	flags.Braille = true
	flags.Dimensions = []int{170, 40}
	flags.Colored = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", asciiArt)
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |    Take the car    |                             |      Walk        |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":

	case "1":
		voiture(p)
	default:
		commitacrimevalidation(p, true)

	}
}
