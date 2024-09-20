package projetred

import"fmt"

type Monstre struct {
	Nom 		string
	PVmaxm 		int
	PVactuelm 	int
	Degat 		int 
}

func InitGoblin() Monstre {
	return Monstre{
		Nom:          "Gobelin d'entrainement",
		PVmaxm:     40,
		PVactuelm: 40,
		Degat:        5,
	}
}

func goblinPattern(goblin Monstre, player User, turns int) {
	for turn := 1; turn <= turns; turn++ {
		damage := goblin.Degat
		if turn%3 == 0 {
			damage = goblin.Degat * 2
		}
		player.PvActuel -= damage
		if player.PvActuel < 0 {
			player.PvActuel = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Nom, player.Nom, damage)
		fmt.Printf("Points de vie de %s : %d/%d\n", player.Nom, player.PvActuel, player.PvMax)
		if player.PvActuel == 0 {
			fmt.Printf("%s est vaincu!\n", player.Nom)
			break
		}
	}
}