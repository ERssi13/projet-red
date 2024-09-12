package main

import (
	"fmt"
)

// Définition du personnage
type Character struct {
	Name       string
	Inventory  []string
	MaxHP      int
	CurrentHP  int
}

// Fonction d'initialisation (constructeur)
func NewCharacter(name string, maxHP int) *Character {
	return &Character{
		Name:       name,
		Inventory:  []string{},
		MaxHP:      maxHP,
		CurrentHP:  maxHP, // Le personnage commence avec des HP pleins
	}
}

// Fonction pour ajouter un item à l'inventaire
func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Item", item, "ajouté à l'inventaire.")
}

// Fonction pour afficher l'inventaire
func (c *Character) ShowInventory() {
	fmt.Println("Inventaire de", c.Name, ":", c.Inventory)
}

// Fonction pour utiliser une potion
func (c *Character) TakePotion() {
	for i, item := range c.Inventory {
		if item == "potion" {
			// Utiliser la potion
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...) // Supprimer la potion de l'inventaire

			// Régénérer les points de vie
			c.CurrentHP += 50
			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP // Ne pas dépasser les points de vie max
			}

			fmt.Println("Potion utilisée. Points de vie actuels :", c.CurrentHP, "/", c.MaxHP)
			return
		}
	}
	fmt.Println("Aucune potion dans l'inventaire.")
}
