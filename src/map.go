package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Map(p *gameEngine.Player) {

	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")

	filePath := "ltatest.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Braille = true
	flags.Dimensions = []int{170, 40}
	flags.Colored = false
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	fmt.Println("                          _____________________                    ___________________                     __________________                    __________________\n                          |  Commit a crime    |                   |   Go to dealer   |                   |    Workshop     |                    |   Go to menu    |\n                          |     (Press 1)      |                   |    (Press 2)     |                   |    (Press 3)    |                    |    (Press 4)    |   \n                          |____________________|                   |__________________|                   |_________________|                    |_________________| \n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "4":
		startmenu(p)
	case "2":
		Dealer(p)
	}
}
