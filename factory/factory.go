package factory

import "fmt"

const (
	SwordType = 1
	BowType = 2
	GunType = 3
)

type Weapon interface {
	Use()
}

type Sword struct {
	Damage  int
	Range   int
	Attack  string
}

func (s *Sword) Use() {
	s.Attack = "Slash!"
	fmt.Printf("Sword attacks with %d damage and %d range, %s", s.Damage, s.Range, s.Attack)
}

type Bow struct {
	Damage  int
	Range   int
	Attack  string
}

func (b *Bow) Use() {
	b.Attack = "Twang!"
	fmt.Printf("Bow attacks with %d damage and %d range, %s", b.Damage, b.Range, b.Attack)
}

type Gun struct {
	Damage    int
	Range     int
	ReloadTime int
	Attack    string
}

func (g *Gun) Use() {
	g.Attack = "Bang!"
	fmt.Printf("Gun attacks with %d damage and %d range, reload time is %d sec, %s", g.Damage, g.Range, g.ReloadTime, g.Attack)
}

type WeaponFactory struct{}

func (w *WeaponFactory) CreateWeapon(t int) (Weapon, error) {
	switch t {
	case 1:
		return &Sword{Damage: 10, Range: 1}, nil
	case 2:
		return &Bow{Damage: 7, Range: 10}, nil
	case 3:
		return &Gun{Damage: 15, Range: 20, ReloadTime: 2}, nil
	default:
		return nil, fmt.Errorf("weapon %d not implemented", t)
	}
}