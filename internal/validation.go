package internal

import (
	"fmt"
	"strings"
)

func Validate(colony Colony) error {

	if colony.Ants <= 0 {
		return fmt.Errorf("ERROR: Invalid Number of Ants")
	}
	if colony.StartCount != 1 {
		return fmt.Errorf("ERROR: You are allowed to have only one starting point")
	}
	if colony.EndCount != 1 {
		return fmt.Errorf("ERROR: You are allowed to have only one ending point")
	}
	for _, room := range colony.Rooms {
		if strings.HasPrefix(room.Name, "L") || strings.HasPrefix(room.Name, "#") {
			return fmt.Errorf("ERROR: Invalid Room Name, you cant have a room name starting with L or #")
		}

	}
	return nil
}
