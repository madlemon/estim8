package model

import (
	"strconv"
	"sync"
)

type GlobalInMemoryState struct {
	Rooms map[string]*Room
	mutex *sync.RWMutex
}

func InitGlobalState() GlobalInMemoryState {
	state := GlobalInMemoryState{}
	state.Rooms = make(map[string]*Room)
	state.mutex = &sync.RWMutex{}
	return state
}

func (global GlobalInMemoryState) GetRoom(roomId string) *Room {
	global.mutex.RLock()
	defer global.mutex.RUnlock()
	return global.Rooms[roomId]
}

var Global = InitGlobalState()

type Estimate int

type User string

func (e Estimate) String() string {
	return strconv.Itoa(int(e))
}
