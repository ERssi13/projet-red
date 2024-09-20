package projetred

import (
	"fmt"
	"strings"
)

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding > 0 {
		return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
	}
	return text
}

func AfficherLigne(label, value string, totalWidth int) {
	const col1Width = 20
	col2Width := totalWidth - col1Width - 4

	if len(label) > col1Width {
		label = label[:col1Width]
	}
	if len(value) > col2Width {
		value = value[:col2Width]
	}

	fmt.Printf("║ %-20s : %-43s ║\n", label, value)
}