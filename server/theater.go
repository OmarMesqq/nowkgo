package main

type Theater struct {
	peopleInRoom       int
	maxTheaterCapacity int
}

func (theater *Theater) IsThereRoom() bool {
	return theater.peopleInRoom < theater.maxTheaterCapacity
}

func (theater *Theater) Leave(observer *Observer) {
	theater.peopleInRoom -= 1
	observer.NotifyExit()
}
