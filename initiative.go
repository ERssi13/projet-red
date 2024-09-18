package projetred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Définir une structure pour les personnages (joueur et gobelin) avec l'initiative
type Character struct {
	Name      string
	Health    int
	Attack    int
	Initiative int
}

// Définir une structure pour les objets
type Item struct {
	Name  string
	Effect func(*Character) // Fonction pour appliquer l'effet de l'objet
}

// Fonction pour afficher l'état actuel des personnages
func displayStatus(attacker, attacked Character) {
	fmt.Printf("%s inflige %d dégâts à %s\n", attacker.Name, attacker.Attack, attacked.Name)
	fmt.Printf("%s : %d/%d PV\n", attacked.Name, attacked.Health, 100) // On suppose ici que le max est 100
}

// Fonction pour gérer le tour du joueur
func charTurn(player *Character, monster *Character, items []Item) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Menu:")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")

	fmt.Print("Choisissez une option (1/2): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		// Attaquer
		damage := 5
		monster.Health -= damage
		fmt.Printf("%s utilise Attaque basique\n", player.Name)
		fmt.Printf("%s inflige %d dégâts à %s\n", player.Name, damage, monster.Name)
		fmt.Printf("%s : %d/%d PV\n", monster.Name, monster.Health, 80) // On suppose ici que le max est 80
		if monster.Health <= 0 {
			fmt.Printf("%s est vaincu! Le joueur gagne!\n", monster.Name)
			return
		}
	case "2":
		// Inventaire
		fmt.Println("Inventaire:")
		for i, item := range items {
			fmt.Printf("%d. %s\n", i+1, item.Name)
		}
		fmt.Print("Choisissez un objet à utiliser (ou appuyez sur Enter pour revenir): ")
		itemInput, _ := reader.ReadString('\n')
		itemInput = strings.TrimSpace(itemInput)

		if itemInput != "" {
			index := 0
			fmt.Sscanf(itemInput, "%d", &index)
			if index > 0 && index <= len(items) {
				selectedItem := items[index-1]
				fmt.Printf("Vous utilisez %s\n", selectedItem.Name)
				selectedItem.Effect(player)
			} else {
				fmt.Println("Choix invalide!")
			}
		}
		return
	default:
		fmt.Println("Option invalide, veuillez réessayer.")
	}
}

// Fonction pour gérer le combat avec un Gobelin d'entraînement
func goblinPattern(player *Character, goblin *Character) {
	turn := 1
	for player.Health > 0 && goblin.Health > 0 {
		fmt.Printf("Tour %d:\n", turn)
		damageMultiplier := 1
		if turn%3 == 0 {
			damageMultiplier = 2
		}
		damage := goblin.Attack * damageMultiplier

		player.Health -= damage
		displayStatus(*goblin, *player)

		if player.Health <= 0 {
			fmt.Printf("%s est vaincu! Le Gobelin d'entraînement gagne!\n", player.Name)
			return
		}

		turn++
	}
}

// Fonction pour gérer le combat d'entraînement avec initiative
func trainingFight(player *Character, goblin *Character, items []Item) {
	// Déterminer qui commence le tour
	if player.Initiative >= goblin.Initiative {
		fmt.Println("Le joueur commence!")
		for player.Health > 0 && goblin.Health > 0 {
			charTurn(player, goblin, items)
			if goblin.Health <= 0 {
				break
			}
			fmt.Println("Au tour du Gobelin.")
			goblinPattern(player, goblin)
			if player.Health <= 0 {
				break
			}
		}
	} else {
		fmt.Println("Le Gobelin commence!")
		for player.Health > 0 && goblin.Health > 0 {
			goblinPattern(player, goblin)
			if player.Health <= 0 {
				break
			}
			fmt.Println("Au tour du joueur.")
			charTurn(player, goblin, items)
			if goblin.Health <= 0 {
				break
			}
		}
	}
}

// Fonction principale du menu de départ
func projetred() {
	player := Character{Name: "Héros", Health: 100, Attack: 20, Initiative: 10}
	goblin := Character{Name: "Gobelin d'entraînement", Health: 80, Attack: 10, Initiative: 5}

	// Exemples d'objets avec effets
	items := []Item{
		{Name: "Potion de soin", Effect: func(c *Character) { c.Health += 10; if c.Health > 100 { c.Health = 100 } }},
		{Name: "Élixir de force", Effect: func(c *Character) { c.Attack += 5 }},
	}

	for {
		fmt.Println("Menu principal:")
		fmt.Println("1. Entraînement")
		fmt.Println("2. Quitter")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choisissez une option (1/2): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			trainingFight(&player, &goblin, items)
		case "2":
			fmt.Println("Quitter le jeu.")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
