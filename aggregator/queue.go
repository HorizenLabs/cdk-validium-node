package aggregator

import (
	"errors"
	"fmt"
)

type FinalProofsQueue struct {
	FinalProofs []finalProofElement
}

func (q *FinalProofsQueue) Enqueue(elem finalProofElement) {
	q.FinalProofs = append(q.FinalProofs, elem)
}

func (q *FinalProofsQueue) Dequeue() (finalProofElement, error) {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return finalProofElement{}, errors.New("empty queue")
	}
	element := q.FinalProofs[0]
	// if q.GetLength() == 1 {
	// 	q.FinalProofs = nil
	// 	return element
	// }
	q.FinalProofs = q.FinalProofs[1:]
	return element, nil // Slice off the element once it is dequeued.
}

func (q *FinalProofsQueue) GetLength() int {
	return len(q.FinalProofs)
}

func (q *FinalProofsQueue) IsEmpty() bool {
	return q.GetLength() == 0
}

func (q *FinalProofsQueue) Peek() (finalProofElement, error) {
	if q.IsEmpty() {
		return finalProofElement{}, errors.New("empty queue")
	}
	return q.FinalProofs[0], nil
}
