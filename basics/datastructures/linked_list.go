package datastructures

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Element[T constraints.Ordered] struct {
	value T
	next  *Element[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Element[T]
	size int
}

func (l *LinkedList[T]) Add(el *Element[T]) {
	if l.head == nil {
		l.head = el
	} else {
		el.next = l.head
		l.head = el
	}
	l.size++
}

func (l *LinkedList[T]) Insert(el *Element[T], marker T) error {
	for current := l.head; current.next != nil; current = current.next {
		if current.value == marker {
			el.next = current.next
			current.next = el
			l.size++
			return nil
		}
	}
	return errors.New("element not found")
}

func (l *LinkedList[T]) Delete(el *Element[T]) error {
	prev := l.head
	current := l.head
	for current != nil {
		if current.value == el.value {
			if current == l.head {
				l.head = current.next
			} else {
				prev.next = current.next
			}
			l.size--
			return nil
		}
		prev = current
		current = current.next
	}
	return errors.New("element not found")
}

func (l *LinkedList[T]) Find(value T) (el *Element[T], err error) {
	for current := l.head; current.next != nil; current = current.next {
		if current.value == value {
			el = current
			break
		}
	}
	if el == nil {
		err = errors.New("element not found")
	}
	return
}

func (l *LinkedList[T])List() (list []*Element[T])  {
	if l.head == nil{
		return []*Element[T]{}
	}
	for current := l.head; current != nil; current = current.next {
		list = append(list, current)
	}
	return
}

func (l *LinkedList[T]) IsEmpty() bool{
	return l.size == 0
}

func (l *LinkedList[T]) Size() int{
	return l.size
}
