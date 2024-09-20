package projetred

import "fmt"

func (u *User) TakePotion(potionType string) {
	if potionType == "Potion de Vie" {
		if u.PvActuel == u.PvMax {
			fmt.Println("Vos points de vie sont déjà au maximum. Vous ne pouvez pas utiliser une potion de vie.")
			return
		}

		if qty, found := u.Inventaire["Potion de Vie"]; found && qty > 0 {
			u.PvActuel += 50
			if u.PvActuel > u.PvMax {
				u.PvActuel = u.PvMax
			}
			u.removeInventory("Potion de Vie", 1)

			fmt.Printf("Potion de vie utilisée. Points de vie actuels : %d/%d\n", u.PvActuel, u.PvMax)
			u.Dead()
		} else {
			fmt.Println("Aucune potion de vie dans l'inventaire.")
		}
	} 
	
	if potionType == "Potion de Poison" {
		if qty, found := u.Inventaire["Potion de Poison"]; found && qty > 0 {
			u.PvActuel -= 30
			if u.PvActuel < 0 {
				u.PvActuel = 0
			}
			u.removeInventory("Potion de Poison", 1)

			fmt.Printf("Potion de poison utilisée. Points de vie actuels : %d/%d\n", u.PvActuel, u.PvMax)
			u.Dead()
		} else {
			fmt.Println("Aucune potion de poison dans l'inventaire.")
		}
	}
}
