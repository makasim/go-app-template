package greeter

import "fmt"

type Greeter struct {
	greet string
}

func New(greet string) *Greeter {
	return &Greeter{
		greet: greet,
	}
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("%s %s!", g.greet, name)
}
