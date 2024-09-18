package projetred

import "fmt"

func ChoisirClasse() string {
	clearTerminal()
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                   Choisissez une classe                          ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Guerrier                                                      ║")
	fmt.Println("║ 2. Mage                                                          ║")
	fmt.Println("║ 3. Archer                                                        ║")
	fmt.Println("║ 4. Voleur                                                        ║")
	fmt.Println("║ 5. Assassin                                                      ║")
	fmt.Println("║ 6. Support                                                       ║")
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
	case 6:
		return "Support"
	default:
		fmt.Println("Choix invalide, la classe par défaut 'Aventurier' sera utilisée.")
		return "Aventurier"
	}
}
