package main

import (
	"fmt"
	"strings"
)

// Structure représentant un personnage
type Character struct {
	Name			string
	MaxHealth		int
	CurrentHealth	int
	Attack			int
	Inventory		[]string
}

// InitGoblin initialise un Gobelin d'entraînement avec des points de vie et une attaque spécifiques
func InitGoblin() Character {
	return Character{
		Name:	"Gobelin d'entrainement",
		MaxHealth:	40,
		CurrentHealth:	40,
		Attack:	5,
		Inventory:	nil, // Le Gobelin n'a pas d'inventaire
	}
}

// Menu d'attaque ou d'inventaire pour le joueur
func charTurn(player *Character, opponent *Character) {
	var choice string
	fmt.Println("C'est votre tour de jouer !")
	fmt.Println("Menu :")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Choisissez une option : ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		// Attaque basique
		fmt.Printf("%s utilise Attaque basique et inflige 5 dégâts à %s\n", player.Name, opponent.Name)
		opponent.CurrentHealth -= 5
		if opponent.CurrentHealth < 0 {
			opponent.CurrentHealth = 0
		}
		fmt.Printf("PV de %s : %d/%d\n", opponent.Name, opponent.CurrentHealth, opponent.MaxHealth)
	case "2":
		// Utilisation de l'inventaire
		if len(player.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide !")
		} else {
			fmt.Println("Inventaire :")
			for i, item := range player.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			var itemChoice int
			fmt.Print("Choisissez un objet à utiliser : ")
			fmt.Scanln(&itemChoice)

			if itemChoice > 0 && itemChoice <= len(player.Inventory) {
				item := player.Inventory[itemChoice-1]
				useItem(player, item)
				player.Inventory = append(player.Inventory[:itemChoice-1], player.Inventory[itemChoice:]...)
				} else {
					fmt.Println("Choix invalide.")
			}
		}
	default:
		fmt.Println("Choix invalide. Vous perdez votre tour.")
	}

	// Si l'adversaire est encore en vie, il attaque à son tour
	if opponent.CurrentHealth > 0 {
		opponentTurn(opponent, player)
	}
}

// Fonction qui applique l'effet des objets de l'inventaire
func useItem(player *Character, item string) {
	switch strings.ToLower(item) {
	case "potion de soin":
		healAmount := 10
		player.CurrentHealth += healAmount
		if player.CurrentHealth > player.MaxHealth {
			player.CurrentHealth = player.MaxHealth
		}
		fmt.Printf("Vous utilisez %s et regagnez %d PV. Vos PV : %d/%d\n", item, healAmount, player.CurrentHealth, player.MaxHealth)
	default:
		fmt.Printf("L'objet %s n'a aucun effet.\n", item)
	}
}

// Tour du monstre
func opponentTurn(opponent *Character, player *Character) {
	fmt.Printf("%s attaque %s et inflige %d dégâts\n", opponent.Name, player.Name, opponent.Attack)
	player.CurrentHealth -= opponent.Attack
	if player.CurrentHealth < 0 {
		player.CurrentHealth = 0
	}
	fmt.Printf("PV de %s : %d/%d\n", player.Name, player.CurrentHealth, player.MaxHealth)
}

func main() {
	// Initialiser le joueur et le Gobelin
	player := Character{
		Name:	"Personnage",
		MaxHealth:	50,
		CurrentHealth:	50,
		Attack:		5,
		Inventory:     []string{"Potion de soin"},
	}
	goblin := InitGoblin()

	// Boucle de combat jusqu'à ce qu'un des deux soit vaincu
	for player.CurrentHealth > 0 && goblin.CurrentHealth > 0 {
		charTurn(&player, &goblin)
	}

	// Résultat final
	if player.CurrentHealth > 0 {
		fmt.Println("Vous avez vaincu le Gobelin d'entrainement !")
	} else {
		fmt.Println("Vous avez été vaincu par le Gobelin d'entrainement.")
	}
}
