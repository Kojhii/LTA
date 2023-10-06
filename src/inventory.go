package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Inventoryentry(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")

	logo()

	//conversation de la photo et affichage
	fmt.Println("\n\n ")
	filePath := "6.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 37}
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
	//si mauvais imput
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	//scan de l'imput et lancement de la prochaine fonction (en rapport au choix)
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		startmenu(p, false)

	case "1":
		Inventory(p, false)
	default:
		Inventoryentry(p, true)
	}
}
func Inventory(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")
	logoinventory()

	//print de l'inventaire
	for key, value := range p.Inventory {
		fmt.Println("\n\n                                                                ", key, ":", value)
	}
	//print des tableau avec choix
	fmt.Println("\n\n                                           _____________________                              ___________________\n                                           |  Use / Equip item  |                             |   Back to Menu   |\n                                           |     (Press 1)      |                             |    (Press 2)     |   \n                                           |____________________|                             |__________________| \n\n\n\n\n\n   ")
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scanln(&imput)
	
	switch imput {
	case "2":
		fmt.Print("\033[H\033[2J")
		startmenu(p, false)
	case "1":
		equiporuse(p, true, true, true, false)
	default:
		Inventory(p, true)
	}
}
