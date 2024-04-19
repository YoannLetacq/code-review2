package main

import "errors"

func (s_a *Stack_a) Push(value int) {
	newNode := &Node{value: value, next: s_a.top}
	s_a.top = newNode
}

func (s_b *Stack_b) Push(value int) {
	newNode := &Node{value: value, next: s_b.top}
	s_b.top = newNode
}

func (s_a *Stack_a) Pop() (int, error) {
	if s_a.isEmpty() {
		return 0, errors.New("Stack_a is empty")
	}
	value := s_a.top.value
	s_a.top = s_a.top.next

	return value, nil
}

func (s_b *Stack_b) Pop() (int, error) {
	if s_b.isEmpty() {
		return 0, errors.New("Stack_b is empty")
	}
	value := s_b.top.value
	s_b.top = s_b.top.next
	return value, nil
}

func (s_a *Stack_a) Top() (int, error) {
	if s_a.isEmpty() {
		return 0, errors.New("Stack_a is empty")
	}
	return s_a.top.value, nil
}

func (s_b *Stack_b) Top() (int, error) {
	if s_b.isEmpty() {
		return 0, errors.New("Stack_b is empty")
	}
	return s_b.top.value, nil
}

func (s *Stack_a) isEmpty() bool {
	return s.top == nil
}

func (s *Stack_b) isEmpty() bool {
	return s.top == nil
}
