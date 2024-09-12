
	fmt.Print("Entrez le nom de votre Personnage: ")
	scanner.Scan()
	nom := scanner.Text()
	nom = strings.TrimSpace(nom)

	classe := choisirClasse()

	inventaire := map[string]int{
		"Potion":     3,
		"Iron sword": 1,
	}
