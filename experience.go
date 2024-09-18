package projetred

import (
	"fmt"
)

// Monstre représente un monstre avec un certain nombre de points d'expérience.
type Monstre struct {
	Nom string
	Exp int
}

// Fonction pour créer un nouveau joueur
func NouveauJoueur(nom string) User {
	return User{
		Nom:          nom,
		Niveau:       1,
		ExpActuelle:  0,
		ExpMax:       100, // Expérience requise pour le niveau 1
		Force:        10,
		Agilite:      10,
		Intelligence: 10,
	}
}

// Fonction pour que le joueur gagne de l'expérience après un combat
func (j *User) GagnerExperience(exp int) {
	fmt.Printf("%s gagne %d points d'expérience.\n", j.Nom, exp)
	j.ExpActuelle += exp

	// Vérifier si le joueur atteint le niveau suivant
	for j.ExpActuelle >= j.ExpMax {
		excesExp := j.ExpActuelle - j.ExpMax
		j.MonterDeNiveau()
		j.ExpActuelle = excesExp
	}
}

// Fonction pour gérer la montée en niveau
func (j *User) MonterDeNiveau() {
	j.Niveau++
	j.ExpMax += j.Niveau * 50 // Augmentation de l'Exp max nécessaire à chaque niveau
	fmt.Printf("%s monte au niveau %d !\n", j.Nom, j.Niveau)

	// Attribuer des bonus de statistiques au joueur
	j.Force += 2
	j.Agilite += 2
	j.Intelligence += 2
	fmt.Printf("%s gagne des bonus de statistiques ! Force: %d, Agilité: %d, Intelligence: %d\n",
		j.Nom, j.Force, j.Agilite, j.Intelligence)
}

// Simuler un combat entre le joueur et un monstre
func Combat(joueur *User, monstre Monstre) {
	fmt.Printf("%s combat %s et le bat !\n", joueur.Nom, monstre.Nom)
	joueur.GagnerExperience(monstre.Exp)
}
