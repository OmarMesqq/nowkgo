package main

import (
	"net"
)

type Queue struct {
	people []net.Conn
}

func (queue *Queue) Enqueue(clientAddress net.Conn) {
	queue.people = append(queue.people, clientAddress)
}

func (queue *Queue) Dequeue() net.Conn {
	if len(queue.people) == 0 {
		return nil
	}

	nextInLine := queue.people[0]
	queue.people = queue.people[:1]
	return nextInLine
}
