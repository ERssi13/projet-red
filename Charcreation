package main

import (
	"fmt"
)

// Character structure
type Character struct {
	Class          string
	MaxHealth      int
	CurrentHealth  int
	Level          int
}

// Init function to create the character
func Init(class string) Character {
	var maxHealth int

	// Définir la santé maximale en fonction de la classe
	switch class {
	case "guerrier":
		maxHealth = 120
	case "archer":
		maxHealth = 100
	case "assassin":
		maxHealth = 100
	case "mage":
		maxHealth = 80
	case "voleur":
		maxHealth = 100
	default:
		fmt.Println("Classe non reconnue.")
		return Character{}
	}

	//Créer le personnage avec 50% de la santé maximale et le niveau 1
	character := Character{
		Class:         class,
		MaxHealth:     maxHealth,
		CurrentHealth: maxHealth / 2, // 50% of max health
		Level:         1,             // Starting level
	}
	return character
}

// CharCreation function to prompt user to create a character
func CharCreation() Character {
	var class string

	// Get user input for class
	fmt.Println("Choisissez une classe (guerrier, archer, assassin, mage, voleur) :")
	fmt.Scanln(&class)

	// Initialize the character
	character := Init(class)
	
	// Display the character details
	if character.Class != "" {
		fmt.Printf("Vous avez créé un %s avec %d/%d points de vie et de niveau %d.\n", 
			character.Class, character.CurrentHealth, character.MaxHealth, character.Level)
	}

	return character
}

func main() {
	// Call CharCreation to create a character
	CharCreation()
}
