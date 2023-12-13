package main

type Theater struct {
	peopleInRoom       int
	maxTheaterCapacity int
}

func (room *Theater) IsThereRoom() bool {
	return room.peopleInRoom < room.maxTheaterCapacity
}
