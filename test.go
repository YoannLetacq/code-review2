package main

import (
	"fmt"
	"os"
	"strconv"
)

func review() {
	check_pa()
	fmt.Println()
	check_pb()
	fmt.Println()
	check_sa()
	fmt.Println()
	check_sb()
}

// this file don't comments, i"m checking if functions works properly

func check_sb() {
	stack_b := &Stack_b{}

	for i, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error sb converting")
			return
		}
		if i >= 2 {
			break
		}
		stack_b.Push(num)
	}

	fmt.Println("SB stack check element")
	for node := stack_b.top; node != nil; node = node.next {

		fmt.Println("Value sb is :", node.value)
	}

	err := stack_b.sb()
	if err != nil {
		fmt.Println("Error sb", err)
		return
	}
}

func check_sa() {
	stack := &Stack_a{}

	for i, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error sa converting")
			return
		}
		if i >= 2 {
			break
		}
		stack.Push(num)
	}

	fmt.Println("SA stack a check element")
	count := 0
	for node := stack.top; node != nil; node = node.next {
		fmt.Println("Value sa is", node.value)
		count++
		if count == 2 {
			break
		}
	}
	err := stack.sa()
	if err != nil {
		fmt.Println("Error sa", err)
		return
	}
}

func check_pb() {
	stack := &Stack_a{}
	stack_b := &Stack_b{}

	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error pb converting")
			return
		}
		stack_b.Push(num)
	}

	for !stack_b.isEmpty() {
		value, err := stack_b.Pop()
		if err != nil {
			fmt.Println("PB Error pop stack_b")
			return
		}
		stack.Push(value)
	}

	fmt.Println("PB stack _ BEFORE")

	for node := stack.top; node != nil; node = node.next {
		fmt.Println("value : ", node.value)
	}

	err := stack_b.pb(stack)
	if err != nil {
		fmt.Println("Error pa main", err)
		return
	}

	println("PB stack _ AFTER")

	for node := stack_b.top; node != nil; node = node.next {
		fmt.Println("value: ", node.value)
	}

}

func check_pa() {
	stack := &Stack_a{}
	stack_b := &Stack_b{}

	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error pa strconv")
			return
		}
		stack.Push(num)
	}

	for !stack.isEmpty() {
		value, err := stack.Pop()
		if err != nil {
			fmt.Println("PA Error pop stack_a:", err)
			return
		}
		stack_b.Push(value)
	}

	fmt.Println("PA stack _ BEFORE")

	for node := stack_b.top; node != nil; node = node.next {
		fmt.Println("value : ", node.value)
	}

	err := stack.pa(stack_b)
	if err != nil {
		fmt.Println("Error pa main", err)
		return
	}

	println("PA stack _ AFTER")

	for node := stack.top; node != nil; node = node.next {
		fmt.Println("value: ", node.value)
	}
}
