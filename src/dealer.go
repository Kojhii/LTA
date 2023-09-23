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
		Shop(p,true)
	}
}

func Shop(p *gameEngine.Player, s bool) {
	fmt.Print("\033[H\033[2J")
	logoshop()
	fmt.Println("\n                                                               [PRESS THE NAME OF WHAT YOU WANT]\n                                                        _______________________________________________\n                                                       |       Objet :        Price :        STOCK :  |\n                                                       |soda(give 15 Health)      20           ", p.Inventory["Soda(give 15 health)"],"    |\n                                                       |                  HP :" ,p.Hp,"|",p.MaxHP,"               | \n                                                       |                  Damage :", p.Damage, "                |\n                                                       |                  Armor :",p.Armor ,"                 |\n                                                       |                  Statpoint :",p.Statpoint,"              |\n                                                       |                  Money : ",p.Money,"                |\n                                                       |______________________________________________| \n\n ")
	if !s {
		fmt.Println("                                                         [You can't buy this , you need more money]")
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
		Shop(p,true)

		} else {
			Shop(p,false)
		}
	}
}