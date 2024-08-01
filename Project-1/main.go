package main

import (
	"fmt"
)

func main() {
	var balance float64 = 2000.0
	var choice int
	var amount float64

	fmt.Printf("Hesabınızda %.2f TL var.\n", balance)
	fmt.Println("Para mı çekmek istiyorsunuz yoksa para mı yatırmak istiyorsunuz?")
	fmt.Println("1. Para çekmek istiyorum")
	fmt.Println("2. Para yatırmak istiyorum")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Ne kadar çekmek istiyorsunuz?")
		fmt.Scan(&amount)
		if amount > balance || amount < 0 {
			fmt.Println("Bu işlem yapılamaz.")
		} else {
			balance -= amount
			fmt.Printf("Hesabınızda artık %.2f TL var.\n", balance)
		}
	case 2:
		fmt.Println("Ne kadar yatırmak istiyorsunuz?")
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Println("Bu işlem yapılamaz.")
		} else {
			balance += amount
			fmt.Printf("Hesabınızda artık %.2f TL var.\n", balance)
		}
	default:
		fmt.Println("Geçersiz seçenek.")
	}
}
