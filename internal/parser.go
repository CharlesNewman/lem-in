package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseData(data string) (Colony, error) {
	datasplit := strings.Split(data, "\n")

	antsFound := false

	colony := Colony{
		Rooms:       make(map[string]Room),
		Connections: []Connection{},
	}

	nextRoomIsStart := false
	nextRoomIsEnd := false

	for i := 0; i < len(datasplit); i++ {
		indata := datasplit[i]

		if indata == "" {
			continue
		}

		if strings.HasPrefix(indata, "#") && indata != "##start" && indata != "##end" {
			continue
		}
		if indata == "##start" {
			nextRoomIsStart = true
			continue
		}
		if indata == "##end" {
			nextRoomIsEnd = true
			continue
		}

		if !antsFound {
			ants, err := strconv.Atoi(indata)
			if err != nil {
				return Colony{}, fmt.Errorf("invalid ants")
			}

			colony.Ants = ants
			antsFound = true
			continue
		}

		parts := strings.Fields(indata)

		if len(parts) == 3 {
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return Colony{}, fmt.Errorf("invalid x coordinate")
			}

			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return Colony{}, fmt.Errorf("invalid y coordinate")
			}

			room := Room{
				Name: parts[0],
				X:    x,
				Y:    y,
			}
			if nextRoomIsStart {
				colony.Start = room.Name
				colony.StartCount++
				nextRoomIsStart = false
			}

			if nextRoomIsEnd {
				colony.End = room.Name
				nextRoomIsEnd = false
				colony.EndCount++
			}

			fmt.Printf("This is the room: %s %d %d\n", room.Name, room.X, room.Y)

			colony.Rooms[room.Name] = room
			continue
		}

		if strings.Contains(indata, "-") {
			connectionParts := strings.Split(indata, "-")

			if len(connectionParts) != 2 {
				return Colony{}, fmt.Errorf("invalid connection")
			}

			connection := Connection{
				NameA: connectionParts[0],
				NameB: connectionParts[1],
			}

			colony.Connections = append(colony.Connections, connection)

			fmt.Println(connection)
		}
	}

	return colony, nil
}
