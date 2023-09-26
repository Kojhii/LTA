package main

import (
	"github.com/common-nighthawk/go-figure"

	"github.com/Kojhii/LTA/src/gameEngine"
)

func logo() {
	myFigure := figure.NewFigure("                Little Theft Auto ", "", true)
	myFigure.Print()
}

func logo2() {
	myFigure := figure.NewFigure("Virginie ", "", true)
	myFigure.Print()
}
func logo3() {
	myFigure := figure.NewFigure("           THANKS FOR PLAYING ", "", true)
	myFigure.Print()
}
func logocrips() {
	myFigure := figure.NewFigure("                             CRIPS ", "", true)
	myFigure.Print()
}
func logobloods() {
	myFigure := figure.NewFigure("                           BLOODS ", "", true)
	myFigure.Print()
}
func logolatinos() {
	myFigure := figure.NewFigure("                         LATINOS ", "", true)
	myFigure.Print()
}

func choicelogo(p *gameEngine.Player) {
	if p.Gang == "Crips" {
		logocrips()
	}
	if p.Gang == "Bloods" {
		logobloods()
	}
	if p.Gang == "Latinos" {
		logolatinos()
	}
}
func logonewbie() {
	myFigure := figure.NewFigure("                            ROOKIE ", "", true)
	myFigure.Print()
}
func choicelogolevel(p *gameEngine.Player) {
	if p.Level <= 2 {
		logonewbie()
	}
}
func logonstat() {
	myFigure := figure.NewFigure("                               STAT ", "", true)
	myFigure.Print()
}
func logostatpoint() {
	myFigure := figure.NewFigure("                          STATPOINT ", "", true)
	myFigure.Print()
}
func logoinventory() {
	myFigure := figure.NewFigure("                       Inventory ", "", true)
	myFigure.Print()
}
func logoshop() {
	myFigure := figure.NewFigure("                                SHOP", "", true)
	myFigure.Print()
}

func logocraft() {
	myFigure := figure.NewFigure("                              Craft", "", true)
	myFigure.Print()
}

func logocar() {
	myFigure := figure.NewFigure("         You are taking the car", "", true)
	myFigure.Print()
}
func logocrime(){
	myFigure := figure.NewFigure("                             CRIME", "", true)
	myFigure.Print()
}