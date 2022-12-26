package game

// the types that the player can have
// Dog
type Dog struct {
	Name  string
	speed int
	x     int
}


func NewDog(name string) *Dog {
	return &Dog[Name: name, Speed: 2 , x:0]
}

// Cat
type Cat struct {
	Name  string
	speed int
	x     int
}

func NewCat(name string) *Cat {
	return &Cat[Name: name, Speed: 1 , x:0]
}

// Owl
type Owl struct {
	Name  string
	speed int
	x     int
}

func NewOwl(name string) *Owl {
	return &Owl[Name: name, Speed: 3 , x:0]
}