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
		Shop(p)
	}
}

func Shop(p *gameEngine.Player) {
	fmt.Print("\033[H\033[2J")
	logoshop()
	fmt.Println("\n\n                                                        _______________________________________________\n                                                       |  Objet   |    Price | Number to press        |\n                                                       |              ","   Level :" ,p.Level, "                  |\n                                                       |                  HP :" ,p.Hp,"|",p.MaxHP,"               | \n                                                       |                  Damage :", p.Damage, "                |\n                                                       |                  Armor :",p.Armor ,"                 |\n                                                       |                  Statpoint :",p.Statpoint,"              |\n                                                       |                  Money : ",p.Money,"                |\n                                                       |______________________________________________| \n\n ")
	fmt.Println("\n\n                                                                    _____________________                              \n                                                                    |   Go back to Map   |                             \n                                                                    |     (Press 1)      |                             \n                                                                    |____________________|                              \n\n\n\n\n\n   ")
	
}