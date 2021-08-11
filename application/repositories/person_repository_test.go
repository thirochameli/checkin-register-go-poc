package repositories

import "testing"

func TestPersonRepositoryDb_PersonExists(t *testing.T) {

	var personRepository PersonRepository
	personRepository = NewPersonRepositoryDb()

	existentCpf := "123.456.789-02"
	exists := personRepository.PersonExists(existentCpf)
	if !exists {
		t.Error("return data should be true")
	}

	nonExistentCpf := "123.456.789-10"
	exists = personRepository.PersonExists(nonExistentCpf)
	if exists {
		t.Error("return data should be false")
	}
}
