package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) sb() {
	if len(p.b) < 2 {
		return
	}
	p.b[0], p.b[1] = p.b[1], p.b[0]
	p.operations = append(p.operations, "sb")
	}

func (p *Piles) ss() {
	if len(p.a) < 2 && len(p.b) < 2 {
		return
	}
	if len(p.a) >= 2 {
		p.a[0], p.a[1] = p.a[1], p.a[0]
	}
	if len(p.b) >= 2 {
		p.b[0], p.b[1] = p.b[1], p.b[0]
	}
	p.operations = append(p.operations, "ss")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}
