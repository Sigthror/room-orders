package repository

import (
	"errors"
	"fmt"
	"sync"
	"top-selection-test/internal/model"
)

var (
	predefinedRooms = roomStorage{
		"redisson_blackhole": model.Room{
			HotelID: "redisson",
			RoomID:  "blackhole",
		},
		"redisson_morningstart": model.Room{
			HotelID: "redisson",
			RoomID:  "morningstart",
		},
		"ritz_italy": model.Room{
			HotelID: "ritz",
			RoomID:  "italy",
		},
		"ritz_spain": model.Room{
			HotelID: "ritz",
			RoomID:  "spain",
		},
	}

	rooms = NewRooms()
)

type roomStorage map[string]model.Room

type Rooms struct {
	mu      sync.RWMutex
	storage roomStorage
}

func NewRooms() *Rooms {
	return &Rooms{
		storage: predefinedRooms,
	}
}

func (r *Rooms) Add(m model.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sKey := fmt.Sprintf("%s_%s", m.HotelID, m.RoomID)
	if _, ok := r.storage[sKey]; ok {
		return errors.New("room already exists")
	}
	r.storage[sKey] = m

	return nil
}

func (r *Rooms) GetByName(hotel, roomName string) *model.Room {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sKey := fmt.Sprintf("%s_%s", hotel, roomName)
	room, ok := r.storage[sKey]
	if !ok {
		return nil
	}

	return &room
}
