package main

import (
	"fmt"
	"lem-in/internal"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error Reading The file")
		return
	}

	datasplit := strings.Split(string(data), "\n")

	antsFound := false

	rooms := make(map[string]internal.Room)

	for i := 0; i < len(datasplit); i++ {
		indata := datasplit[i]
		if indata == "" {
			continue
		}
		if strings.HasPrefix(indata, "#") && indata != "##start" && indata != "##end" {
			continue
		}
		if !antsFound {
			ants, err := strconv.Atoi(indata)
			if err != nil {
				fmt.Println("ERROR: Invalid data format")
				return
			}
			fmt.Println("Ants", ants)
			antsFound = true
			continue
		}

		parts := strings.Fields(indata)
		if len(parts) == 3 {
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return
			}
			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return
			}
			room := internal.Room{
				Name: parts[0],
				X:    x,
				Y:    y,
			}
			fmt.Printf("This is the room: %s %d %d\n", room.Name, room.X, room.Y)
			rooms[room.Name] = room
		}
	}
}
