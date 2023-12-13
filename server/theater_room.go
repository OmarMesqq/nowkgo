package main

type TheaterRoom struct {
	peopleInRoom       int
	maxTheaterCapacity int
}

func (room TheaterRoom) IsThereRoom() bool {
	return room.peopleInRoom < room.maxTheaterCapacity
}

func (room TheaterRoom) IsEmpty() bool {
	return room.peopleInRoom == 0
}
