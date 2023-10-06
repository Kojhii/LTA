package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"time"
)

func combat(p *gameEngine.Player, cop *gameEngine.Ennemy) {
	fmt.Print("\033[H\033[2J")
	logofight()
	//affichage des stat des deux combattants

	fmt.Println("\n\n\n                                                           ", p.Name, "                                     ", cop.Name, "\n                                                    ", "LEVEL :", p.Level, "                                ", "LEVEL :", cop.Level, "\n                                                       ", "HP :", p.Hp, "|", p.MaxHP, "                             ", "HP :", cop.Hp, "|", cop.MaxHP, " \n                                                     Armor :", p.Armor, "                                 Armor :", cop.Armor, "\n                                                    Damage :", p.Damage, "                               Damage :", cop.Damage)
	fmt.Println("\n\n\n\n\n\n\n                                                _____________________                              ___________________\n                                                |      attack        |                             |   use a object   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)
	// choix entre attaquer , ou utilser un objet
	switch imput {
	default:
		combat(p, cop)
	case "1":
		attack(p, cop)
	case "2":
		Useobjectcombat(p, cop, false)

	}
}

func attack(p *gameEngine.Player, cop *gameEngine.Ennemy) {
	// Attaque = vie ennemie - degat de base - degat d'arme + armure
	cop.Hp -= p.Damage
	cop.Hp -= p.WeaponDamage
	cop.Hp += cop.Armor
	// si il est vivant , affichage des degat , pause , puis tour ennemie
	if cop.Hp > 0 {
		fmt.Println("                                                         [You juste made", p.WeaponDamage+p.Damage-cop.Armor, "damage to", cop.Name, "| The ennemie has now ", cop.Hp, "HP]")
		time.Sleep(800 * time.Millisecond)
		ennemyturn(p, cop)
	} else {
		// si il est mort , annonce de sa mort , pause , puis fonction de victoire
		fmt.Println("                                                                     [   YOU JUST GAVE HIM THE FATAL BLOW   ]")
		time.Sleep(800 * time.Millisecond)
		Win(p, cop)
	}

}

func ennemyturn(p *gameEngine.Player, cop *gameEngine.Ennemy) {
	//meme fonctionnement que l'attaque mais inversé , sans les degat d'arme qui sont compris dans les degat de base de l'ennemie
	p.Hp -= cop.Damage
	p.Hp += p.Armor
	time.Sleep(1 * time.Second)
	time.Sleep(2 * time.Second)
	// si on est mort
	if p.Hp <= 0 {
		time.Sleep(800 * time.Millisecond)
		Death(p)
	// si on est vivant , relancement de la fonction combat pour reproposer un choix (attaque ou objet)
	} else {
		//affichage des degat reçu
		fmt.Println("                                                         [Cop juste made", cop.Damage-p.Armor, "damage to you| You have now ", p.Hp, "HP]")
		time.Sleep(800 * time.Millisecond)
		combat(p, cop)
	}

}

func Useobjectcombat(p *gameEngine.Player, cop *gameEngine.Ennemy, s bool) {
	//choix de l'objet a choisir et affichage de l'inventaire
	fmt.Println("                                                              [Choose which object you want to use]")
	fmt.Println("                                                               molotov :", p.Inventory["molotov"],"soda",p.Inventory["soda(give 10 health)"])
	//si mauvais imput 
	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	default:
		Useobjectcombat(p, cop, true)
	//si on utilise soda 
	case "soda": {
		//Si on en a un => vie = vie + 10
		if p.Inventory["soda(give 10 health)"] >0 {
			p.Hp +=10
			if p.Hp > p.MaxHP {
				p.Hp = p.MaxHP
			}
			ennemyturn(p,cop)
		//si jamais on en a pas , relancement du choix
		} else {
			fmt.Println("                                                                    [YOU DONT HAVE ANY SODA]")
			time.Sleep(2 * time.Second)
			combat(p,cop)
		}
	}
	case "molotov":
		//si jamais on en a pas , relancement du choix
		if p.Inventory["molotov"] <= 0 {
			fmt.Println("                                                                    [YOU DONT HAVE ANY MOLOTOV]")
			time.Sleep(2 * time.Second)
			combat(p, cop)
		} else {
		// armure ennemie -20
			p.Inventory["molotov"] -= 1
			cop.Armor -= 20
			//si on a detruit l'armure
			if cop.Armor <= 0 {
				cop.Armor = 0
				fmt.Println("                                                         [You just destroyed the armor of the ennemy]")
				time.Sleep(1 * time.Second)
				combat(p, cop)
			} else {
				combat(p, cop)
			}
		}
	}
}
