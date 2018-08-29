package domain

// UserRoom type
type UserRoom struct {
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
