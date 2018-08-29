package domain

import "github.com/tokopedia/tdk/go/log"

// Message type
type Message struct {
	ID          int    `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	Sender      string `json:"sender"`
	MessageType string `json:"message_type"`
	Receiver    string `json:"receiver"`
	Text        string `json:"text"`
}

// MessageSorter sorter
type MessageSorter []Message

func (a MessageSorter) Len() int           { return len(a) }
func (a MessageSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MessageSorter) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }

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

// GetPrivateMessages domain
func (messageDomain MessageDomain) GetPrivateMessages(username string) ([]Message, error) {
	return messageDomain.resource.GetPrivateMessages(username)
}

// GetRoomMessages domain
func (messageDomain MessageDomain) GetRoomMessages(roomID string) ([]Message, error) {
	return messageDomain.resource.GetRoomMessages(roomID)
}

// CreateMessage domain
func (messageDomain MessageDomain) CreateMessage(message *Message) error {
	return messageDomain.resource.CreateMessage(message)
}

// GetPrivateMessages fake
func (messageResource MessageResourceFake) GetPrivateMessages(username string) ([]Message, error) {
	messages := []Message{
		Message{1, 1, "turfaa", "private", "kamu", "halo kamu"},
		Message{2, 2, "kamu", "private", "turfaa", "halo juga"},
		Message{3, 3, "turfaa", "room", "room1", "hai semua"},
	}

	var privateMessages []Message

	for _, message := range messages {
		log.Info(message.Receiver, message.Sender, message.MessageType, username)
		if message.MessageType == "private" && (message.Sender == username || message.Receiver == username) {
			privateMessages = append(privateMessages, message)
		}
	}
	log.Info(privateMessages)
	return privateMessages, nil
}

// GetRoomMessages fake
func (messageResource MessageResourceFake) GetRoomMessages(roomID string) ([]Message, error) {
	messages := []Message{
		Message{1, 1, "turfaa", "private", "kamu", "halo kamu"},
		Message{2, 2, "kamu", "private", "turfaa", "halo juga"},
		Message{3, 3, "turfaa", "room", "room1", "hai semua"},
	}

	var roomMessages []Message

	for _, message := range messages {
		if message.MessageType == "room" && (message.Receiver == roomID) {
			roomMessages = append(roomMessages, message)
		}
	}

	return roomMessages, nil
}

// CreateMessage fake
func (messageResource MessageResourceFake) CreateMessage(message *Message) error {
	return nil
}
