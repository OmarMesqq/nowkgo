package main

type Theater struct {
	peopleInRoom       int
	maxTheaterCapacity int
}

func (room *Theater) IsThereRoom() bool {
	return room.peopleInRoom < room.maxTheaterCapacity
}

func (room *Theater) IsEmpty() bool {
	return room.peopleInRoom == 0
}
