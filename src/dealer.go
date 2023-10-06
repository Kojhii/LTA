package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Dealer(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")

	logo()
	//import de l'image
	fmt.Println("\n\n ")
	filePath := "3.jpeg"
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

	//print du tableau
	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |     Buy stuff      |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")
	
	//si jamais mauvais imput
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	//scan du imput et redirection vers le menu choisis
	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p, false)
	case "1":
		Shop(p, true, true, false)
	default:
		Dealer(p, true)
	}
}

func Shop(p *gameEngine.Player, nomoney bool, already bool, s bool) {

	fmt.Print("\033[H\033[2J")
	logoshop()
	fmt.Println("\n                                                               [PRESS THE NAME OF WHAT YOU WANT]\n                                                        _______________________________________________\n                                                       |       Objet :        Price :        STOCK :  |\n                                                       |soda(give 10 Health)      20           ", p.Inventory["soda(give 10 health)"], "    |\n                                                       |cotton                    20           ", p.Inventory["cotton"], "    |\n                                                       |vodka                     10           ", p.Inventory["vodka"], "    |\n                                                       |knife(+10 damage)         40           ", p.Inventory["knife(+10 damage)"], "    |\n                                                       |glock(+20 damage)         100          ", p.Inventory["Glock(+20 damage)"], "    | \n                                                       |ammo(give 20 ammo)        20           ", p.Inventory["ammo(give 20 ammo)"], "    |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |______________________________________________| \n\n                                                                        [Money ", p.Money, "] ")

	//si pas d'argent
	if !nomoney {
		fmt.Println("                                                         [You can't buy this , you need more money]")
	}
	// si objet deja equipé
	if !already {
		fmt.Println("                                                  [You already equip this item,why dit you buy it you dumbass ?]")
	}
	fmt.Println("\n\n                                                                    _____________________                              \n                                                                    |   Go back to Map   |                             \n                                                                    |     (Press 1)      |                             \n                                                                    |____________________|                              \n\n\n\n\n\n   ")
	
	//si mauvais imput
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	//retour map
	case "1":
		Map(p, false)
	// achat soda si argent > prix
	case "soda":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["soda(give 10 health)"] = p.Inventory["soda(give 10 health)"] + 1
			Shop(p, true, true, false)

		} else {
			Shop(p, false, true, false)
		}
	case "knife":
		if p.Money >= 40 {
			p.Money -= 40
			p.Inventory["knife(+10 damage)"] = p.Inventory["knife(+10 damage)"] + 1
			//si jamais l'arme equipé est pas le couteau , on propose de l'equipé
			if p.Weapon != "Knife" {
				equipeornotknife(p, false)
			} else {
				Shop(p, true, false, false)
			}

		} else {
			Shop(p, false, true, false)
		}
	case "glock":
		if p.Money >= 100 {
			p.Money -= 100
			p.Inventory["glock(+20 damage)"] = p.Inventory["lock(+20 damage)"] + 1
			//si jamais l'arme equipé est pas le glock , on propose de l'equipé
			if p.Weapon != "Glock" {
				equipeornotglock(p, false)
			} else {
				Shop(p, true, false, false)
			}

		} else {
			Shop(p, false, true, false)
		}
	case "cotton":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["cotton"] = p.Inventory["cotton"] + 1
			Shop(p, true, true, false)

		} else {
			Shop(p, false, true, false)
		}
	case "vodka":
		if p.Money >= 10 {
			p.Money -= 10
			p.Inventory["vodka"] = p.Inventory["vodka"] + 1
			Shop(p, true, true, false)

		} else {
			Shop(p, false, true, false)
		}
	case "ammo":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["ammo(give 20 ammo)"] = p.Inventory["ammo(give 20 ammo)"] + 1
			Shop(p, true, true, false)

		} else {
			Shop(p, false, true, false)
		}
	default:
		Shop(p, true, true, true)
	}
}
