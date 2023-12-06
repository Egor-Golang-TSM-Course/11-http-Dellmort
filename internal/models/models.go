package models

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
