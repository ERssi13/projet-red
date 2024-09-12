package main

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
	Or         int
	PvMax      int
	PvActuel   int
	Inventaire map[string]int
}

func Init(nom string, classe string, niveau int, or int, pvMax int, pvActuel int, inventaire map[string]int) User {
	return User{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		Or:         or,
		PvMax:      pvMax,
		PvActuel:   pvActuel,
		Inventaire: inventaire,
	}
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding > 0 {
		return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
	}
	return text
}

func afficherLigne(label, value string, totalWidth int) {
	const col1Width = 20
	col2Width := totalWidth - col1Width - 4 // 4 is for borders and the space between label and value

	if len(label) > col1Width {
		label = label[:col1Width]
	}
	if len(value) > col2Width {
		value = value[:col2Width]
	}

	fmt.Printf("║ %-20s : %-43s ║\n", label, value)
}

func (u User) AfficherInfos() {
	const largeurTotale = 70

	fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
	fmt.Println("║" + centerText("Informations du Personnage", largeurTotale-2) + "║")
	fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

	afficherLigne("Nom", u.Nom, largeurTotale)
	afficherLigne("Classe", u.Classe, largeurTotale)
	afficherLigne("Niveau", fmt.Sprintf("%d", u.Niveau), largeurTotale)
	afficherLigne("Or", fmt.Sprintf("%d", u.Or), largeurTotale)
	afficherLigne("Points de vie max", fmt.Sprintf("%d", u.PvMax), largeurTotale)
	afficherLigne("Points de vie actuels", fmt.Sprintf("%d", u.PvActuel), largeurTotale)

	fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")
}

func (u User) AfficherInventaire() {
	const largeurTotale = 70

	fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
	fmt.Println("║" + centerText("Inventaire", largeurTotale-2) + "║")
	fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

	if len(u.Inventaire) == 0 {
		fmt.Println("║ " + centerText("L'inventaire est vide.", largeurTotale-2) + "║")
	} else {
		for item, quantity := range u.Inventaire {
			afficherLigne(item, fmt.Sprintf("%d", quantity), largeurTotale)
		}
	}

	fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")
}

func afficherMenu() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         Menu Principal                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Afficher les informations du personnage                       ║")
	fmt.Println("║ 2. Afficher l'Inventaire                                         ║")
	fmt.Println("║ 3. Quitter                                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

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

	fmt.Print("Entrez le nom de votre Personnage: ")
	scanner.Scan()
	nom := scanner.Text()
	nom = strings.TrimSpace(nom)

	classe := choisirClasse()

	inventaire := map[string]int{
		"Potion":     3,
		"Iron sword": 1,
	}

	p1 := Init(nom, classe, 1, 100, 100, 40, inventaire)

	var choix int
	for {
		afficherMenu()
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

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
