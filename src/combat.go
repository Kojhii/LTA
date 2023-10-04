package main

import (
	"fmt"
	"github.com/Kojhii/LTA/src/gameEngine"
	"time"
)

func combat(p *gameEngine.Player, cop *gameEngine.Ennemy) {
	fmt.Print("\033[H\033[2J")
	logofight()

	fmt.Println("\n\n\n                                                           ", p.Name, "                                     ", cop.Name, "\n                                                    ", "LEVEL :", p.Level, "                                ", "LEVEL :", cop.Level, "\n                                                       ", "HP :", p.Hp, "|", p.MaxHP, "                             ", "HP :", cop.Hp, "|", cop.MaxHP, " \n                                                     Armor :", p.Armor, "                                 Armor :", cop.Armor, "\n                                                    Damage :", p.Damage, "                               Damage :", cop.Damage)
	fmt.Println("\n\n\n\n\n\n\n                                                _____________________                              ___________________\n                                                |      attack        |                             |   use a object   |\n                                                |     (Press 1)      |                             |    (Press 2)     |   \n                                                |____________________|                             |__________________| \n\n\n\n\n\n   ")

	var imput string
	fmt.Scan(&imput)

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
	cop.Hp -= p.Damage
	cop.Hp -= p.WeaponDamage
	cop.Hp += cop.Armor
	if cop.Hp > 0 {
		fmt.Println("                                                         [You juste made", p.WeaponDamage+p.Damage-cop.Armor, "damage to", cop.Name, "| The ennemie has now ", cop.Hp, "HP]")
		time.Sleep(800 * time.Millisecond)
		ennemyturn(p, cop)
	} else {
		fmt.Println("                                                                     [   YOU JUST GAVE HIM THE FATAL BLOW   ]")
		time.Sleep(800 * time.Millisecond)
		Win(p, cop)
	}

}

func ennemyturn(p *gameEngine.Player, cop *gameEngine.Ennemy) {
	p.Hp -= cop.Damage
	p.Hp += p.Armor
	time.Sleep(1 * time.Second)
	time.Sleep(2 * time.Second)
	if p.Hp <= 0 {
		time.Sleep(800 * time.Millisecond)
		Death(p)
	} else {
		fmt.Println("                                                         [Cop juste made", cop.Damage-p.Armor, "damage to you| You have now ", p.Hp, "HP]")
		time.Sleep(800 * time.Millisecond)
		combat(p, cop)
	}

}

func Useobjectcombat(p *gameEngine.Player, cop *gameEngine.Ennemy, s bool) {
	fmt.Println("                                                              [Choose which object you want to use]")
	fmt.Println("                                                               molotov :", p.Inventory["molotov"],"soda",p.Inventory["soda(give 10 health)"])

	if s {
		fmt.Println("                                                                           [BAD IMPUT BRO]")
	}
	var imput string
	fmt.Scan(&imput)

	switch imput {
	default:
		Useobjectcombat(p, cop, true)
	case "soda": {
		if p.Inventory["soda(give 10 health)"] >0 {
			p.Hp +=10
			if p.Hp > p.MaxHP {
				p.Hp = p.MaxHP
			}
			ennemyturn(p,cop)
		} else {
			fmt.Println("                                                                    [YOU DONT HAVE ANY SODA]")
			time.Sleep(2 * time.Second)
			combat(p,cop)
		}
	}
	case "molotov":
		if p.Inventory["molotov"] <= 0 {
			fmt.Println("                                                                    [YOU DONT HAVE ANY MOLOTOV]")
			time.Sleep(2 * time.Second)
			combat(p, cop)
		} else {
			p.Inventory["molotov"] -= 1
			cop.Armor -= 20
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
