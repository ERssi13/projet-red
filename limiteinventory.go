package main

import (
	"fmt"
)

// Définition du personnage
type Character struct {
	Name      string
	Inventory []string
	MaxItems  int
}

// Fonction d'initialisation (constructeur)
func NewCharacter(name string) *Character {
	return &Character{
		Name:      name,
		Inventory: []string{},
		MaxItems:  10, // Limite de 10 items
	}
}

// Fonction pour vérifier si l'inventaire est plein
func (c *Character) IsInventoryFull() bool {
	return len(c.Inventory) >= c.MaxItems
}

// Fonction pour ajouter un item à l'inventaire
func (c *Character) AddItem(item string) {
	// Vérifier si l'inventaire est plein
	if c.IsInventoryFull() {
		fmt.Println("Impossible d'ajouter", item, ": inventaire plein.")
		return
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Item", item, "ajouté à l'inventaire.")
}
