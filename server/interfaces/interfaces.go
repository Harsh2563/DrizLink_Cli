package interfaces

import (
	"encoding/json"
	"net"
	"sync"
	"time"
)

type Server struct {
	Address     string
	Connections map[string]*User
	IpAddresses map[string]*User
	Messages    chan Message
	Mutex       sync.Mutex
}

// MessageType defines the type of message being sent
type MessageType string

const (
	MessageTypeChat           MessageType = "chat"
	MessageTypeSystem         MessageType = "system"
	MessageTypeCommand        MessageType = "command"
	MessageTypeFileTransfer   MessageType = "file_transfer"
	MessageTypeFolderTransfer MessageType = "folder_transfer"
	MessageTypeStatus         MessageType = "status"
	MessageTypePing           MessageType = "ping"
	MessageTypePong           MessageType = "pong"
	MessageTypeError          MessageType = "error"
	MessageTypeUserList       MessageType = "user_list"
	MessageTypeLookup         MessageType = "lookup"
	MessageTypeDownload       MessageType = "download"
)

// Message represents a structured message with JSON support
type Message struct {
	Type           MessageType            `json:"type"`
	SenderId       string                 `json:"sender_id"`
	SenderUsername string                 `json:"sender_username"`
	Content        string                 `json:"content"`
	Timestamp      string                 `json:"timestamp"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// NewMessage creates a new message with current timestamp
func NewMessage(msgType MessageType, senderId, senderUsername, content string) *Message {
	return &Message{
		Type:           msgType,
		SenderId:       senderId,
		SenderUsername: senderUsername,
		Content:        content,
		Timestamp:      time.Now().Format(time.RFC3339),
		Metadata:       make(map[string]interface{}),
	}
}

// ToJSON converts the message to JSON bytes
func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// FromJSON creates a message from JSON bytes
func FromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}

// AddMetadata adds metadata to the message
func (m *Message) AddMetadata(key string, value interface{}) {
	if m.Metadata == nil {
		m.Metadata = make(map[string]interface{})
	}
	m.Metadata[key] = value
}

type User struct {
	UserId        string
	Username      string
	StoreFilePath string
	Conn          net.Conn
	IsOnline      bool
	IpAddress     string
}
