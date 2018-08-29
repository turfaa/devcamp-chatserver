package domain

import (
	"chatserver/pkg/lib/utils"

	"github.com/tokopedia/tdk/go/app/resource"
)

// Room type
type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RoomResourceItf interface of room resource
type RoomResourceItf interface {
	GetAllRooms() ([]Room, error)
	FindRoom(string) (Room, bool, error)
	CreateRoom(*Room) error
}

// RoomResourceDB room resource from db
type RoomResourceDB struct {
	DB resource.SQLDB
}

// RoomResourceFake fake data for room resource
type RoomResourceFake struct {
}

// RoomDomain room domain
type RoomDomain struct {
	resource RoomResourceItf
}

// InitRoomDomain init room domain
func InitRoomDomain(rsc RoomResourceItf) RoomDomain {
	return RoomDomain{
		resource: rsc,
	}
}

// GetAllRooms domain
func (roomDomain RoomDomain) GetAllRooms() ([]Room, error) {
	return roomDomain.resource.GetAllRooms()
}

// FindRoom domain
func (roomDomain RoomDomain) FindRoom(roomID string) (Room, bool, error) {
	return roomDomain.resource.FindRoom(roomID)
}

// CreateRoom domain
func (roomDomain RoomDomain) CreateRoom(room *Room) error {
	return roomDomain.resource.CreateRoom(room)
}

// GetAllRooms fake
func (roomResource RoomResourceFake) GetAllRooms() ([]Room, error) {
	return []Room{
		Room{"room1", "Room 1"},
		Room{"room2", "Room 2"},
		Room{"room3", "Room 3"},
	}, nil
}

// FindRoom fake
func (roomResource RoomResourceFake) FindRoom(id string) (Room, bool, error) {
	rooms := []Room{
		Room{"room1", "Room 1"},
		Room{"room2", "Room 2"},
		Room{"room3", "Room 3"},
	}

	for _, room := range rooms {
		if id == room.ID {
			return room, true, nil
		}
	}

	return Room{}, false, nil
}

// CreateRoom fake
func (roomResource RoomResourceFake) CreateRoom(room *Room) error {
	id := utils.GenerateRandomString(32)

	_, found, err := roomResource.FindRoom(id)

	if err != nil {
		return err
	}

	for found {
		id = utils.GenerateRandomString(32)

		if _, found, err = roomResource.FindRoom(id); err != nil {
			return err
		}
	}

	room.ID = id
	return nil
}
