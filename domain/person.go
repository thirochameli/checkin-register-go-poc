package domain

type Person struct {
	Cpf string
	Name string
}

func NewPerson(cpf string, name string) *Person {
	return &Person{Cpf: cpf, Name: name}
}
