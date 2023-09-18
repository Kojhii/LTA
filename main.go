package main 

import "fmt"

type Personne struct {
	name string
}


func main() {	fmt.Println("                     Commencer ?                 ")
				fmt.Println("                     appuyer sur C...            ")
}


func startmenu() {
	fmt.Println("Personnage                         Inventaire                         Quitter ")
	fmt.Println("(appuyer sur P )                  (appuyer sur I)                  (appuyer sur Q) ")	
}