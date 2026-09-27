package internal

import "fmt"

func Validate(Colony Colony) error {

	if Colony.Ants <= 0 {
		return fmt.Errorf("ERROR: Invalid Number of Ants")
	}

	if Colony.StartCount != 1 {
		return fmt.Errorf("ERROR: You are allowed to have only one starting point")
	}
	if Colony.EndCount != 1 {
		return fmt.Errorf("ERROR: You are allowed to have only one ending point")
	}
	return nil
}
