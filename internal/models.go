package internal

type Room struct {
	Name string
	X    int
	Y    int
}

type Connection struct {
	NameA string
	NameB string
}

type Colony struct {
	Ants        int
	Rooms       map[string]Room
	Connections []Connection
}
