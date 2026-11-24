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
func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) rb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append(p.b[1:], p.b[0])
	p.operations = append(p.operations, "rb")
}

func (p *Piles) rr() {
	if len(p.a) < 2 && len(p.b) < 2 {
		return
	}
	if len(p.a) >= 2 {
		p.a = append(p.a[1:], p.a[0])
	}
	if len(p.b) >= 2 {
		p.b = append(p.b[1:], p.b[0])
	}
	p.operations = append(p.operations, "rr")
}

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	dernier := p.a[len(p.a)-1]
	p.a = append([]int{dernier}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	dernier := p.b[len(p.b)-1]
	p.b = append([]int{dernier}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

func (p *Piles) rrr() {
	if len(p.a) < 2 && len(p.b) < 2 {
		return
	}
	if len(p.a) >= 2 {
		dernierA := p.a[len(p.a)-1]
		p.a = append([]int{dernierA}, p.a[:len(p.a)-1]...)
	}
	if len(p.b) >= 2 {
		dernierB := p.b[len(p.b)-1]
		p.b = append([]int{dernierB}, p.b[:len(p.b)-1]...)
	}
	p.operations = append(p.operations, "rrr")
}

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}

func (p *Piles) estValide() bool {
	return p.aEstTriee() && len(p.b) == 0
}

func (p *Piles) posMin() int {
	minIdx := 0
	for i := 1; i < len(p.a); i++ {
		if p.a[i] < p.a[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

func (p *Piles) trierTrois() {
	if len(p.a) != 3 || p.aEstTriee() {
		return
	}
	x, y, z := p.a[0], p.a[1], p.a[2]
	if x < z && z < y {
		p.sa()
		p.ra()
	} else if y < x && x < z {
		p.sa()
	} else if z < x && x < y {
		p.rra()
	} else if y < z && z < x {
		p.ra()
	} else if z < y && y < x {
		p.sa()
		p.rra()
	}
}

func (p *Piles) trier() {
	if p.aEstTriee() {
		return
	}
	for len(p.a) > 3 && !p.aEstTriee() {
		idx := p.posMin()
		if idx < len(p.a)/2 {
			for idx > 0 {
				p.ra()
				idx--
			}
		} else {
			for idx < len(p.a) {
				p.rra()
				idx++
			}
		}
		if p.aEstTriee() {
			break
		}
		p.pb()
	}
	if len(p.a) == 3 {
		p.trierTrois()
	} else if len(p.a) == 2 && !p.aEstTriee() {
		p.sa()
	}
	for len(p.b) > 0 {
		p.pa()
	}
}
func lireArgs(args []string) ([]int, bool) {
	if len(args) == 0 {
		return nil, true
	}
	var res []int

	for _, arg := range args {
		morceaux := strings.Fields(arg)
		if len(morceaux) == 0 {
			return nil, false
		}
		for _, m := range morceaux {
			n, err := strconv.Atoi(m)
			if err != nil {
				return nil, false
			}
			for _, existant := range res {
				if existant == n {
					return nil, false
				}
			}
			res = append(res, n)
		}
	}
	return res, true
}
