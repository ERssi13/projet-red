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
	nom := scanner.Text()
	nom = strings.TrimSpace(nom)

	classe := projetred.ChoisirClasse()

	inventaire := map[string]int{
		"Potion de Vie": 3,
		"Iron sword":    1,
	}
	//resourcej := map[string]int {
		//"Fourrure" : 1,
		//"Dent" : 5,
	//}

	p1 := projetred.Init(nom, classe, 1, 100, 100, 40, inventaire, false, 10)

	var choix int
	for {
		projetred.AfficherMenu()
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			p1.AfficherInfos()
		case 2:
			p1.AccesInventaire()
		case 3:
			p1.Marchand()
		case 0:
			fmt.Println("A bientot Aventurier!")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}

func main() {
	AppelMenu()
}
