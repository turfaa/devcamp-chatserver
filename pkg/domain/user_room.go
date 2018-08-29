package domain

import (
	"fmt"
)

// UserRoom type
type UserRoom struct {
	id       int    `json:"id"`
	username string `json:"username"`
	roomID   string `json:"roomID"`
}

// UserRoomResourceItf interface of room resource
type UserRoomResourceItf interface {
	GetRoomMembers(string) ([]string, error)
	GetUserRooms(string) ([]string, error)
	CreateUserRoom(*UserRoom) error
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
func (userRoomResource UserRoomResourceFake) CreateUserRoom(userRoom *UserRoom) error {
	joinedRooms, err := userRoomResource.GetUserRooms(userRoom.username)

	if err != nil {
		return err
	}

	for _, joinedRoomID := range joinedRooms {
		if joinedRoomID == userRoom.roomID {
			return fmt.Errorf("Already joined")
		}
	}

	userRoom.id = 1
	return nil
}
