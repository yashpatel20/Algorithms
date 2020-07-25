package basics

import (
	"fmt"
)

//Stack basic stack in go
func Stack() {
	//create a stack
	stack := make([]int, 0)
	fmt.Println("create", stack)

	//push on stack
	stack = append(stack, 1)
	stack = append(stack, 2)
	fmt.Println("push", stack)

	//pop from stack
	poppedElement := stack[len(stack)-1]
	fmt.Println("popped element", poppedElement)

	stack = stack[:len(stack)-1]
	fmt.Println("pop", stack)

}
