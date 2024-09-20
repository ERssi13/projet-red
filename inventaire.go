package projetred

import (
	"fmt"
	"strings"
)

func (u *User) AccesInventaire() {
	const largeurTotale = 70

	for {
		clearTerminal()
		fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
		fmt.Println("║" + centerText("Inventaire", largeurTotale-2) + "║")
		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

		if len(u.Inventaire) == 0 {
			fmt.Println("║ " + centerText("L'inventaire est vide.", largeurTotale-2) + "║")
		} else {
			for item, quantity := range u.Inventaire {
				AfficherLigne(item, fmt.Sprintf("%d", quantity), largeurTotale)
			}

			fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")
			fmt.Println("║ 1. Utiliser une potion                                             ║")
			fmt.Println("║ 2. Utiliser Livre de Sort : Boule de Feu                           ║")
			fmt.Println("║ 3. Équiper un objet                                                ║") 
			fmt.Println("║ 0. Retourner au menu                                               ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

			var choix int
			fmt.Print("Choisissez une option: ")
			fmt.Scan(&choix)

			switch choix {
			case 1:
				u.ChoisirPotion()
			case 2:
				if _, found := u.Inventaire["Livre de Sort : Boule de Feu"]; found {
					u.spellBook()
					u.removeInventory("Livre de Sort : Boule de Feu", 1) 
				} else {
					fmt.Println("Vous ne possédez pas de 'Livre de Sort : Boule de Feu'.")
				}
			case 3:
				u.EquiperObjetDepuisInventaire()
			case 0:
				return
			default:
				fmt.Println("Option invalide.")
			}
		}
	}
}

func (u *User) AugmenterInventaire() {
    if u.Maxinventaire >= 40 {
        fmt.Println("Votre inventaire est déjà à la taille maximale.")
        return
    }
    cout := (u.Maxinventaire / 10) * 20
    if u.Or < cout {
        fmt.Printf("Vous n'avez pas assez d'or pour augmenter la taille de l'inventaire. Il vous faut %d or.\n", cout)
        return
    }
    u.Or -= cout
    u.Maxinventaire += 10
    fmt.Printf("Votre inventaire a été augmenté à %d emplacements.\n", u.Maxinventaire)
}

func (u *User) EquiperObjetDepuisInventaire() {
	fmt.Println("Objets disponibles à équiper:")

	var equipementDispo []string
	for item, quantity := range u.Inventaire {
		if item == "Chapeau de l’aventurier" || item == "Tunique de l’aventurier" || item == "Bottes de l’aventurier" {
			equipementDispo = append(equipementDispo, item)
			fmt.Printf("%d. %s (Quantité: %d)\n", len(equipementDispo), item, quantity)
		}
	}

	if len(equipementDispo) == 0 {
		fmt.Println("Aucun équipement disponible dans l'inventaire.")
		return
	}
	var choix int
	fmt.Print("Entrez le numéro de l'objet à équiper: ")
	fmt.Scan(&choix)

	if choix >= 1 && choix <= len(equipementDispo) {
		u.EquiperObjet(equipementDispo[choix-1]) 
	} else {
		fmt.Println("Choix invalide.")
	}
}

func (u *User) ChoisirPotion() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║ 1. Potion de Vie                                                   ║")
	fmt.Println("║ 2. Potion de Poison                                                ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	var choixPotion int
	fmt.Print("Choisissez une potion à utiliser: ")
	fmt.Scan(&choixPotion)

	switch choixPotion {
	case 1:
		u.TakePotion("Potion de Vie")
	case 2:
		u.TakePotion("Potion de Poison")
	default:
		fmt.Println("Potion invalide.")
	}
}

func (u *User) MaxInventaire() bool {
	return len(u.Inventaire) >= u.Maxinventaire
}

func (u *User) addInventaire(item string, quantity int) bool {
	if u.Inventaire == nil {
		u.Inventaire = make(map[string]int)
	}

	totalObjets := 0
	for _, q := range u.Inventaire {
		totalObjets += q
	}

	espaceRestant := u.Maxinventaire - totalObjets
	if quantity > espaceRestant {
		fmt.Printf("Vous n'avez pas assez d'espace dans l'inventaire. Il reste seulement %d emplacements.\n", espaceRestant)
		return false
	}

	u.Inventaire[item] += quantity
	fmt.Printf("Vous avez ajouté %d %s à votre inventaire.\n", quantity, item)
	return true
}

func (u *User) removeInventory(item string, qty int) {
	if currentQty, found := u.Inventaire[item]; found {
		if currentQty > qty {
			u.Inventaire[item] -= qty
		} else {
			delete(u.Inventaire, item)
		}
	}
}

