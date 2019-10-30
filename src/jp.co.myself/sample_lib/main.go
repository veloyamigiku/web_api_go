package sample_lib

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age int
}

// Say function
func (p *Person) Say() {
	fmt.Printf("My name is %s, I'm %d\n", p.Name, p.Age)
}
