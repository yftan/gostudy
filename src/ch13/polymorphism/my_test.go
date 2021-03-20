package polymorphism

import (
	"fmt"
	"testing"
)

type A interface {
	test1()
	test2()
}

type B interface {
	test2()
	test4()
}


type C interface {
	A
	B
	test3()
}

type Dog struct {

}

func (dog *Dog) test1()  {
	fmt.Print("test1()")
}

func (dog *Dog) test2()  {
	fmt.Print("test2()")
}

func (dog *Dog) test4()  {
	fmt.Print("test4()")
}

func (dog *Dog) test3()  {
	fmt.Print("test3()")
}

func TestCat(t *testing.T) {
	dog := Dog{}
	var a1 A;
	var b1 B;
	var c1 C;
	//  如果接受者是 (dog *Dog)指针类型 a1 = &dog ，如果是(dog Dog) a1 = dog
	a1 = &dog
	b1 = &dog
	c1 = &dog

	a1.test1()
	a1.test2()
	b1.test2()
	b1.test4()
	c1.test3()
}

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Printf("%s Person show()\n" , p.name)
}

type People = Person

func (p People) show1()  {
	fmt.Printf("%s People show1()\n" , p.name)
}

func TestPerson(t *testing.T) {
	p1 := Person{"王二狗"}
	p2 := People{"李小花"}
	p1.show()
	p1.show1()
	p2.show()
	p2.show1()
}

