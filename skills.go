package projetred

import "fmt"

type Skill struct {
	Nom      string
	Degats   int
	Achete   bool
}

func (s Skill) UtiliserSkill() {
    fmt.Printf("Vous utilisez %s et infligez %d points de dégâts.\n", s.Nom, s.Degats)
}

func (u *User) spellBook() {
	for _, skill := range u.Skills {
		if skill.Nom == "Boule de feu" {
			fmt.Println("Vous avez déjà appris le sort 'Boule de feu'.")
			return
		}
	}
	u.Skills = append(u.Skills, Skill{Nom: "Boule de feu", Degats: 50, Achete: true})
	fmt.Println("Vous avez appris le sort 'Boule de feu'.")
}


func (u *User) UtiliserCompetence() {
    if len(u.Skills) == 0 {
        fmt.Println("Vous n'avez aucune compétence.")
        return
    }

    fmt.Println("Compétences disponibles :")
    for i, skill := range u.Skills {
        fmt.Printf("%d. %s (Dégâts: %d)\n", i+1, skill.Nom, skill.Degats)
    }

    var choix int
    fmt.Print("Choisissez une compétence à utiliser: ")
    fmt.Scan(&choix)

    if choix > 0 && choix <= len(u.Skills) {
        u.Skills[choix-1].UtiliserSkill()
    } else {
        fmt.Println("Compétence invalide.")
    }
}
