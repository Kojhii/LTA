package main

import ("github.com/TheZoraiz/ascii-image-converter/aic_package"
		"fmt"
)

func Map() {
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
	fmt.Println("                          _____________________                    ___________________                   ___________________                    ___________________\n                          |  Commit a crime    |                   |   Go to dealer   |\n                          |     (Press 1)      |                   |    (Press 2)     |   \n                          |____________________|                   |__________________| \n\n\n\n   ")

}