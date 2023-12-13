package main

import "net"

type Observer struct {
	queue   Queue
	theater Theater
	left    bool
	next    net.Conn
}

func (observer *Observer) NotifyExit() {
	observer.next = observer.queue.Dequeue()
	observer.left = true
}

func (observer *Observer) GetNext() net.Conn {
	nextInLine := observer.next
	observer.next = nil
	observer.left = false
	return nextInLine
}
