package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Workshopmenu(p *gameEngine.Player) {
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")
	filePath := "garage.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 35}
	flags.Colored = true
	flags.Braille = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |       Enter        |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p)
	
	case "1":
		Workshop(p)
	}
}

func Workshop(p *gameEngine.Player){
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")
	filePath := "atelier.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 35}
	flags.Colored = true
	flags.Braille = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |    Craft item      |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p)
	
	case "1":
		Craftitem(p)
	}
}

func Craftitem(p *gameEngine.Player){
	fmt.Print("\033[H\033[2J")
	logocraft()
}