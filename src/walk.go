package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"time"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"

)


func Walk(p *gameEngine.Player) {
	fmt.Print("\033[H\033[2J")
	//conversion et affichage de l'image
	filePath := "10.webp"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{180, 40}
	flags.Colored = true
	flags.Braille = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)

	if err != nil {
		fmt.Println(err)
	}
	logowalk()
	fmt.Printf("%v\n", asciiArt)
	time.Sleep(2 * time.Second)
	Crimewalk(p)
}

func Crimewalk(p *gameEngine.Player) {
	//init du gangster
	var gangster gameEngine.Ennemy
	gangster.Inito(40, 40, 3, "Gangster", 15, map[string]int{}, 8, 25,25)


	fmt.Print("\033[H\033[2J")
	logocrime()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                    _____________________________________________________\n                                                    |        You were simply walking  when...            |\n                                                    |     SUDDENLY A MENBER OF RIVAL GANG ATTACK YOU     |\n                                                    |____________________________________________________|\n\n\n   ")
	time.Sleep(3 * time.Second)
	combat(p, &gangster)

}