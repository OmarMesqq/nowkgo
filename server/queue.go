package main

import (
	"net"
)

type Queue struct {
	people []net.Addr
}

func (queue Queue) GetPeopleInLine() int {
	return len(queue.people)
}

func (queue *Queue) Enqueue(clientAddress net.Addr) {
	queue.people = append(queue.people, clientAddress)
}

func (queue *Queue) Dequeue() net.Addr {
	if len(queue.people) == 0 {
		return nil
	}

	nextInLine := queue.people[0]
	queue.people = queue.people[:1]
	return nextInLine
}
