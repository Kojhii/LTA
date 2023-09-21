package gameEngine

type Player struct {
	Hp        int
	MaxHP     int
	Level     int
	Name      string
	Statpoint int
	Gang      string
	Damage    int
	Inventory map[string]int
}

type Engine struct {
	CharacterT Player
}

func (p *Player) Init(Name string, Gang string, Level int, MaxHP int, Hp int, Damage int,Statpoint int, Inventory map[string]int) {
    p.Name = Name
    p.Gang = Gang
    p.Level = Level
    p.MaxHP = MaxHP
    p.Hp = Hp
	p.Damage = Damage
	p.Statpoint = Statpoint
	p.Inventory = Inventory

}