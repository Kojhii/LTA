package main

import ("github.com/TheZoraiz/ascii-image-converter/aic_package"
		"fmt"
		"github.com/Kojhii/LTA/src/gameEngine"
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
		Shop(p,true,true)
	}
}

func Shop(p *gameEngine.Player, nomoney bool,already bool) {
	
	fmt.Print("\033[H\033[2J")
	
	logoshop()
	fmt.Println("\n                                                               [PRESS THE NAME OF WHAT YOU WANT]\n                                                        _______________________________________________\n                                                       |       Objet :        Price :        STOCK :  |\n                                                       |soda(give 15 Health)      20           ", p.Inventory["Soda(give 15 health)"],"    |\n                                                       |knife(+10 damage)         40           ", p.Inventory["Knife(+10 damage)"],"    |\n                                                       |glock(+30 damage)         100          ", p.Inventory["Glock(+30 damage)"],"    | \n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |                                              |\n                                                       |______________________________________________| \n\n                                                                        [Money ",p.Money,"] ")
	
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
		p.Inventory ["Soda(give 15 health)"] = p.Inventory["Soda(give 15 health)"] + 1
		Shop(p,true,true)

		} else {
			Shop(p,false,true)
		}
	case "knife":
		if p.Money >= 40 {
		p.Money -= 40
		p.Inventory ["Knife(+10 damage)"] = p.Inventory["Knife(+10 damage)"] + 1
		if p.Weapon != "Knife" {
			equipeornotknife(p)
		} else {
			Shop(p,true,false)
		}

		} else {
			Shop(p,false,true)
		}
	case "glock":
		if p.Money >= 100 {
		p.Money -= 100
		p.Inventory ["Glock(+30 damage)"] = p.Inventory["Glock(+30 damage)"] + 1
		if p.Weapon != "Glock" {
			equipeornotglock(p)
		} else {
			Shop(p,true,false)
		}

		} else {
			Shop(p,false,true)
		}	
	}
}