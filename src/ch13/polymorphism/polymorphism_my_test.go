package polymorphism

import (
	"fmt"
	"reflect"
	"testing"
)

type Animal interface {
	eat() string
}

type Cat struct {
	name string
	age  int
}

func (cat *Cat) eat() string {
	return "小猫咪吃猫粮"
}

type Bird struct {
	name string
	age  int
}

func (bird Bird) eat() string {
	return "小鸟吃鸟食"
}

func DoEat(animal Animal) {
	animal.eat()
}

func TestAnimal(t *testing.T) {
	bird := new(Bird)
	bird1 := Bird{}
	fmt.Println(reflect.TypeOf(bird1))
	DoEat(bird)
}
