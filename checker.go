package main

import (
	"bufio"
	"fmt"
	"os"
)

func executerLigne(p *Piles, ligne string) bool {
	switch ligne {
	case "sa":
		p.sa()
	case "sb":
		p.sb()
	case "ss":
		p.ss()
	case "pa":
		p.pa()
	case "pb":
		p.pb()
	case "ra":
		p.ra()
	case "rb":
		p.rb()
	case "rr":
		p.rr()
	case "rra":
		p.rra()
	case "rrb":
		p.rrb()
	case "rrr":
		p.rrr()
	default:
		return false
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	nums, ok := lireArgs(os.Args[1:])
	if !ok {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	p := Piles{a: nums}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		if !executerLigne(&p, ligne) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if p.estValide() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
