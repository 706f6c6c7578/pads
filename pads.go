package main

import (
    "flag"
    "fmt"
    "github.com/google/go-tpm/legacy/tpm2"
    "io"
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
                    // Generate unbiased letters
                    for m := 0; m < 5; m++ {
                        letter := getUnbiasedLetter(rwc)
                        fmt.Printf("%c", letter)
                    }
                }
                if d {
                    // Generate unbiased 5-digit number
                    num := getUnbiasedNumber(rwc)
                    fmt.Printf("%05d", num)
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

func getUnbiasedLetter(rwc io.ReadWriteCloser) byte {
    for {
        random, _ := tpm2.GetRandom(rwc, 1)
        // 234 is the first value that creates bias (9 * 26)
        if random[0] < 234 {
            return 'A' + (random[0] % 26)
        }
    }
}

func getUnbiasedNumber(rwc io.ReadWriteCloser) uint32 {
    maxValid := uint32(100000 * (0xFFFFFF / 100000))
    for {
        random, _ := tpm2.GetRandom(rwc, 3)
        num := uint32(random[0]) | uint32(random[1])<<8 | uint32(random[2])<<16
        if num < maxValid {
            return num % 100000
        }
    }
}
