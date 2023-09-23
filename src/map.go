package main

import ("github.com/TheZoraiz/ascii-image-converter/aic_package"
		"fmt"
)

func Map() {
	logo()
	fmt.Println("\n\n ")
	filePath := "cartelta.webp"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{60, 35}
	flags.Colored = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}