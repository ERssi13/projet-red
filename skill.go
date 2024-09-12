package main

import (
	"fmt"
)

// Définition du personnage
type Character struct {
	Name      string
	Spellbook []string
}

// Fonction d'initialisation (constructeur)
func NewCharacter(name string) *Character {
	character := &Character{
		Name:      name,
		Spellbook: []string{"coup de poing"}, // Sort de base
	}
	return character
}

// Fonction pour ajouter un sort à la liste
func (c *Character) AddSpell(spell string) {
	// Vérifier si le sort est déjà dans la liste
	for _, s := range c.Spellbook {
		if s == spell {
			fmt.Println("Le sort", spell, "est déjà dans le livre de sorts.")
			return
		}
	}
	c.Spellbook = append(c.Spellbook, spell)
	fmt.Println("Sort", spell, "ajouté au livre de sorts.")
}

// Fonction pour ajouter le sort "boule de feu"
func (c *Character) AddFireball() {
	c.AddSpell("boule de feu")
}

// Fonction pour afficher les sorts dans le livre
func (c *Character) ShowSpellbook() {
	fmt.Println("Livre de sorts de", c.Name, ":", c.Spellbook)
}

func main() {
	// Création du personnage
	character := NewCharacter("Mage")

	// Afficher les sorts
	character.ShowSpellbook()

	// Ajouter "boule de feu" à la liste de sorts
	character.AddFireball()

	// Réessayer d'ajouter "boule de feu"
	character.AddFireball()

	// Afficher les sorts après ajout
	character.ShowSpellbook()
}
