package main

import "container/list"
import "fmt"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func isValid(s string) bool {
	parentheses := map[string]string{"(": ")", "[": "]", "{": "}"}
	// 去掉空格
	stack := NewStack()
	for _, c_rune := range s {
		c := string(c_rune)
		if c == "(" || c == "[" || c == "{" {
			stack.Push(c)
			continue
		}
		if c == ")" || c == "]" || c == "}" {
			a := stack.Pop()
			if a == nil {
				return false
			}
			p := a.(string)
			if parentheses[p] == c {
				continue
			} else {
				return false
			}
		}
	}
	if stack.Empty() {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("()[]{"))
}
