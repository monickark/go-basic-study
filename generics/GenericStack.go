package main

type Stack[T any] struct {
	list []T
}

func (s *Stack[T]) Push(x T) {
	s.list = append(s.list, x)
}

func (s *Stack[T]) Pop() T {
	val := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return val
}

func (s *Stack[T]) peek() T {

}

func main() {

}
