package usecase

import (
	"chatserver/pkg/domain"
	"chatserver/pkg/lib/config"
	"fmt"
)

// ChatUsecase chat use case
type ChatUsecase struct {
	config         config.Config
	roomDomain     domain.RoomDomain
	userRoomDomain domain.UserRoomDomain
	messageDomain  domain.MessageDomain
}

// InitChatUsecase init
func InitChatUsecase(
	config config.Config,
	roomDomain domain.RoomDomain,
	userRoomDomain domain.UserRoomDomain,
	messageDomain domain.MessageDomain,
) *ChatUsecase {

	return &ChatUsecase{
		config,
		roomDomain,
		userRoomDomain,
		messageDomain,
	}
}

// GetMessages get messages
func (chatUsecase *ChatUsecase) GetMessages(username string) ([]domain.Message, error) {
	var messages []domain.Message

	if privateMessages, err := chatUsecase.messageDomain.GetPrivateMessages(username); err == nil {
		messages = append(messages, privateMessages...)
	} else {
		return []domain.Message{}, err
	}

	if roomIDs, err := chatUsecase.userRoomDomain.GetUserRooms(username); err != nil {
		for _, roomID := range roomIDs {
			if roomMessages, err2 := chatUsecase.messageDomain.GetRoomMessages(roomID); err2 == nil {
				messages = append(messages, roomMessages...)
			} else {
				return []domain.Message{}, err
			}
		}
	}

	// sort.Sort(domain.MessageSorter(messages))
	return messages, nil
}

// SendMessage send message
func (chatUsecase *ChatUsecase) SendMessage(message *domain.Message) error {
	if message.MessageType == "room" {
		if _, found, err := chatUsecase.roomDomain.FindRoom(message.Receiver); !found || err != nil {
			if err != nil {
				return err
			}

			return fmt.Errorf("Room not found")
		}
	}

	if err := chatUsecase.messageDomain.CreateMessage(message); err != nil {
		return err
	}

	return nil
}

// NewRoom new chat room
func (chatUsecase *ChatUsecase) NewRoom(room *domain.Room) error {
	if err := chatUsecase.roomDomain.CreateRoom(room); err != nil {
		return err
	}

	return nil
}

// JoinRoom join chat room
func (chatUsecase *ChatUsecase) JoinRoom(userRoom *domain.UserRoom) error {
	if _, found, err := chatUsecase.roomDomain.FindRoom(userRoom.RoomID); !found || err != nil {
		if err != nil {
			return err
		}

		return fmt.Errorf("Room not found")
	}

	if err := chatUsecase.userRoomDomain.CreateUserRoom(userRoom); err != nil {
		return err
	}

	return nil
}

// GetAllRooms get all rooms
func (chatUsecase *ChatUsecase) GetAllRooms() ([]domain.Room, error) {
	return chatUsecase.roomDomain.GetAllRooms()
}

// GetUserRooms get user rooms
func (chatUsecase *ChatUsecase) GetUserRooms(username string) ([]domain.Room, error) {
	if roomIDs, err := chatUsecase.userRoomDomain.GetUserRooms(username); err != nil {
		return []domain.Room{}, err
	} else {
		var rooms []domain.Room

		for _, roomID := range roomIDs {
			if room, found, err := chatUsecase.roomDomain.FindRoom(roomID); !found || err != nil {
				if err != nil {
					return []domain.Room{}, err
				} else {
					return []domain.Room{}, err
				}
			} else {
				rooms = append(rooms, room)
			}
		}

		return rooms, nil
	}
}
