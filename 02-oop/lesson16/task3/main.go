package main

import "fmt"

type Hero struct {
	Name  string
	HP    int
	MaxHP int
}

func (h *Hero) TakeDamage(dmg int) {
	h.HP -= dmg
	if h.HP < 0 {
		h.HP = 0
	}
}
func (h *Hero) Heal(amount int) {
	h.HP += amount
	if h.HP > h.MaxHP {
		h.HP = h.MaxHP
	}
}
func (h Hero) IsAlive() bool  { return h.HP > 0 }
func (h Hero) Status() string { return fmt.Sprintf("%v: %d/%d", h.Name, h.HP, h.MaxHP) }

func main() {
	hero := Hero{Name: "Anton", HP: 100, MaxHP: 100}
	fmt.Println(hero.Status())
	hero.TakeDamage(70)
	fmt.Println(hero.Status())
	hero.Heal(30)
	fmt.Println(hero.Status())
	fmt.Println(hero.IsAlive())
	hero.TakeDamage(100)
	fmt.Println(hero.Status())
	fmt.Println(hero.IsAlive())
}
