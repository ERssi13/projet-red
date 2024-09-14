package main

import (
	"fmt"
)

// Structure représentant un personnage
type Character struct {
	Name          string
	MaxHealth     int
	CurrentHealth int
	Attack        int
}

// InitGoblin initialise un Gobelin d'entraînement avec des points de vie et une attaque spécifiques
func InitGoblin() Character {
	return Character{
		Name:          "Gobelin d'entrainement",
		MaxHealth:     100,
		CurrentHealth: 100,
		Attack:        5,
	}
}

// goblinPattern simule le combat entre le Gobelin et un joueur
func goblinPattern(goblin Character, player Character, turns int) {
	for turn := 1; turn <= turns; turn++ {
		// Calcul des dégâts infligés par le Gobelin
		damage := goblin.Attack
		if turn%3 == 0 {
			// Tous les 3 tours, le Gobelin inflige 200% de son attaque
			damage = goblin.Attack * 2
		}

		// Réduire les points de vie du joueur en fonction des dégâts infligés
		player.CurrentHealth -= damage
		if player.CurrentHealth < 0 {
			player.CurrentHealth = 0 // Empêche les points de vie négatifs
		}

		// Affichage des informations de combat
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, damage)
		fmt.Printf("Points de vie de %s : %d/%d\n", player.Name, player.CurrentHealth, player.MaxHealth)

		// Si le joueur est à 0 PV, le combat s'arrête
		if player.CurrentHealth == 0 {
			fmt.Printf("%s est vaincu!\n", player.Name)
			break
		}
	}
}

func main() {
	// Initialiser le Gobelin et le joueur
	goblin := InitGoblin()
	player := Character{
		Name:          "Personnage",
		MaxHealth:     50,
		CurrentHealth: 50,
		Attack:        10,
	}

	// Simuler un combat de 10 tours
	goblinPattern(goblin, player, 10)
}
