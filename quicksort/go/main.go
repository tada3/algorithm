package main

import (
	"fmt"
	"sort"
)

func main() {
	println("hello")


	a := []int{2, 3, 5, 1, 8, 4}
	test(a)

	a = []int{2, 2, 2, 2, 2}
	test(a)

	a = []int{1, 2, 3, 1, 2, 3}
	test(a)

	a = []int{1}
	test(a)

	a = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	test(a)
	
	a = []int{4, 3, 7, 1, 5}
	test(a)

}

func test(a []int) {
	a1 := make([]int, len(a))
	copy(a1, a)
	sort.Ints(a1)
	fmt.Printf("%v\n", a1)

	a2 := make([]int, len(a))
	copy(a2, a)
	quicksort(a2, 0, len(a)-1)
	fmt.Printf("%v\n", a2)

	a3 := make([]int, len(a))
	copy(a3, a)
	quicksortNoRecursion(a3)
	fmt.Printf("%v\n", a3)

}

func quicksort(a []int, l int, r int) {
	//fmt.Println(l, r)
	if l >= r {
		return
	}
	m := partition(a, l, r)
	quicksort(a, l, m-1)
	quicksort(a, m+1, r)
}

// Queue requires O(N) space
// Stack requires O(log N) space
func quicksortNoRecursion(a []int) {
	stack := Stack{}
	stack.Push(len(a) - 1)
	stack.Push(0)

	for {
		l, ok := stack.Pop()
		if !ok {
			break
		}
		r, _ := stack.Pop()
		//fmt.Println(l, r)
		if l >= r {
			continue
		}

		m := partition(a, l, r)
		stack.Push(m)
		stack.Push(l)
		stack.Push(r)
		stack.Push(m + 1)
	}
}

func partition(a []int, l int, r int) int {
	p := getPivot(a, l, r)
	fmt.Printf("p = %d\n", p)
	s := l
	t := r
	for {
		fmt.Println("s, t", s, t)
		for a[s] < p {
			s++
		}
		for a[t] > p {
			t--
		}
		if s >= t {
			break
		}
		if a[s] > a[t] {
			a[s], a[t] = a[t], a[s]
		} else {
			s++
			t--
		}
		
	}
	fmt.Println("a = ", a)
	fmt.Println("bunki = ", t)
	return t
}

func getPivot(a []int, l int, r int) int {
	m := l + (r-l)/2
	x := a[l]
	y := a[m]
	z := a[r]
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		}
		return x
	} else {
		if x < z {
			return y
		} else if y < z {
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
