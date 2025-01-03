package main

import (
	"flag"
	"fmt"
	"github.com/google/go-tpm/legacy/tpm2"
	"os"
)

func main() {
	var n, gpl, lps int
	var l, d, b bool
	flag.IntVar(&n, "n", 1, "Number of sheets")
	flag.BoolVar(&d, "d", false, "Digits")
	flag.BoolVar(&l, "l", false, "Letters")
	flag.BoolVar(&b, "b", false, "Binary")
	flag.IntVar(&gpl, "gpl", 5, "Groups per line")
	flag.IntVar(&lps, "lps", 10, "Lines per sheet")
	flag.Parse()

	if l == d && d == b {
		fmt.Println("Usage:")
		fmt.Println("  -n Number of sheets")
		fmt.Println("  -d Digits")
		fmt.Println("  -l Letters")
		fmt.Println("  -b Binary")
		fmt.Println("  -gpl Groups per line - default 5")
		fmt.Println("  -lps Lines per sheet - default 10")
		os.Exit(1)
	}

	rwc, err := tpm2.OpenTPM()
	if err != nil {
		fmt.Printf("TPM öffnen fehlgeschlagen: %v\n", err)
		return
	}
	defer rwc.Close()

	for i := 1; i <= n; i++ {
		fmt.Printf("*DESTROY AFTER USE*\n")
		for j := 0; j < lps; j++ {
			for k := 0; k < gpl; k++ {
				if l {
					random, _ := tpm2.GetRandom(rwc, 5)
					for m := 0; m < 5; m++ {
						fmt.Printf("%c", 'A'+(random[m]%26))
					}
				}
				if d {
					random, _ := tpm2.GetRandom(rwc, 3)
					num := uint32(random[0]) | uint32(random[1])<<8 | uint32(random[2])<<16
					fmt.Printf("%05d", num%100000)
				}
				if b {
					random, _ := tpm2.GetRandom(rwc, 1)
					for m := 0; m < 5; m++ {
						fmt.Printf("%01d", (random[0]>>m)&1)
					}
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
