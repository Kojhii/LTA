package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Dealer(p *gameEngine.Player) {
	fmt.Print("\033[H\033[2J")

	logo()

	fmt.Println("\n\n ")
	filePath := "dealos.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{160, 35}
	flags.Colored = false
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

	fmt.Println("\n\n                                                _____________________                              ___________________\n                                                |     Buy stuff      |                             |   Back to MAP    |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "2":
		Map(p)
	case "1":
		Shop(p, true, true)
	}
}

func Shop(p *gameEngine.Player, nomoney bool, already bool) {
	
	fmt.Print("\033[H\033[2J")
	logoshop()
	fmt.Println("\n                                                               [PRESS THE NAME OF WHAT YOU WANT]\n                                                        _______________________________________________\n                                                       |       Objet :        Price :        STOCK :  |\n                                                       |soda(give 10 Health)      20           ", p.Inventory["soda(give 10 health)"], "    |\n                                                       |cotton                    20           ", p.Inventory["cotton"],"    |\n                                                       |vodka                     10           ", p.Inventory["vodka"],"    |\n                                                       |knife(+10 damage)         40           ", p.Inventory["Knife(+10 damage)"], "    |\n                                                       |glock(+30 damage)         100          ", p.Inventory["Glock(+30 damage)"], "    | \n                                                       |ammo(give 20 ammo)        20           ",p.Inventory["ammo(give 20 ammo)"],"    |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |______________________________________________| \n\n                                                                        [Money ", p.Money, "] ")

	if !nomoney {
		fmt.Println("                                                         [You can't buy this , you need more money]")
	}
	if !already {
		fmt.Println("                                                  [You already equip this item,why dit you buy it you dumbass ?]")
	}
	fmt.Println("\n\n                                                                    _____________________                              \n                                                                    |   Go back to Map   |                             \n                                                                    |     (Press 1)      |                             \n                                                                    |____________________|                              \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

	switch imput {
	case "1":
		Map(p)
	case "soda":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["soda(give 10 health)"] = p.Inventory["soda(give 10 health)"] + 1
			Shop(p, true, true)

		} else {
			Shop(p, false, true)
		}
	case "knife":
		if p.Money >= 40 {
			p.Money -= 40
			p.Inventory["knife(+10 damage)"] = p.Inventory["knife(+10 damage)"] + 1
			if p.Weapon != "Knife" {
				equipeornotknife(p)
			} else {
				Shop(p, true, false)
			}

		} else {
			Shop(p, false, true)
		}
	case "glock":
		if p.Money >= 100 {
			p.Money -= 100
			p.Inventory["glock(+30 damage)"] = p.Inventory["lock(+30 damage)"] + 1
			if p.Weapon != "Glock" {
				equipeornotglock(p)
			} else {
				Shop(p, true, false)
			}

		} else {
			Shop(p, false, true)
		}
	case "cotton":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["cotton"] = p.Inventory["cotton"] + 1
			Shop(p, true, true)

		} else {
			Shop(p, false, true)
		}
	case "vodka":
		if p.Money >= 10 {
			p.Money -= 10
			p.Inventory["vodka"] = p.Inventory["vodka"] + 1
			Shop(p, true, true)

		} else {
			Shop(p, false, true)
		}
	case "ammo":
		if p.Money >= 20 {
			p.Money -= 20
			p.Inventory["ammo(give 20 ammo)"] = p.Inventory["ammo(give 20 ammo)"] + 1
			Shop(p, true, true)

		} else {
			Shop(p, false, true)
		}
	}
}
