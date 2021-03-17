package extension

import (
	"fmt"
	"testing"
)

type Animal struct {
}

func (animal *Animal) say() {
	fmt.Println("Animal say...")
}

func (animal Animal) walk() {
	fmt.Println("Animal walk...")
}

type Bird struct {
	Animal
}

func (bird *Bird) walk() {
	fmt.Println("bird fly better than walk")
}

func TestAnimal(t *testing.T) {
	bird := new(Bird)
	bird.say()
	bird.walk()
}
