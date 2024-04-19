package main

import (
	"errors"
)

// Pop topB value on the stack_a
func (s_a *Stack_a) pa(s_b *Stack_b) error {
	if s_b.isEmpty() {
		return errors.New("stack b (pa) is empty")
	}
	topB, err := s_b.Pop()
	if err != nil {
		return err
	}
	s_a.Push(topB)
	return nil
}

// Pop topA value on the stack_b
func (s_b *Stack_b) pb(s_a *Stack_a) error {
	if s_a.isEmpty() {
		return errors.New("stack_a (pb) is empty")
	}
	topA, err := s_a.Pop()
	if err != nil {
		return err
	}
	s_b.Push(topA)
	return nil
}

// This method swap first element on stack_a
func (s_a *Stack_a) sa() error {
	if s_a.top == nil || s_a.top.next == nil {
		return errors.New("stack a (sa) is empty <2 elements")
	}

	first := s_a.top.value
	second := s_a.top.next.value

	s_a.top.value = second
	s_a.top.next.value = first
	return nil
}

// This method swap first element on stack_b
func (s_b *Stack_b) sb() error {

	if s_b.top == nil || s_b.top.next == nil {
		return errors.New("stack b (sb) is empty <2 elements")
	}

	first := s_b.top.value
	second := s_b.top.next.value

	s_b.top.next.value = first
	s_b.top.value = second

	return nil
}
