package projetred

import "fmt"

func (u *User) Dead() {
	if u.PvActuel <= 0 {
		fmt.Printf("%s est mort. Il est ressuscité avec 50%% de sa vie maximale.\n", u.Nom)
		u.PvActuel = u.PvMax / 2
	} else {
		fmt.Printf("%s est encore vivant avec %d points de vie.\n", u.Nom, u.PvActuel)
	}
}
