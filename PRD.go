/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Nom        string
	Classe     string
	Niveau     int
	PvMax      int
	PvActuel   int
	Inventaire map[string]int
}

func Init(nom string, classe string, niveau int, pvMax int, pvActuel int, inventaire map[string]int) User {
	return User{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PvMax:      pvMax,
		PvActuel:   pvActuel,
		Inventaire: inventaire,
	}
}

func (u User) AfficherInfos() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Informations du Personnage                    ║")
	fmt.Printf("║ Nom : %v                                                          ║\n", u.Nom)
	fmt.Printf("║ Classe : %v                                                       ║\n", u.Classe)
	fmt.Printf("║ Niveau :	%v										                ║\n", u.Niveau)
	fmt.Printf("║ Points de vie maximum : %v						      ║\n", u.PvMax)
	fmt.Printf("║ Points de vie actuels : %v                                        ║\n", u.PvActuel)
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

func (u User) AfficherInventaire() {
	fmt.Println("=== Inventaire ===")
	if len(u.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
	} else {
		for item, quantity := range u.Inventaire {
			fmt.Printf("%s: %d\n", item, quantity)
		}
	}
	fmt.Println("===================")
}

func afficherMenu() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         Menu Principal                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Afficher les informations du personnage                       ║")
	fmt.Println("║ 2. Afiicher l'Inventaire                                         ║")
	fmt.Println("║ 3. Quitter                                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

// Nouvelle fonction pour choisir la classe
func choisirClasse() string {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                   Choisissez une classe                          ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Guerrier                                                      ║")
	fmt.Println("║ 2. Mage                                                          ║")
	fmt.Println("║ 3. Archer                                                        ║")
	fmt.Println("║ 4. Voleur                                                        ║")
	fmt.Println("║ 5. Assassin                                                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ Entrez le numéro de la classe:                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
	var choix int
	fmt.Scan(&choix)

	// Utilisation du switch pour définir la classe
	switch choix {
	case 1:
		return "Guerrier"
	case 2:
		return "Mage"
	case 3:
		return "Archer"
	case 4:
		return "Voleur"
	case 5:
		return "Assassin"
	default:
		fmt.Println("Choix invalide, la classe par défaut 'Aventurier' sera utilisée.")
		return "Aventurier"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Demander le nom du personnage
	fmt.Print("Entrez le nom de votre Personnage: ")
	scanner.Scan()
	nom := scanner.Text()
	nom = strings.TrimSpace(nom)

	// Utiliser la fonction choisirClasse pour définir la classe
	classe := choisirClasse()

	// Initialiser l'inventaire avec 3 potions
	inventaire := map[string]int{
		"Potion": 3,
	}

	// Initialiser le personnage
	p1 := Init(nom, classe, 1, 100, 40, inventaire)

	// Boucle principale pour afficher le menu
	var choix int
	for {
		afficherMenu()
		fmt.Println("")
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)
		fmt.Println("")
		switch choix {
		case 1:
			p1.AfficherInfos()
		case 2:
			p1.AfficherInventaire()
		case 3:
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
*/