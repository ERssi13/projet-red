package projetred

import ("fmt"
		"strings"
)

func AfficherMenu() {
	clearTerminal()
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         Menu Principal                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Duel                                                          ║")
	fmt.Println("║ 2. Afficher les informations du personnage                       ║")
	fmt.Println("║ 3. Afficher l'Inventaire                                         ║")
	fmt.Println("║ 4. Marchand                                                      ║")
	fmt.Println("║ 5. Forgeron                                                      ║")
	fmt.Println("║ 0. Quitter                                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

func (u User) AfficherInfos() {
	const largeurTotale = 70

	for {
		clearTerminal()
		fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
		fmt.Println("║" + centerText("Informations du Personnage", largeurTotale-2) + "║")
		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")

		AfficherLigne("Nom", u.Nom, largeurTotale)
		AfficherLigne("Classe", u.Classe, largeurTotale)
		AfficherLigne("Niveau", fmt.Sprintf("%d", u.Niveau), largeurTotale)
		AfficherLigne("Points de vie", fmt.Sprintf("%d/%d", u.PvActuel, u.PvMax), largeurTotale)
		AfficherLigne("Or", fmt.Sprintf("%d", u.Or), largeurTotale)

		fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")
		fmt.Println("╔" + strings.Repeat("═", largeurTotale-2) + "╗")
		fmt.Println("║" + centerText("Equipements", largeurTotale-2) + "║")
		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")
		AfficherLigne("Tête", u.getNomEquipement("Chapeau"), largeurTotale)
		AfficherLigne("Torse", u.getNomEquipement("Tunique"), largeurTotale)
		AfficherLigne("Jambes", u.getNomEquipement("Bottes"), largeurTotale)
		fmt.Println("╠" + strings.Repeat("═", largeurTotale-2) + "╣")
		fmt.Println("║ 0. Retourner au menu                                               ║")
		fmt.Println("╚" + strings.Repeat("═", largeurTotale-2) + "╝")
		var choix int
		fmt.Print("Choisissez une option: ")
		fmt.Scan(&choix)

		if choix == 0 {
			return
		} else {
			fmt.Println("Option invalide.")
		}
	}
}

