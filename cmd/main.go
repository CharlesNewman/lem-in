package main

import (
	"fmt"
	"lem-in/internal"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid data format")
		return
	}

	input := os.Args[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	colony, err := internal.ParseData(string(data))
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	fmt.Println(colony)
}
