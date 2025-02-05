package main

import "fmt"

// max stack size
const MaxSize = 100

// Stack struct with array and a top index
type Stack struct {
    data [MaxSize]int // this array stores stack elements
    top  int          // keep track of the top index
}

// fun to push element to stack
func (s *Stack) Push(value int) {
    if s.top >= MaxSize { // check overflow
        fmt.Println("stack overflow!!! Cannot push", value)
        return
    }
    s.data[s.top] = value // Store value to top
    s.top++               // update pointer
}

// this is pop function
func (s *Stack) Pop() int {
    if s.top == 0 { // if stack is empty?
        fmt.Println("Nothing to pop")
        return -1 
    }
    s.top--              // Move pointer down
    return s.data[s.top] // Return last stored value
}


func main() {
    stack := Stack{} 

    // lets push some values onto the stack
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    // Pop values and print 
    fmt.Println("Popped:", stack.Pop()) // 30
    fmt.Println("Popped:", stack.Pop()) // 20
    fmt.Println("Popped:", stack.Pop()) // 10

   
    popped := stack.Pop()
    if popped != -1 { // Only print valid values
        fmt.Println("Popped:", popped)
    }
}
