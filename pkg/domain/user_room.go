package domain

import (
	"fmt"
)

// UserRoom type
type UserRoom struct {
	id       int
	username string
	roomID   string
}

// UserRoomResourceItf interface of room resource
type UserRoomResourceItf interface {
	GetRoomMembers(string) ([]string, error)
	GetUserRooms(string) ([]string, error)
	CreateUserRoom(string) (string, string, error)
}

// UserRoomResourceFake fake
type UserRoomResourceFake struct {
}

// UserRoomDomain room domain
type UserRoomDomain struct {
	resource UserRoomResourceItf
}

// InitUserRoomDomain init room domain
func InitUserRoomDomain(rsc UserRoomResourceItf) UserRoomDomain {
	return UserRoomDomain{
		resource: rsc,
	}
}

// GetRoomMembers fake
func (userRoomResource UserRoomResourceFake) GetRoomMembers(roomID string) ([]string, error) {
	if roomID == "room1" {
		return []string{"turfa"}, nil
	}

	return []string{}, nil
}

// GetUserRooms fake
func (userRoomResource UserRoomResourceFake) GetUserRooms(username string) ([]string, error) {
	if username == "turfa" {
		return []string{"room1"}, nil
	}

	return []string{}, nil
}

// CreateUserRoom fake
func (userRoomResource UserRoomResourceFake) CreateUserRoom(username string, roomID string) (UserRoom, error) {
	joinedRooms, err := userRoomResource.GetUserRooms(username)

	if err != nil {
		return UserRoom{}, err
	}

	for _, joinedRoomId := range joinedRooms {
		if joinedRoomId == roomID {
			return UserRoom{}, fmt.Errorf("Already joined")
		}
	}

	return UserRoom{1, username, roomID}, nil
}
