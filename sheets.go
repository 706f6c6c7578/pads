// For educational purposes, when using the Diana Cryptosystem.

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "Number of sheets")
	flag.Parse()

	for i := 1; i <= n; i++ {
		fmt.Printf("Sheet %d:\n", i)
		for j := 0; j < 15; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 5; l++ {
					num, _ := rand.Int(rand.Reader, big.NewInt(26))
					fmt.Printf("%c", 'A'+num.Int64())
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		if i != n {
			fmt.Println()
		}
	}
}

