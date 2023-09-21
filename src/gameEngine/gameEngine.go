package gameEngine

type Player struct {
	Hp        int
	MaxHP     int
	Level     int
	Name      string
	Statpoint int
	Gang      string
	Damage    int
}

type Engine struct {
	CharacterT Player
}

func (g *Player) Init() {
	Crips := Player{
		10,
		10,
		1,
		"sam",
		1,
		"Crips",
		7,
	}
	Bloods := Player{
		7,
		7,
		1,
		"sam",
		1,
		"Bloods",
		10,
	}
	Latinos := Player{
		8,
		8,
		1,
		"sam",
		1,
		"Latinos",
		8,
	}
	type ennemie struct {
		Name   string
		HP     int
		MaxHP  int
		Damage int
		Gang   string
	}
}
