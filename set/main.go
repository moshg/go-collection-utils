package set

import "iter"

type Set[T comparable] struct {
	elems map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elems: make(map[T]struct{})}
}

func (s *Set[T]) Len() int {
	return len(s.elems)
}

func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.elems[elem]
	return ok
}

func (s *Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s.elems[elem] = struct{}{}
	}
}

func (s *Set[T]) Remove(elems ...T) {
	for _, elem := range elems {
		delete(s.elems, elem)
	}
}

func (s *Set[T]) Clear() {
	s.elems = make(map[T]struct{})
}

func (s *Set[T]) Range() iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}
