package projetred

import "fmt"

func (u *User) addInventaire(item string, quantity int) {
	if u.Inventaire == nil {
		u.Inventaire = make(map[string]int)
	}
	u.Inventaire[item] += quantity
	fmt.Printf("Vous avez ajouté %d %s à votre inventaire.\n", quantity, item)
}

func (u *User) spendOr(amount int) bool {
	if u.Or >= amount {
		u.Or -= amount
		fmt.Printf("Vous avez dépensé %d or. Il vous reste %d or.\n", amount, u.Or)
		return true
	} else {
		fmt.Println("Vous n'avez pas assez d'or pour acheter cet article.")
		return false
	}
}

func (u *User) Marchand() {
	for {
		clearTerminal()
		fmt.Printf("Or: %d\n", u.Or)
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                         Interface du Marchand                    ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Acheter une potion de vie (gratuit la première fois)          ║")
		fmt.Println("║ 2. Acheter une potion de poison (10 or)                          ║")
		fmt.Println("║ 3. Acheter une épée en bois (15 or)                              ║")
		fmt.Println("║ 4. Acheter une épée en bronze (30 or)                            ║")
		fmt.Println("║ 5. Acheter une épée en fer (50 or)                               ║")
		fmt.Println("║ 6. Retourner au menu                                             ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════╝")

		var choix int
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if !u.hasReceivedFreePotion {
				u.addInventaire("Potion de Vie", 1)
				fmt.Println("Vous avez reçu une potion de vie gratuitement.")
				u.hasReceivedFreePotion = true
			} else {
				if u.spendOr(5) {
					u.addInventaire("Potion de Vie", 1)
				}
			}
		case 2:
			if u.spendOr(10) {
				u.addInventaire("Potion de Poison", 1)
			}
		case 3:
			if u.spendOr(15) {
				u.addInventaire("Épée en bois", 1)
			}
		case 4:
			if u.spendOr(30) {
				u.addInventaire("Épée en bronze", 1)
			}
		case 5:
			if u.spendOr(50) {
				u.addInventaire("Épée en fer", 1)
			}
		case 6:
			return
		default:
			fmt.Println("Option invalide.")
		}
	}
}
