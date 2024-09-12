package main

import "fmt"

// Définition de la structure du joueur
type Joueur struct {
	nom         string
	vieActuelle int
	vieMax      int
}

// Fonction qui vérifie si le joueur est mort, et le ressuscite avec 50% de sa vie max
func (j *Joueur) Dead() {
	if j.vieActuelle <= 0 {
		fmt.Printf("%s est mort. Il est ressuscité avec 50%% de sa vie maximale.\n", j.nom)
		j.vieActuelle = j.vieMax / 2
	} else {
		fmt.Printf("%s est encore vivant avec %d points de vie.\n", j.nom, j.vieActuelle)
	}
}

func main() {
	joueur := Joueur{nom: "Hero", vieActuelle: 0, vieMax: 100}
    
    // Vérification si le joueur est mort et ressuscitation
	joueur.Dead()
    
    // Affichage des points de vie après la vérification
	fmt.Printf("%s a maintenant %d points de vie.\n", joueur.nom, joueur.vieActuelle)
}
