package main

type Theater struct {
	peopleInRoom       int
	maxTheaterCapacity int
	observedQueue      Queue
}

func (theater *Theater) IsThereRoom() bool {
	return theater.peopleInRoom < theater.maxTheaterCapacity
}

func (theater *Theater) Leave(queue *Queue) {
	theater.peopleInRoom -= 1

}
