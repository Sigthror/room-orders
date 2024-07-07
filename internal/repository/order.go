package repository

import (
	"fmt"
	"slices"
	"top-selection-test/internal/model"
)

type Orders map[model.Room][]model.Order

func NewOrders() Orders {
	o := make(Orders)
	for _, room := range predefinedRooms {
		o[room] = make([]model.Order, 0)
	}

	return o
}

func (o Orders) Add(order model.Order) error {
	// TODO Add mutex
	room := rooms.GetByName(order.HotelID, order.RoomID)
	if room == nil {
		return formatOrderError(ErrRoomDoesntExist, order, false)
	}

	roomOrders := o[*room]
	i, inOrders := slices.BinarySearchFunc(roomOrders, order, func(a, b model.Order) int {
		// TODO Analyze this if
		if a.From == b.From && a.To == b.To {
			return 0
		}

		return a.From.Compare(b.From)
	})

	if inOrders {
		return formatOrderError(ErrRoomNotAvailable, order, true)
	}

	// After binary search we have invariant:
	// roomOrders[i-i].From < order.From < roomOrders.[i].From

	// Check that roomOrders[i-1].To < order.From
	if i >= 1 && roomOrders[i-1].To.After(order.From) {
		return formatOrderError(ErrRoomNotAvailable, order, true)
	}

	// Check that roomOrders[i].From > order.To
	if i < len(roomOrders) && roomOrders[i].From.Before(order.To) {
		return formatOrderError(ErrRoomNotAvailable, order, true)
	}

	o[*room] = slices.Insert(roomOrders, i, order)

	fmt.Printf("room boked from %s to %s\n", order.From.Format("2006/01/02"), order.To.Format("2006/01/02"))

	return nil
}
