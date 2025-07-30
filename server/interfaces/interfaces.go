package interfaces

import (
	"encoding/json"
	"net"
	"sync"
)

type Server struct {
	Address     string
	Connections map[string]*User
	IpAddresses map[string]*User
	Messages    chan Message
	Mutex       sync.Mutex
}

type Message struct {
	SenderId       string `json:"senderId"`
	SenderUsername string `json:"senderUsername"`
	Content        string `json:"content"`
	Timestamp      string `json:"timestamp"`
	Type           string `json:"type"` // "chat", "system", "command"
}

// ToJSON converts a Message to JSON string
func (m *Message) ToJSON() (string, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON creates a Message from JSON string
func MessageFromJSON(jsonStr string) (*Message, error) {
	var msg Message
	err := json.Unmarshal([]byte(jsonStr), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

type User struct {
	UserId        string
	Username      string
	StoreFilePath string
	Conn          net.Conn
	IsOnline      bool
	IpAddress     string
}
