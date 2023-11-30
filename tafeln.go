// Für Lernzwecke, wenn die "Dein Star" Kodier/Dekodiertabelle genutzt wird.

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "Anzahl der Tafeln")
	flag.Parse()

	for i := 1; i <= n; i++ {
		fmt.Printf("Tafel %d:\n", i)
		for j := 0; j < 10; j++ {
			for k := 0; k < 5; k++ {
				num, _ := rand.Int(rand.Reader, big.NewInt(100000))
				fmt.Printf("%05d ", num.Int64())
			}
			fmt.Println()
		}
		if i != n {
			fmt.Println()
		}
	}
}

