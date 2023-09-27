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
	Armor int
	Money int
	Weapon string
	Armorequiped string
	WeaponDamage int
	Levelbar int
}



func (p *Player) Init(Name string, Gang string, Level int, MaxHP int, Hp int, Damage int,Statpoint int,Armor int, Inventory map[string]int,Money int,Weapon string,Armoreequiped string,WeaponDamage int,Levelbar int) {
    p.Name = Name
    p.Gang = Gang
    p.Level = Level
    p.MaxHP = MaxHP
    p.Hp = Hp
	p.Damage = Damage
	p.Statpoint = Statpoint
	p.Inventory = Inventory
	p.Armor = Armor
	p.Money = Money
	p.Weapon = Weapon
	p.Armorequiped = Armoreequiped
	p.WeaponDamage = WeaponDamage
	p.Levelbar = Levelbar

}

type Ennemy struct {
	Hp        int
	MaxHP     int
	Level     int
	Name      string
	Damage    int
	Inventory map[string]int
	Armor int
	Money int
	
}
func (g *Ennemy) Inito(Hp int,MaxHP int,Level int,Name string,Damage int,Inventory map[string]int,Armor int, Money int ){
	g.Name = Name
	g.Hp = Hp 
	g.MaxHP = MaxHP
	g.Level = Level
	g.Damage = Damage
	g.Inventory = Inventory
	g.Armor = Armor
	g.Money = Money
}