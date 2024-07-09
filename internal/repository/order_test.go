//go:build unit

package repository

import (
	"errors"
	"slices"
	"testing"
	"time"
	"top-selection-test/internal/model"
)

const (
	expectedOrderCount = 4
)

func TestOrders_Add(t *testing.T) {
	orders := NewOrders()
	startDate := time.Date(2024, 7, 1, 0, 0, 0, 0, &time.Location{})

	// Test must be run exact this order
	tests := []struct {
		name  string
		order model.Order
		err   error
	}{
		{
			name: "insert not existing room",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "noroom",
				UserEmail: "test_user@mail.com",
				From:      startDate,
				To:        startDate.AddDate(0, 0, 3),
			},
			err: ErrRoomDoesntExist,
		},
		{
			name: "insert in empty slice",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate,
				To:        startDate.AddDate(0, 0, 3),
			},
			err: nil,
		},
		{
			name: "insert exact same dates",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate,
				To:        startDate.AddDate(0, 0, 3),
			},
			err: ErrRoomNotAvailable,
		},
		{
			name: "intersect left border",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate.AddDate(0, 0, 2),
				To:        startDate.AddDate(0, 0, 6),
			},
			err: ErrRoomNotAvailable,
		},
		{
			name: "intersect right border",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate.AddDate(0, 0, -1),
				To:        startDate.AddDate(0, 0, 1),
			},
			err: ErrRoomNotAvailable,
		},
		{
			name: "no intersection insert left",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate.AddDate(0, 0, -10),
				To:        startDate.AddDate(0, 0, -5),
			},
			err: nil,
		},
		{
			name: "no intersection insert with the same dates",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate.AddDate(0, 0, -5),
				To:        startDate,
			},
			err: nil,
		},
		{
			name: "no intersection insert right",
			order: model.Order{
				HotelID:   "redisson",
				RoomID:    "blackhole",
				UserEmail: "test_user@mail.com",
				From:      startDate.AddDate(0, 0, 3),
				To:        startDate.AddDate(0, 0, 7),
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := orders.Create(tt.order); !errors.Is(err, tt.err) {
				t.Fatalf("%s failed: want '%s' got '%s'", tt.name, tt.err, err)
			}
		})
	}

	resOrders := orders[predefinedRooms["redisson_blackhole"]]
	if len(resOrders) != expectedOrderCount {
		t.Fatalf("orders count mismatch: want %d got %d", expectedOrderCount, len(resOrders))
	}

	if !slices.IsSortedFunc(resOrders, func(a, b model.Order) int {
		return a.From.Compare(b.From)
	}) {
		t.Fatalf("order's slice is not sorted")
	}
}
