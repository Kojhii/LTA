package main

import "github.com/Kojhii/LTA/src/gameEngine"

//fonction pour associer le niveau du joueur en rapport avec avec la barre d'xp
func Levelbar(p *gameEngine.Player) {
	if p.Levelbar == 0 {
		p.Level = 1
	}
	if p.Levelbar == 15 {
		p.Level = 2
		p.Statpoint += 1
	}
	if p.Levelbar == 30 {
		p.Level = 3
		p.Statpoint += 1
	}
	if p.Levelbar == 50 {
		p.Level = 4
		p.Statpoint += 1
	}
	if p.Levelbar == 80 {
		p.Level = 5
		p.Statpoint += 1
	}
	if p.Levelbar == 115 {
		p.Level = 6
		p.Statpoint += 1
	}
	if p.Levelbar == 150 {
		p.Level = 7
		p.Statpoint += 1
	}
	if p.Levelbar == 190 {
		p.Level = 8
		p.Statpoint += 1
	}
	if p.Levelbar == 240 {
		p.Level = 9
		p.Statpoint += 1
	}
	if p.Levelbar == 300 {
		p.Level = 10
		p.Statpoint += 1
	}
}
