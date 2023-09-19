package main

type gang struct {
	Name              string
	Gang              string
	Level             int
	MaxHP             int
	CurrentLifePoints int
	damage            int
}

func (p *gang) Init(Name string, Gang string, MaxHP int, CurrentLifePoints int, damage int, Level int) {
	p.Name = Name
	p.Gang = Gang
	p.MaxHP = MaxHP
	p.CurrentLifePoints = CurrentLifePoints
	p.damage = damage
	p.Level = Level
}