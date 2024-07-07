package repository

import (
	"errors"
	"fmt"
	"top-selection-test/internal/model"
)

const (
	dateLayout = "2007/01/02"
)

var (
	ErrRoomNotAvailable = errors.New("room is not available")
	ErrRoomDoesntExist  = errors.New("room doesn't exist")
)

func formatOrderError(err error, o model.Order, withDate bool) error {
	if withDate {
		return fmt.Errorf(
			"%w: room ID %s from %s to %s",
			err,
			fmt.Sprintf("%s_%s", o.HotelID, o.RoomID),
			o.From.Format(dateLayout),
			o.To.Format(dateLayout),
		)
	}

	return fmt.Errorf(
		"%w: room ID %s",
		err,
		fmt.Sprintf("%s_%s", o.HotelID, o.RoomID),
	)
}
