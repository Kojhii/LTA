package main

import "fmt"

type Personne struct {
	name string
}

func main() {
	fmt.Println("                     Start ?                 ")
	fmt.Println("                     Press C...           ")
	var first string
	fmt.Scan(&first)
	if first == "C" {
		startmenu()
	}
}

func startmenu() {
	fmt.Println("Personnage                         Inventaire                         Quitter ")
	fmt.Println("(Press P )               	    (Press C)                         (Press Q) ")
}
