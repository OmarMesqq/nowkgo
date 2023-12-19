package main

type Theater struct {
	peopleInRoom       int
	maxTheaterCapacity int
	ready              chan bool
	queue              *Queue
}

func (theater *Theater) IsThereRoom() bool {
	return theater.peopleInRoom < theater.maxTheaterCapacity
}

func (theater *Theater) Leave() {
	theater.peopleInRoom -= 1
	if theater.queue.IsEmpty() {
		return
	}
	theater.ready <- true
}

func (theater *Theater) Enter() {
	theater.peopleInRoom += 1
}
