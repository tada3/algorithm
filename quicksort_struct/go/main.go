package main

import (
	"fmt"
	"sort"
)

type pair struct {
    a, b int
}

func less(x, y pair) bool {
	if x.a < y.a {
		return true
	}
	if x.a == y.a {
		return x.b < y.b
	}
	return false
}

func main() {
	println("hello")
	a := []int{2, 3, 5, 1, 8, 4}
	b := []int{2, 3, 5, 1, 8, 4}
	test(a, b)

	a = []int{2, 2, 2, 2, 2}
	b = []int{2, 2, 2, 2, 2}
	test(a, b)

	a = []int{1, 2, 3, 1, 2, 3}
	b = []int{1, 2, 3, 1, 2, 3}
	test(a, b)

	a = []int{1}
	b = []int{1}
	test(a,b)

	a = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	b = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	test(a, b)

	a = []int{4, 3, 7, 1, 5}
	b = []int{5, 4, 3, 2, 1}
	test(a, b)

	a = []int{4, 4, 4, 4, 4}
	b = []int{5, 4, 9, 2, 1}
	test(a, b)
}

func test(a,b []int) {
	a1 := make([]int, len(a))
	copy(a1, a)
	sort.Ints(a1)
	fmt.Printf("%v\n", a1)

	a2 := make([]pair, len(a))
	makePair(a2, a, b)
	sortPair(a2)
	fmt.Printf("%v\n", a2)
}


func makePair(dest []pair, srcA, srcB []int) {
	for i, v := range srcA {
		dest[i] = pair{v, srcB[i]}
	}
}

func sortPair(a []pair) {
	quicksort2(a, 0, len(a)-1)
}


func quicksort2(a []pair, l int, r int) {
	fmt.Println("quicksort2", l, r)
	if r <= l {
		return
	}
	m := partition2(a, l, r)
	quicksort2(a, l, m)
	quicksort2(a, m+1, r)
}

func partition2(a []pair, l int, r int) int {
	p := getPivot2(a, l, r)
	fmt.Printf("p = %d\n", p)
	s := l
	t := r
	for {
		for less(a[s], p) {
			s++
		}
		for less(p, a[t]) {
			t--
		}
		if t <= s {
			break
		}
		fmt.Println("s, t = ",s, t)
		fmt.Println("a = ", a)
		if less(a[t], a[s]) {
			a[s], a[t] = a[t], a[s]
		} else {
			s++
			t--
		}
	}
	return t
}

func getPivot2(a []pair, l int, r int) pair {
	m := l + (r-l)/2
	x := a[l]
	y := a[m]
	z := a[r]
	if less(x, y) {
		if less(y, z) {
			return y
		} else if less(x,z) {
			return z
		}
		return x
	} else {
		if less(x , z) {
			return y
		} else if less (y , z) {
			return z
		}
		return y
	}
}
