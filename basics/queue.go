package basics

import "fmt"

//Queue basic queue in go
func Queue() {
	//create a stack
	queue := make([]int, 0)
	fmt.Println("create", queue)

	//push on stack
	queue = append(queue, 1)
	queue = append(queue, 2)
	fmt.Println("enqueue", queue)

	//pop from stack
	poppedElement := queue[0]
	fmt.Println("dequeued element", poppedElement)

	queue = queue[1:]
	fmt.Println("dequeue", queue)
}
