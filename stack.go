package gostructures

import "errors"

/*
	TODO:
		- Change inner array with linkedin list
		- Support max item
*/
type Stack struct {
	innerArray []interface{}
}

func (s *Stack) Push(data interface{}) error {
	s.innerArray = append(s.innerArray, data)
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	popItem := s.innerArray[len(s.innerArray)-1]
	s.innerArray = s.innerArray[:len(s.innerArray)-1]
	return popItem, nil
}

func (s *Stack) PopWithoutErr() interface{} {
	popItem, _ := s.Pop()
	return popItem
}

func (s *Stack) IsEmpty() bool {
	return len(s.innerArray) <= 0
}

func (s *Stack) Length() int {
	return len(s.innerArray)
}
