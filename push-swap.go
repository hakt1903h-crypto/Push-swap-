package main

import (
	"fmt"
	"os"
)

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
	p.trier()

	for _, op := range p.operations {
		fmt.Println(op)
	}
}
