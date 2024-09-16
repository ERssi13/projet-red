package projetred

import (
	"fmt"
	"strings"
)

type User struct {
	Nom                   string
	Classe                string
	Niveau                int
	Or                    int
	PvMax                 int
	PvActuel              int
	Inventaire            map[string]int
	hasReceivedFreePotion bool
}

func Init(nom string, classe string, niveau int, or int, pvMax int, pvActuel int, inventaire map[string]int, hasreceivedFreePotion bool) User {
	return User{
		Nom:                   nom,
		Classe:                classe,
		Niveau:                niveau,
		Or:                    or,
		PvMax:                 pvMax,
		PvActuel:              pvActuel,
		Inventaire:            inventaire,
		hasReceivedFreePotion: hasreceivedFreePotion,
	}
}

func (u User) AfficherInfos() {
	const largeurTotale = 70

	for {
		clearTerminal()
		fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
		fmt.Println("║" + centerText("Informations du Personnage", largeurTotale-2) + "║")
		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

		AfficherLigne("Nom", u.Nom, largeurTotale)
		AfficherLigne("Classe", u.Classe, largeurTotale)
		AfficherLigne("Niveau", fmt.Sprintf("%d", u.Niveau), largeurTotale)
		AfficherLigne("Or", fmt.Sprintf("%d", u.Or), largeurTotale)
		AfficherLigne("Points de vie max", fmt.Sprintf("%d", u.PvMax), largeurTotale)
		AfficherLigne("Points de vie actuels", fmt.Sprintf("%d", u.PvActuel), largeurTotale)

		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")
		fmt.Println("║ 1. Retourner au menu                                               ║")
		fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")

		var choix int
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

		if choix == 1 {
			return
		} else {
			fmt.Println("Option invalide.")
		}
	}
}

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
			fmt.Println("║ 2. Retourner au menu                                               ║")
			fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

			var choix int
			fmt.Print("Choisissez une option: ")
			fmt.Scan(&choix)

			switch choix {
			case 1:
				u.TakePotion()
			case 2:
				return
			default:
				fmt.Println("Option invalide.")
			}
		}
	}
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

func (u *User) TakePotion() {
	if qty, found := u.Inventaire["Potion"]; found && qty > 0 {
		u.PvActuel += 50
		if u.PvActuel > u.PvMax {
			u.PvActuel = u.PvMax

			u.removeInventory("Potion", 1)

			fmt.Printf("Potion utilisée. Points de vie actuels : %d/%d\n", u.PvActuel, u.PvMax)
			u.Dead()
		} else {
			fmt.Println("Aucune potion dans l'inventaire.")
		}
	}
}
