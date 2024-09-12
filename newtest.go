/*package main

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

// Initialisez votre personnage
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

// Centrer le texte dans une ligne de largeur définie
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding > 0 {
		return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
	}
	return text
}

// Fonction pour afficher une ligne de texte alignée avec deux colonnes
func afficherLigne(label, value string, totalWidth int) {
	// Définir la largeur de la première colonne (label) et la deuxième colonne (valeur)
	const col1Width = 20
	col2Width := totalWidth - col1Width - 4 // 4 est l'espace pour les bords et les espaces

	// Si la chaîne dépasse, elle sera coupée
	if len(label) > col1Width {
		label = label[:col1Width]
	}
	if len(value) > col2Width {
		value = value[:col2Width]
	}

	// Afficher la ligne formatée
	fmt.Printf("║ %-20s : %-45s ║\n", label, value)
}

// Méthode pour afficher les informations du personnage
func (u User) AfficherInfos() {
	// Définir la largeur totale du cadre (par exemple, 70 caractères)
	const largeurTotale = 72

	// Afficher le cadre supérieur
	fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
	fmt.Println("║" + centerText("Informations du Personnage", largeurTotale-2) + "║")
	fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

	// Afficher les informations du personnage avec alignement des colonnes
	afficherLigne("Nom", u.Nom, largeurTotale)
	afficherLigne("Classe", u.Classe, largeurTotale)
	afficherLigne("Niveau", fmt.Sprintf("%d", u.Niveau), largeurTotale)
	afficherLigne("Points de vie max", fmt.Sprintf("%d", u.PvMax), largeurTotale)
	afficherLigne("Points de vie actuels", fmt.Sprintf("%d", u.PvActuel), largeurTotale)

	// Afficher le cadre inférieur
	fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")
}

// Méthode pour afficher l'inventaire
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

// Fonction pour afficher le menu principal
func afficherMenu() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         Menu Principal                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Afficher les informations du personnage                       ║")
	fmt.Println("║ 2. Afficher l'Inventaire                                         ║")
	fmt.Println("║ 3. Quitter                                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

// Fonction pour choisir la classe du personnage
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
		"Potion":     3,
		"Iron Sword": 1,
	}

	// Initialiser le personnage
	p1 := Init(nom, classe, 1, 100, 40, inventaire)

	// Boucle principale pour afficher le menu
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
*/