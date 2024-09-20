package projetred

import "fmt"

func displayStatus(player User, monster Monstre) {
	fmt.Printf("Joueur: %s, Santé: %d\n", player.Nom, player.PvActuel)
	fmt.Printf("Monstre: %s, Santé: %d\n", monster.Nom, monster.PVactuelm)
}

func playerTurn(player *User, monster *Monstre) {
	var choice int
	fmt.Println("Menu :")
	fmt.Println("1. Attaquer avec une compétence")
	fmt.Println("2. Utiliser une potion de soin")
	fmt.Print("Choisissez une option : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if len(player.Skills) > 0 {
			fmt.Println("Compétences disponibles :")
			for i, skill := range player.Skills {
				fmt.Printf("%d. %s (Dégâts: %d)\n", i+1, skill.Nom, skill.Degats)
			}
			var skillChoice int
			fmt.Print("Choisissez une compétence à utiliser : ")
			fmt.Scan(&skillChoice)
			if skillChoice > 0 && skillChoice <= len(player.Skills) {
				damage := player.Skills[skillChoice-1].Degats
				monster.PVactuelm -= damage
				fmt.Printf("%s utilise %s et inflige %d dégâts à %s!\n", player.Nom, player.Skills[skillChoice-1].Nom, damage, monster.Nom)
			} else {
				fmt.Println("Choix invalide.")
			}
		} else {
			fmt.Println("Vous n'avez aucune compétence à utiliser.")
		}
	case 2:
		player.TakePotion("Potion de Vie")
	default:
		fmt.Println("Choix invalide. Vous perdez votre tour.")
	}
}

func monsterTurn(player *User, monster *Monstre) {
	fmt.Println("Tour du monstre:")
	player.PvActuel -= monster.Degat
	fmt.Printf("%s attaque %s pour %d dégâts!\n", monster.Nom, player.Nom, monster.Degat)
}

func TrainingFight(player User, monster Monstre) {
	for player.PvActuel > 0 && monster.PVactuelm > 0 {
		playerTurn(&player, &monster)
		displayStatus(player, monster)
		if monster.PVactuelm <= 0 {
			fmt.Printf("%s a vaincu %s!\n", player.Nom, monster.Nom)
			return
		}
		monsterTurn(&player, &monster)
		displayStatus(player, monster)
		if player.PvActuel <= 0 {
			fmt.Printf("%s est vaincu!\n", player.Nom)
			player.Dead() 
		}
	}
}
