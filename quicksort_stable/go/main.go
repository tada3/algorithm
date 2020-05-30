package main

import (
	"fmt"
)

type Pair struct {
	x int
	y int
}

func (p *Pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
	println("hello")
	a := []*Pair{
		&Pair{2, 0},
		&Pair{2, 1},
		&Pair{1, 0},
		&Pair{3, 0},
		&Pair{4, 0},
		&Pair{2, 2},
	}
	test(a)

}

func test(a []*Pair) {

	a2 := make([]*Pair, len(a))
	copy(a2, a)
	quicksort(a2, 0, len(a)-1)
	printOut(a2)
}

func printOut(a []*Pair) {
	if len(a) == 0 {
		fmt.Println()
		return
	}

	out := ""
	for _, p := range a {
		out += (p.String() + ", ")
	}
	fmt.Println(out[:len(out)-2])
}

func quicksort(a []*Pair, l int, r int) {
	//fmt.Println(l, r)
	if l >= r {
		return
	}
	m := partition(a, l, r)
	quicksort(a, l, m)
	quicksort(a, m+1, r)
}

func quicksortStable(a []*Pair, l int, r int) {
	//fmt.Println(l, r)
	if l >= r {
		return
	}
	m := partition(a, l, r)
	quicksort(a, l, m)
	quicksort(a, m+1, r)
}

func partition(a []*Pair, l int, r int) int {
	p := getPivot(a, l, r)
	//fmt.Printf("p = %d\n", p)
	s := l
	t := r
	for {
		for a[s].x < p.x {
			s++
		}
		for a[t].x > p.x {
			t--
		}
		if s >= t {
			break
		}
		a[s], a[t] = a[t], a[s]
		s++
		t--
	}
	return t
}

func partitionStable(a []*Pair, l int, r int) int {
	p := getPivot(a, l, r)
	//fmt.Printf("p = %d\n", p)
	s := l
	t := r
	for {
		for a[s].x < p.x {
			s++
		}
		for a[t].x > p.x {
			t--
		}
		if s >= t {
			break
		}
		a[s], a[t] = a[t], a[s]
		s++
		t--
	}
	return t
}

func getPivot(a []*Pair, l int, r int) *Pair {
	m := l + (r-l)/2
	x := a[l]
	y := a[m]
	z := a[r]
	if x.x < y.x {
		if y.x < z.x {
			return y
		} else if x.x < z.x {
			return z
		}
		return x
	} else {
		if x.x < z.x {
			return y
		} else if y.x < z.x {
			return z
		}
		return y
	}
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
