package d18

import (
	"fmt"
	"strconv"
	"strings"
)

type operator int

const (
	addition operator = iota
	multiplication
	numeric
	lbracket
	rbracket
)

// Node ..
type Node struct {
	typ operator
	val int
	fn  func(a, b int) int
}

// NodeQueue ..
type NodeQueue []*Node

// Push ..
func (q *NodeQueue) Push(n *Node) {
	*q = append(*q, n)
}

// Pop ..
func (q *NodeQueue) Pop() (*Node, bool) {
	if len(*q) == 0 {
		return nil, false
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out, true
}

// MustPop ..
func (q *NodeQueue) MustPop() *Node {
	if len(*q) == 0 {
		panic("nothing to pop")
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

// NodeStack ..
type NodeStack []*Node

// Push ..
func (s *NodeStack) Push(n *Node) {
	*s = append(*s, n)
}

// Pop ..
func (s *NodeStack) Pop() (*Node, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out, true
}

// MustPop ..
func (s *NodeStack) MustPop() *Node {
	if len(*s) == 0 {
		panic("nothing to pop")
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}

// ModifiedShuntingYard converts tokens to RPN but without operator precedence
func ModifiedShuntingYard(tokens []string) (int, error) {
	queue := make(NodeQueue, 0)
	stack := make(NodeStack, 0)

	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			// token is a number
			queue.Push(&Node{numeric, num, nil})
			continue
		}
		// token is an operator or bracket
		switch v {
		case "+":
			for len(stack) != 0 && stack[len(stack)-1].typ != lbracket {
				if popped, ok := stack.Pop(); ok {
					queue.Push(popped)
				}
			}
			stack.Push(&Node{addition, 0, func(a, b int) int { return a + b }})
		case "*":
			for len(stack) != 0 && stack[len(stack)-1].typ != lbracket {
				if popped, ok := stack.Pop(); ok {
					queue.Push(popped)
				}
			}
			stack.Push(&Node{multiplication, 0, func(a, b int) int { return a * b }})
		case "(":
			stack.Push(&Node{lbracket, 0, func(a, b int) int { panic("I am a left paren") }})
		case ")":
			for len(stack) != 0 && stack[len(stack)-1].typ != lbracket {
				if popped, ok := stack.Pop(); ok {
					queue.Push(popped)
				}
			}

			if _, ok := stack.Pop(); !ok {
				return 0, fmt.Errorf("mismatched parens")
			}
		}
	}

	for len(stack) > 0 {
		queue.Push(stack.MustPop())
	}

	for len(queue) > 0 {
		curr := queue.MustPop()
		switch curr.typ {
		case numeric:
			stack.Push(curr)
		default:
			a := stack.MustPop()
			b := stack.MustPop()
			stack.Push(&Node{numeric, curr.fn(a.val, b.val), nil})
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("what is this")
	}

	return stack[0].val, nil
}

// ParseInput ..
func ParseInput(f []byte) [][]string {
	splitLines := strings.Split(string(f), "\n")
	out := make([][]string, len(splitLines))
	for i, v := range splitLines {
		v = strings.ReplaceAll(v, " ", "")
		out[i] = strings.Split(v, "")
	}
	return out
}

// SolvePart1 ..
func SolvePart1(tokens [][]string) (int, error) {
	sum := 0
	for _, v := range tokens {
		res, err := ModifiedShuntingYard(v)
		if err != nil {
			return 0, err
		}
		sum += res
	}
	return sum, nil
}

// ModifiedShuntingYard2 converts tokens to RPN but without operator precedence
func ModifiedShuntingYard2(tokens []string) (int, error) {
	queue := make(NodeQueue, 0)
	stack := make(NodeStack, 0)

	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			// token is a number
			queue.Push(&Node{numeric, num, nil})
			continue
		}
		// token is an operator or bracket
		switch v {
		case "+":
			for len(stack) != 0 && stack[len(stack)-1].typ == addition {
				queue.Push(stack.MustPop())
			}

			stack.Push(&Node{addition, 0, func(a, b int) int { return a + b }})
		case "*":
			for len(stack) != 0 && stack[len(stack)-1].typ != lbracket {
				if popped, ok := stack.Pop(); ok {
					queue.Push(popped)
				}
			}
			stack.Push(&Node{multiplication, 0, func(a, b int) int { return a * b }})
		case "(":
			stack.Push(&Node{lbracket, 0, func(a, b int) int { panic("I am a left paren") }})
		case ")":
			for len(stack) != 0 && stack[len(stack)-1].typ != lbracket {
				if popped, ok := stack.Pop(); ok {
					queue.Push(popped)
				}
			}

			if _, ok := stack.Pop(); !ok {
				return 0, fmt.Errorf("mismatched parens")
			}
		}
	}

	for len(stack) > 0 {
		queue.Push(stack.MustPop())
	}

	for len(queue) > 0 {
		curr := queue.MustPop()
		switch curr.typ {
		case numeric:
			stack.Push(curr)
		default:
			a := stack.MustPop()
			b := stack.MustPop()
			stack.Push(&Node{numeric, curr.fn(a.val, b.val), nil})
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("what is this")
	}

	return stack[0].val, nil
}

// SolvePart2 ..
func SolvePart2(tokens [][]string) (int, error) {
	sum := 0
	for _, v := range tokens {
		res, err := ModifiedShuntingYard2(v)
		if err != nil {
			return 0, err
		}
		sum += res
	}
	return sum, nil
}
