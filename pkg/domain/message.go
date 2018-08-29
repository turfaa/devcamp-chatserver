package domain

// Message type
type Message struct {
	id          int `json:"id"`
	timestamp   int64 `json:"timestamp"`
	sender      string `json:"sender"`
	messageType string `json:"message_type"`
	receiver    string `json:"receiver"`
	text        string `json:"text"`
}

// MessageResourceItf interface of room resource
type MessageResourceItf interface {
	GetPrivateMessages(string) ([]Message, error)
	GetRoomMessages(string) ([]Message, error)
	CreateMessage(*Message) error
}

// MessageResourceFake fake
type MessageResourceFake struct {
}

// MessageDomain room domain
type MessageDomain struct {
	resource MessageResourceItf
}

// InitMessageDomain init room domain
func InitMessageDomain(rsc MessageResourceItf) MessageDomain {
	return MessageDomain{
		resource: rsc,
	}
}

// GetPrivateMessages fake
func (messageResource MessageResourceFake) GetPrivateMessages(username string) ([]Message, error) {
	messages := []Message{
		Message{1, 1, "turfaa", "private", "kamu", "halo kamu"},
		Message{2, 2, "kamu", "private", "turfaa", "halo juga"},
		Message{3, 3, "turfaa", "room", "room1", "hai semua"}
	}

	var privateMessages []Message

	for _, message := range messages {
		if message.messageType == "private" && (message.sender == username || message.receiver == username) {
			privateMessages = append(privateMessages, message)
		}
	}

	return privateMessages, nil
}

// GetRoomMessages fake
func (messageResource MessageResourceFake) GetRoomMessages(roomID string) ([]Message, error) {
	messages := []Message{
		Message{1, 1, "turfaa", "private", "kamu", "halo kamu"},
		Message{2, 2, "kamu", "private", "turfaa", "halo juga"},
		Message{3, 3, "turfaa", "room", "room1", "hai semua"}
	}

	var roomMessages []Message

	for _, message := range messages {
		if message.messageType == "room" && (message.receiver == roomID) {
			roomMessages = append(roomMessages, message)
		}
	}

	return roomMessages, nil
}

// CreateRoomMessage fake
func (messageResource MessageResourceFake) CreateRoomMessage(message *Message) error {
	return nil
}