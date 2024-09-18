package projetred

import "fmt"

// Définition de la structure Monstre
type Monstre struct {
    Nom             string
    PvMax           int
    PvActuels       int
    PointsAttaque   int
}

// Fonction pour initialiser un gobelin d'entraînement
func InitGoblin() Monstre {
    gobelin := Monstre{
        Nom:           "Gobelin d'entraînement",
        PvMax:         40,
        PvActuels:     40,  // Initialisé à PvMax
        PointsAttaque: 5,
    }
    return gobelin
}

// Méthode d'attaque pour un monstre
func (m *Monstre) Attaquer(cible *Monstre) {
    fmt.Printf("%s attaque %s pour %d points de dégâts !\n", m.Nom, cible.Nom, m.PointsAttaque)
    cible.PvActuels -= m.PointsAttaque
    if cible.PvActuels < 0 {
        cible.PvActuels = 0
    }
}

func charactere() {
    // Initialisation de deux monstres pour le test
    gobelin := InitGoblin()
    
    orc := Monstre{
        Nom:           "Orc d'entraînement",
        PvMax:         60,
        PvActuels:     60,
        PointsAttaque: 8,
    }
    
    // Exemple de tour d'attaque : le gobelin attaque l'orc, puis l'orc attaque le gobelin
    gobelin.Attaquer(&orc)
    fmt.Printf("PV restants de l'orc : %d/%d\n", orc.PvActuels, orc.PvMax)

    orc.Attaquer(&gobelin)
    fmt.Printf("PV restants du gobelin : %d/%d\n", gobelin.PvActuels, gobelin.PvMax)
}
