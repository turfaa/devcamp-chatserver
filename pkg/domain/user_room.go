package domain

import (
	"fmt"
)

// UserRoom type
type UserRoom struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoomID   string `json:"roomID"`
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

// GetRoomMembers domain
func (userRoomDomain UserRoomDomain) GetRoomMembers(roomID string) ([]string, error) {
	return userRoomDomain.resource.GetRoomMembers(roomID)
}

// GetUserRooms domain
func (userRoomDomain UserRoomDomain) GetUserRooms(username string) ([]string, error) {
	return userRoomDomain.resource.GetUserRooms(username)
}

// CreateUserRoom domain
func (userRoomDomain UserRoomDomain) CreateUserRoom(userRoom *UserRoom) error {
	return userRoomDomain.resource.CreateUserRoom(userRoom)
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
	joinedRooms, err := userRoomResource.GetUserRooms(userRoom.Username)

	if err != nil {
		return err
	}

	for _, joinedRoomID := range joinedRooms {
		if joinedRoomID == userRoom.RoomID {
			return fmt.Errorf("Already joined")
		}
	}

	userRoom.ID = 1
	return nil
}
