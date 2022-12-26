package game

// the types that the player can have
// Dog
type Dog struct {
	Name  string
	speed int
	x     int
}


func NewDog(name string) *Dog {
	return &Dog[name, Speed: 2 , x:0]
}

func(c *Dog) Position() int {
	return c.x
	}
	
	func(c *Dog) Name() string {
		return c.name 
	}
	
	func(c *Dog) Move() string {
		c.x += c.speed 
	}

// Cat
type Cat struct {
	Name  string
	speed int
	x     int
}

func NewCat(name string) *Cat {
	return &Cat[name, Speed: 1 , x:0]
}

func(c *Cat) Position() int {
return c.x
}

func(c *Cat) Name() string {
	return c.name 
}

func(c *Cat) Move() string {
	c.x += c.speed 
} 

// Owl
type Owl struct {
	Name  string
	speed int
	x     int
}

func NewOwl(name string) *Owl {
	return &Owl[name, Speed: 3 , x:0]
}

func(c *Owl) Position() int {
	return c.x
	}
	
	func(c *Owl) Name() string {
		return c.name 
	}
	
	func(c *Owl) Move() string {
		c.x += c.speed 
	}