package main
// import de la fonction fmt et de notre gameengine
import (
	"fmt"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func main() {
	// creation de notre variable player et ininiation
	var p gameEngine.Player
	p.Init("sam", "zero", 0, 0, 0, 0, 0, 0, map[string]int{}, 0, "punch", "nothing", 0, 0)
	
	//Print du logo et texte de base
	logo()
	fmt.Println("                                                                             Start ?                 ")
	fmt.Println("                                                                             Press yes...           ")

	//scan de l'imput du joueur
	var first string
	fmt.Scan(&first)

	//easter egg lukas ou lancement de l'histoire, "fmt.Print("\033[H\033[2J")" = print pour nettoyer le terminal
	switch first {
	case "lukas":
		fmt.Print("\033[H\033[2J")
		Lukas()
	default:
		fmt.Print("\033[H\033[2J")
		history(&p, false)
	}
}

func startmenu(p *gameEngine.Player, s bool) {

	fmt.Print("\033[H\033[2J")
	levelmax := 10

	// Condition de jeux tant que le joueur n'est pas niveau 10
	if p.Level <= levelmax {

		logo()
		
		fmt.Println("\n\n\n \n ")
		
		choicelogo(p)
		
		fmt.Println("\n                                                        _______________________________________________\n                                                       |Map (Press 1)                                 |\n                                                       |                                              |\n                                                       |Character (Press 2)                           |\n                                                       |                                              | \n                                                       |Inventory (Press 3)                           |\n                                                       |                                              |\n                                                       |Quit    (Press 4)                             |\n                                                       |______________________________________________| \n ")
		
		choicelogolevel(p)
		
		// si jamais le joueur se trompe d'imput
		if s {
			fmt.Println("                                                                       [BAD IMPUT BRO]")
		}
		
		// scan de l'imput pour rediriger vers le menu approprié
		var imput string
		fmt.Scan(&imput)
		switch imput {

		case "3":
			fmt.Print("\033[H\033[2J")
			Inventoryentry(p, false)

		case "2":
			fmt.Print("\033[H\033[2J")
			DisplayInfo(p, false)

		case "4":
			quit()
		case "1":
			fmt.Print("\033[H\033[2J")
			fmt.Println("                         ")
			Map(p, false)
		default:
			startmenu(p, true)
		}
	// si jamais le joueur est niveau 10 , il a donc gagné
	} else {
		endgame()
	}
}
