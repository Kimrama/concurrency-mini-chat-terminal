package models

import (
	"sync"
)

type Room struct {
	Name    string
	Message chan string
	Member  map[string]bool
}
type RoomManager struct {
	Rooms map[string]*Room // map room name -> room instance
	mu    sync.RWMutex
}

var Manager = &RoomManager{
	Rooms: make(map[string]*Room),
}

func (rm *RoomManager) CreateRoom(roomName string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, exist := rm.Rooms[roomName]; !exist {
		rm.Rooms[roomName] = &Room{
			Name:    roomName,
			Message: make(chan string, 100),
			Member:  make(map[string]bool),
		}
	}
}
func (rm *RoomManager) JoinRoom(roomName, userName string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if room, ok := rm.Rooms[roomName]; ok {
		room.Member[userName] = true
	}
}

func (rm *RoomManager) Broadcast(roomName, msg string) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	if room, ok := rm.Rooms[roomName]; ok {
		room.Message <- msg
	}
}

func (rm *RoomManager) GetRoom(roomName string) *Room {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	return rm.Rooms[roomName]
}
