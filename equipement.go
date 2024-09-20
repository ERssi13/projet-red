package projetred

import "fmt"

func equipmentMenu(player *User) {
	var choice int
	equipments := getEquipmentForClass(player.Classe)

	fmt.Printf("\nMenu de l'équipement pour %s :\n", player.Classe)
	for i, eq := range equipments {
		fmt.Printf("%d. %s (Coût: %d pièces d'or)\n", i+1, eq.Name, eq.Cost)
	}
	fmt.Println("0. Retour")

	fmt.Scan(&choice)

	if choice >= 1 && choice <= 3 {
		buyItem(player, equipments[choice-1])
	} else if choice == 0 {
		return 
	} else {
		fmt.Println("Choix invalide, essayez encore.")
		equipmentMenu(player)
	}
}



func (u *User) EquiperObjet(equipement string) {
	switch equipement {
	case "Chapeau de l’aventurier":
		if u.Equipements.Chapeau != "" {
			u.addInventaire(u.Equipements.Chapeau, 1)
		}
		u.Equipements.Chapeau = equipement
		u.PvMax += 10
		fmt.Println("Vous avez équipé le Chapeau de l’aventurier.")
	case "Tunique de l’aventurier":
		if u.Equipements.Tunique != "" {
			u.addInventaire(u.Equipements.Tunique, 1)
		}
		u.Equipements.Tunique = equipement
		u.PvMax += 25
		fmt.Println("Vous avez équipé la Tunique de l’aventurier.")
	case "Bottes de l’aventurier":
		if u.Equipements.Bottes != "" {
			u.addInventaire(u.Equipements.Bottes, 1)
		}
		u.Equipements.Bottes = equipement
		u.PvMax += 15
		fmt.Println("Vous avez équipé les Bottes de l’aventurier.")
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
	}
	u.removeInventory(equipement, 1)
}

func (u User) getNomEquipement(emplacement string) string {
	if equip, ok := u.Equipement[emplacement]; ok {
		return equip.Name
	}
	return "Aucun équipement" 
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
