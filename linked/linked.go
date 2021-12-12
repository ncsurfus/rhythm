package linked

import "errors"

var ElementNotFound = errors.New("linked element not found.")

type Element[T any] struct {
	Value T
	Next  *Element[T]
}

func (el *Element[T]) InsertAfter(value T) {
	newElement := &Element[T]{
		Value: value,
		Next:  el.Next,
	}
	el.Next = newElement
}

func (el *Element[T]) Last() *Element[T] {
	current := el
	for current.Next != nil {
		current = current.Next
	}
	return current
}

func FromSlice[T any](values []T) *Element[T] {
	if len(values) == 0 {
		return nil
	}

	root := &Element[T]{Value: values[0]}
	current := root
	for _, value := range values[1:] {
		element := &Element[T]{Value: value}
		current.Next, current = element, element
	}
	return root
}

func FindElementByValue[T comparable](element *Element[T], value T) (*Element[T], error) {
	for {
		if element == nil {
			return nil, ElementNotFound
		} else if element.Value == value {
			return element, nil
		}
		element = element.Next
	}
}
