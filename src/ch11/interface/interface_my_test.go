package _interface

import (
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

func TestCat(t *testing.T) {
	var a Animal
	a = new(Cat)
	t.Log(a.eat())
}
