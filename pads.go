package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var n, gpl, lps int
	var l, d bool
	flag.IntVar(&n, "n", 1, "Number of sheets")
	flag.BoolVar(&d, "d", false, "Digits")
	flag.BoolVar(&l, "l", false, "Letters")
	flag.IntVar(&gpl, "gpl", 5, "Groups per line")
	flag.IntVar(&lps, "lps", 10, "Lines per sheet")
	flag.Parse()

	if l == d {
		fmt.Println("Usage:")
		fmt.Println("  -n Number of sheets")
		fmt.Println("  -d Digits")
		fmt.Println("  -l Letters")
		fmt.Println("  -gpl Groups per line - default 5")
		fmt.Println("  -lps Lines per sheet - default 10")
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("Sheet %d:\n", i)
		for j := 0; j < lps; j++ {
			for k := 0; k < gpl; k++ {
				if l {
					for m := 0; m < 5; m++ {
						num, _ := rand.Int(rand.Reader, big.NewInt(26))
						fmt.Printf("%c", 'A'+num.Int64())
					}
				}
				if d {
					num, _ := rand.Int(rand.Reader, big.NewInt(100000))
					fmt.Printf("%05d", num.Int64())
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

