package projetred

import "fmt"

func (u *User) spendOr(amount int) bool {
    if u.MaxInventaire() {
        fmt.Println("L'inventaire est plein. Vous ne pouvez pas dépenser d'or.")
        return false
    }
    
    if u.Or >= amount {
        u.Or -= amount
        fmt.Printf("Vous avez dépensé %d or. Il vous reste %d or.\n", amount, u.Or)
        return true
    } else {
        fmt.Println("Vous n'avez pas assez d'or pour acheter cet article.")
        return false
    }
}


func (u *User) Marchand() {
    for {
        clearTerminal()
        fmt.Printf("Or: %d\n", u.Or)
        fmt.Println()
        fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
        fmt.Println("║                         Interface du Marchand                    ║")
        fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
        fmt.Println("║ 1. Acheter une potion de vie (gratuit la première fois)          ║")
        fmt.Println("║ 2. Acheter une potion de poison (10 or)                          ║")
        fmt.Println("║ 3. Acheter une épée en bois (15 or)                              ║")
        fmt.Println("║ 4. Acheter une épée en bronze (30 or)                            ║")
        fmt.Println("║ 5. Acheter une épée en fer (50 or)                               ║")
        fmt.Println("║ 6. Acheter un Livre de Sort : Boule de Feu (50 or)               ║") 
        fmt.Println("║ 7. Augmenter la taille de l'inventaire (+10 emplacements)        ║")
        fmt.Println("║ 0. Retourner au menu                                             ║")
        fmt.Println("╚══════════════════════════════════════════════════════════════════╝")

        var choix int
        fmt.Print("Choisissez une option: ")
        fmt.Scan(&choix)

        switch choix {
        case 1:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'ajouter une potion.")
            } else if u.addInventaire("Potion de Vie", 1) {
                if !u.hasReceivedFreePotion {
                    fmt.Println("Vous avez reçu une potion de vie gratuitement.")
                    u.hasReceivedFreePotion = true
                } else if u.spendOr(5) {
                    fmt.Println("Vous avez acheté une potion de vie.")
                }
            }
        case 2:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'acheter cet objet.")
            } else if u.addInventaire("Potion de Poison", 1) && u.spendOr(10) {
                fmt.Println("Vous avez acheté une potion de poison.")
            }
        case 3:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'acheter cet objet.")
            } else if u.addInventaire("Épée en bois", 1) && u.spendOr(5) {
                fmt.Println("Vous avez acheté une épée en bois.")
            }
        case 4:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'acheter cet objet.")
            } else if u.addInventaire("Épée en bronze", 1) && u.spendOr(30) {
                fmt.Println("Vous avez acheté une épée en bronze.")
            }
        case 5:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'acheter cet objet.")
            } else if u.addInventaire("Épée en fer", 1) && u.spendOr(50) {
                fmt.Println("Vous avez acheté une épée en fer.")
            }
        case 6:
            if u.MaxInventaire() {
                fmt.Println("Votre inventaire est plein. Impossible d'ajouter un Livre de Sort.")
            } else if u.addInventaire("Livre de Sort : Boule de Feu", 1) && u.spendOr(50) {
                fmt.Println("Vous avez acheté un Livre de Sort : Boule de Feu.")
            }
        case 7:
            u.AugmenterInventaire()
        case 0:
            return
        default:
            fmt.Println("Option invalide.")
        }
    }
}
