package main

import (
	"fmt"
)

type Player struct {
	Gold         int
	Inventory    []string
	Class        string
	Resources    map[string]int
	MaxInventory int
}

type Equipment struct {
	Name      string
	Cost      int
	Resources map[string]int
}

func main() {
	player := Player{
		Gold:         100,
		MaxInventory: 10,
		Resources: map[string]int{
			"Plume de Corbeau":   1,
			"Cuir de Sanglier":   2,
			"Fourrure de Loup":   3,
			"Peau de Troll":      1,
			"Pierre Magique":     2,
			"Essence de Fantôme": 1,
			"Acier Trempé":       1,
			"Bois de Chêne":      2,
		},
	}

	chooseClass(&player)
}

// Choisir la classe du joueur
func chooseClass(player *Player) {
	var choice int
	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Assassin")
	fmt.Println("3. Voleur")
	fmt.Println("4. Archer")
	fmt.Println("5. Mage")
	fmt.Println("6. Quitter")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		player.Class = "Guerrier"
		equipmentMenu(player)
	case 2:
		player.Class = "Assassin"
		equipmentMenu(player)
	case 3:
		player.Class = "Voleur"
		equipmentMenu(player)
	case 4:
		player.Class = "Archer"
		equipmentMenu(player)
	case 5:
		player.Class = "Mage"
		equipmentMenu(player)
	case 6:
		fmt.Println("Quitter le jeu...")
		return
	default:
		fmt.Println("Choix invalide, essayez encore.")
		chooseClass(player)
	}
}

func equipmentMenu(player *Player) {
	var choice int
	equipments := getEquipmentForClass(player.Class)

	fmt.Printf("\nMenu de l'équipement pour %s :\n", player.Class)
	for i, eq := range equipments {
		fmt.Printf("%d. %s (Coût: %d pièces d'or)\n", i+1, eq.Name, eq.Cost)
	}
	fmt.Println("4. Retour")

	fmt.Scan(&choice)

	if choice >= 1 && choice <= 3 {
		buyItem(player, equipments[choice-1])
	} else if choice == 4 {
		chooseClass(player)
	} else {
		fmt.Println("Choix invalide, essayez encore.")
		equipmentMenu(player)
	}
}

func getEquipmentForClass(class string) []Equipment {
	switch class {
	case "Guerrier":
		return []Equipment{
			{Name: "Chapeau de l'aventurier", Cost: 50, Resources: map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}},
			{Name: "Tunique de l'aventurier", Cost: 50, Resources: map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}},
			{Name: "Bottes de l'aventurier", Cost: 50, Resources: map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}},
		}
	case "Assassin":
		return []Equipment{
			{Name: "Dague de l'ombre", Cost: 50, Resources: map[string]int{"Cuir de sanglier": 1, "Acier Trempé": 1}},
			{Name: "Cape de l'ombre", Cost: 50, Resources: map[string]int{"Fourrure de Loup": 1, "Plume de Corbeau": 1}},
			{Name: "Bottes furtives", Cost: 50, Resources: map[string]int{"Peau de Troll": 1, "Cuir de Sanglier": 1}},
		}
	case "Voleur":
		return []Equipment{
			{Name: "Gants de pickpocket", Cost: 50, Resources: map[string]int{"Cuir de Sanglier": 1, "Plume de Corbeau": 1}},
			{Name: "Cagoule du brigand", Cost: 50, Resources: map[string]int{"Fourrure de Loup": 1, "Peau de Troll": 1}},
			{Name: "Chaussures silencieuses", Cost: 50, Resources: map[string]int{"Cuir de Sanglier": 1, "Fourrure de Loup": 1}},
		}
	case "Archer":
		return []Equipment{
			{Name: "Arc du chasseur", Cost: 50, Resources: map[string]int{"Bois de Chêne": 2, "Acier Trempé": 1}},
			{Name: "Carquois renforcé", Cost: 50, Resources: map[string]int{"Cuir de Sanglier": 2, "Fourrure de Loup": 1}},
			{Name: "Gants de tir", Cost: 50, Resources: map[string]int{"Cuir de Sanglier": 1, "Plume de Corbeau": 1}},
		}
	case "Mage":
		return []Equipment{
			{Name: "Bâton du sorcier", Cost: 50, Resources: map[string]int{"Pierre Magique": 2, "Bois de Chêne": 1}},
			{Name: "Robe de l'enchanteur", Cost: 50, Resources: map[string]int{"Essence de Fantôme": 1, "Peau de Troll": 1}},
			{Name: "Anneau magique", Cost: 50, Resources: map[string]int{"Pierre Magique": 1, "Essence de Fantôme": 1}},
		}
	default:
		return nil
	}
}

func hasResources(player *Player, required map[string]int) bool {
	for resource, quantity := range required {
		if player.Resources[resource] < quantity {
			fmt.Printf("Vous n'avez pas assez de %s.\n", resource)
			return false
		}
	}
	return true
}

func consumeResources(player *Player, required map[string]int) {
	for resource, quantity := range required {
		player.Resources[resource] -= quantity
	}
}

func buyItem(player *Player, item Equipment) {
	if len(player.Inventory) >= player.MaxInventory {
		fmt.Println("Votre inventaire est plein, vous ne pouvez pas ajouter d'équipement supplémentaire.")
		return
	}

	if player.Gold < item.Cost {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet objet.")
		return
	}

	if !hasResources(player, item.Resources) {
		return
	}

	player.Gold -= item.Cost
	consumeResources(player, item.Resources)

	player.Inventory = append(player.Inventory, item.Name)
	fmt.Printf("Vous avez fabriqué : %s ! Il vous reste %d pièces d'or.\n", item.Name, player.Gold)

	fmt.Println("Vos ressources restantes :")
	for resource, quantity := range player.Resources {
		fmt.Printf("%s : %d\n", resource, quantity)
	}
	equipmentMenu(player)
}
