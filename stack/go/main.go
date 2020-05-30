package main

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(9)
	println(s.Size())
	v, ok := s.Pop()
	if !ok {
		println("Empty!")
	} else {
		println(v)
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
