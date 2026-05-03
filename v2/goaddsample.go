// A simple package with an add function	
package goaddsample

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two numbers together
// If you want to learn more about adding, read https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a,b T) T {
	return a+b
}
