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
	a1 := make([]*Pair, len(a))
	copy(a1, a)
	quicksort(a1, 0, len(a)-1)
	printOut(a1)

	a2 := make([]*Pair, len(a))
	copy(a2, a)
	quicksortStable(a2, 0, len(a)-1)
	printOut(a2)

	a3 := make([]*Pair, len(a))
	copy(a3, a)
	quicksortStableNoRecursion(a3)
	fmt.Printf("%v\n", a3)
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
	fmt.Println(l, r)
	if l >= r {
		return
	}
	smallEnd, largeStart := partitionStable(a, l, r)
	quicksortStable(a, l, smallEnd)
	quicksortStable(a, largeStart, r)
}

func quicksortStableNoRecursion(a []*Pair) {
	stack := &Stack{}
	stack.Push(len(a) - 1)
	stack.Push(0)

	for {
		l, ok := stack.Pop()
		if !ok {
			break
		}
		r, _ := stack.Pop()

		fmt.Println(l, r)

		if l >= r {
			continue
		}

		m1, m2 := partitionStable(a, l, r)
		stack.Push(r)
		stack.Push(m2)

		stack.Push(m1)
		stack.Push(l)
	}
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
		if a[s].x > a[t].x {
			a[s], a[t] = a[t], a[s]
		}else {
			s++
			t--
		}
	}
	return t
}

func partitionStable(a []*Pair, l int, r int) (int, int) {
	p := getPivot(a, l, r)
	//fmt.Printf("p = %d\n", p)
	small := []*Pair{}
	mid := []*Pair{}
	large := []*Pair{}

	for i := l; i <= r; i++ {
		if a[i].x < p.x {
			small = append(small, a[i])
		} else if a[i].x > p.x {
			large = append(large, a[i])
		} else {
			mid = append(mid, a[i])
		}
	}
	midPos := l + len(small)
	largePos := midPos + len(mid)
	copy(a[l:midPos], small)
	copy(a[midPos:largePos], mid)
	copy(a[largePos:], large)

	return midPos - 1, largePos
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
