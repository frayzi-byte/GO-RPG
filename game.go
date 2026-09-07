package main
import (
    "fmt"
    "math/rand/v2"
)

type player struct {
	name string
	hp int 
	isAlive bool
}

type enemy struct {
	name string
	hp int
	isAlive bool
}

func NewPlayer(name string, hp int, isAlive bool) *player {
	return &player{name : name, hp : hp, isAlive : isAlive}	
}

func NewEnemy(name string, hp int, isAlive bool) *enemy {
	return &enemy{name : name, hp : hp, isAlive : isAlive}
}

func (p *player) TakeDamageEnemy(e *enemy) {
	damage := 0
	fmt.Println("(Max : 50) Enter the damage value:")
	fmt.Scanln(&damage)
	if damage > 50 {
		fmt.Println("Damage value cannot be more than 50")
		return
	} else {
		e.hp -= damage
		fmt.Println(p.name, "attacked", e.name, "for", damage, "damage")
		if e.hp <= 0 {
			e.isAlive = false
			fmt.Println(e.name, "is dead")
		}
	}
}

func (e *enemy) TakeDamagePlayer(p *player) {
	damage := rand.IntN(51)
	number := rand.IntN(11)
	if number < 5 {
		fmt.Println(e.name, "missed the attack")
	} else {
		p.hp -= damage
		fmt.Println(e.name, "attacked", p.name, "for", damage, "damage")
		if p.hp <= 0 {
			p.isAlive = false
			fmt.Println(p.name, "is dead")
		}
	}
}

func createPlayer() *player {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	player := NewPlayer(name, 100, true)
	return player
}

func main() {
	fmt.Println("Welcome to the game!")
	player := createPlayer()
	enemy := NewEnemy("Goblin", 100, true)

	for player.isAlive && enemy.isAlive {
		fmt.Printf("\n[Статус] %s: %d HP | %s: %d HP\n", player.name, player.hp, enemy.name, enemy.hp)

		fmt.Println("\nТвой ход!")
		player.TakeDamageEnemy(enemy)

		if !enemy.isAlive {
			fmt.Println("\n🎉 You won! You defeated the enemy!")
			break
		}

		fmt.Println("\nХод врага...")
		enemy.TakeDamagePlayer(player)

		if !player.isAlive {
			fmt.Println("\n💀 You died! Game over.")
			break
		}
	}

	fmt.Println("\nThank you for playing!")
}
