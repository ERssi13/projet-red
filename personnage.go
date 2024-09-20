package projetred

type User struct {
	Nom                   string
	Classe                string
	Niveau                int
	Or                    int
	PvMax                 int
	PvActuel              int
	Inventaire            map[string]int
	ResourcesJ            map[string]int
	Skills                []Skill
	hasReceivedFreePotion bool
	Maxinventaire         int
	Equipement  		  map[string]Equipment  
	Equipements 		  Equipemnt
	Degats				  int
}

func Init(nom string, classe string, niveau int, or int, pvMax int, pvActuel int, inventaire map[string]int, resourcesj map[string]int, hasreceivedFreePotion bool, maxinventaire int,degats int) User {
	skillsDeBase := []Skill{
		{Nom: "Coup de poing", Degats: 10, Achete: false},
	}

	return User{
		Nom:                   nom,
		Classe:                classe,
		Niveau:                niveau,
		Or:                    or,
		PvMax:                 pvMax,
		PvActuel:              pvActuel,
		Inventaire:            inventaire,
		ResourcesJ:            resourcesj,
		Skills:                skillsDeBase,
		Equipement:            make(map[string]Equipment), 
		Equipements:           Equipemnt{},         
		hasReceivedFreePotion: hasreceivedFreePotion,
		Maxinventaire:         maxinventaire,
		Degats: 			   degats,		
	}
}

