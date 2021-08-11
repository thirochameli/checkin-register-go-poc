package repositories

import "github.com/thirocha/domain"

var persons = make([]*domain.Person, 0)

type PersonRepository interface {
	PersonExists(cpf string) bool
}

type PersonRepositoryDb struct {
}

func init() {
	persons = append(persons,
		domain.NewPerson("123.456.789-01", "João"),
		domain.NewPerson("123.456.789-02", "Maria"),
		domain.NewPerson("123.456.789-03", "José"))
}

func NewPersonRepositoryDb() *PersonRepositoryDb {
	return &PersonRepositoryDb{}
}

func (p *PersonRepositoryDb) PersonExists(cpf string) bool {

	for _, person := range persons {
		if person.Cpf == cpf {
			return true
		}
	}

	return false
}
