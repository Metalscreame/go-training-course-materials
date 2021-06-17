package main

import "fmt"

type Greter interface {
	SayHello()
}

type Person struct {
	name string
}

func (p *Person) SayHello() {
	fmt.Printf("%s is saying 'Hi'\n", p.name)
}

func (p *Person) Sleep() {
	fmt.Printf("%s is sleeping\n", p.name)
}

func (p *Person) Walk() {
	fmt.Printf("%s is making his way downtown, walking fast...\n", p.name)
}

type JamesBond struct {
	Person
	gun string
}

func (j *JamesBond) Shot() string {
	return "TOP SECRET"
}

func (j *JamesBond) Walk() {
	fmt.Printf("%s is walking with his %v gun in his hand\n", j.name, j.gun)
}

func main() {

	p := Person{
		name: "Harry Potter",
	}

	p.SayHello()
	p.Walk()

	fmt.Println(p == Greter)

	bond := JamesBond{
		Person: Person{
			name: "James",
		},
		gun: "Magnum",
	}

	bond.SayHello()
	bond.Walk()
	bond.Sleep()


	/*
	we can call
	💡 We call it dynamic because we can assign s with a new struct of a different struct type which also implements the interface Shape.
	 */
	var human Human
	human = &p
	fmt.Printf("(%v, %T)\n", human, human)

	human = &bond
	fmt.Printf("(%v, %T)\n", human, human)
}


/*
Embedding interfaces
In Go, an interface cannot implement other interfaces or extend them,
but we can create a new interface by merging two or more interfaces.
*/

type Human interface {
	Greter
	Walker
}

type Walker interface {
	Walk()
}