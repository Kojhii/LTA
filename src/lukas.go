package main

import (
	"fmt"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Lukas() {
	//easter egg avec lukas, n'en parlons pas
	filePath := "8.jpeg"
	flags := aic_package.DefaultFlags()
	flags.Braille = true
	flags.Dimensions = []int{70, 40}
	flags.Colored = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}
