package projetred

import "fmt"


func hasResources(player *User, required map[string]int) bool {
	for resource, quantity := range required {
		if player.ResourcesJ[resource] < quantity {
			fmt.Printf("Vous n'avez pas assez de %s.\n", resource)
			return false
		}
	}
	return true
}

func consumeResources(player *User, required map[string]int) {
	for resource, quantity := range required {
		player.ResourcesJ[resource] -= quantity
	}
}

func buyItem(player *User, item Equipment) {
	if player.MaxInventaire() {
		fmt.Println("Votre inventaire est plein, vous ne pouvez pas ajouter d'équipement supplémentaire.")
		return
	}

	if player.Or < item.Cost {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet objet.")
		return
	}

	if !hasResources(player, item.Resources) {
		return
	}

	player.Or -= item.Cost
	consumeResources(player, item.Resources)

	player.addInventaire(item.Name, 1)
	fmt.Printf("Vous avez fabriqué : %s ! Il vous reste %d pièces d'or.\n", item.Name, player.Or)

	fmt.Println("Vos ressources restantes :")
	for resource, quantity := range player.ResourcesJ {
		fmt.Printf("%s : %d\n", resource, quantity)
	}

	equipmentMenu(player)
}
func Forgeron(player *User) {
    fmt.Println("Bienvenue chez le Forgeron !")
    equipmentMenu(player) 
}