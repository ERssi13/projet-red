package projetred 

import (
	"fmt"
)

// Définir une structure pour les personnages (joueur et monstre)
type Character struct {
	Name  string
	Health int
	Attack int
}

// Fonction pour afficher l'état actuel des personnages
func displayStatus(player, monster Character) {
	fmt.Printf("Joueur: %s, Santé: %d\n", player.Name, player.Health)
	fmt.Printf("Monstre: %s, Santé: %d\n", monster.Name, monster.Health)
}

// Fonction pour gérer le tour du joueur
func playerTurn(player, monster *Character) {
	fmt.Println("Tour du joueur:")
	monster.Health -= player.Attack
	fmt.Printf("%s attaque %s pour %d dégâts!\n", player.Name, monster.Name, player.Attack)
}

// Fonction pour gérer le tour du monstre
func monsterTurn(player, monster *Character) {
	fmt.Println("Tour du monstre:")
	player.Health -= monster.Attack
	fmt.Printf("%s attaque %s pour %d dégâts!\n", monster.Name, player.Name, monster.Attack)
}

var monster Character {
	Name 			int
	Healt			int
	Mana 			int 
	PvActuel 		int
	Niveau			int
	Attackconst		int
	Force			int
}

func projetred() {
	player := Character{Name: "Héros", Health: 100, Attack: 20}
	monster := Character{Name: "Dragon", Health: 80, Attack: 15}

	trainingFight(player, monster)
}

