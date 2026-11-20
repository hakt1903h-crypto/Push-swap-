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
