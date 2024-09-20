//il faut creer un autre dossier pour y mettre le main 
package main

import (
	"bufio"
	"fmt"
	"os"
	"projetred"
	"strings"
)

func AppelMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez le nom de votre Personnage: ")
	scanner.Scan()
	nom := strings.TrimSpace(scanner.Text())
	classe := projetred.ChoisirClasse()

	inventaire := map[string]int{
		"Potion de Vie": 3,
		"Iron sword":    1,
	}

	p1 := projetred.Init(nom, classe, 1, 1000, 100, 100, inventaire, nil, false, 10, 10)

	var choix int
	for {
		projetred.AfficherMenu()
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			duel(&p1)
		case 2:
			p1.AfficherInfos()
		case 3:
			p1.AccesInventaire()
		case 4:
			p1.Marchand()
		case 5:
			projetred.Forgeron(&p1)
		case 6:
			p1.UtiliserCompetence()
		case 0:
			fmt.Println("À bientôt Aventurier!")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}

func duel(player *projetred.User) {
	gobelin := projetred.InitGoblin()
	fmt.Println("Vous commencez un duel contre le", gobelin.Nom)
	projetred.TrainingFight(*player, gobelin)
}

func main() {
	AppelMenu()
}
