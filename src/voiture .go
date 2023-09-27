package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"time"
)

func voiture(p *gameEngine.Player) {
	fmt.Print("\033[H\033[2J")
	
	filePath := "9.jpg"
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
	logocar()
	fmt.Printf("%v\n", asciiArt)
	time.Sleep(2 * time.Second)
	crimevoiture(p)
}
	

func crimevoiture(p *gameEngine.Player) {
	var cop gameEngine.Ennemy
	cop.Inito(20,20,2,"Cop",10,map[string]int{},4,15)
	fmt.Print("\033[H\033[2J")
	logocrime()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n                                                    _____________________________________________________\n                                                    |        wow you exceeded the speed limit            |\n                                                    |        A cop is trying to arrest you, KILL HIM     |\n                                                    |____________________________________________________|\n\n\n   ")
	time.Sleep(3 * time.Second)
	combat(p,&cop)

}

