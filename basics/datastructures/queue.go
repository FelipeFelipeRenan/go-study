package datastructures

import "errors"

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(el any)  {
	q.elements = append(q.elements, el)
}

func (q *Queue) Dequeue() (el any, err error){
	if q.IsEmpty(){
		err = errors.New("empty queue")
		return
	}
	el = q.elements[0]
	q.elements = q.elements[1:]
	return
}

func (q *Queue) Peek()(el any, err error)  {
	if q.IsEmpty(){
		err = errors.New("empty queue")
		return
	}
	el = q.elements[0]
	return
}

func (q *Queue) IsEmpty() bool{
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}