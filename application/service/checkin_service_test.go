package service

import "testing"

var personExists func(cpf string) bool

type personRepositoryMock struct {
}

func (p personRepositoryMock) PersonExists(cpf string) bool {
	return personExists(cpf)
}

func TestCheckinService_Register(t *testing.T) {

	repositoryMock := personRepositoryMock{}
	checkinService := NewCheckinService(repositoryMock)

	personExists = func(cpf string) bool {
		return true
	}

	err := checkinService.Register("234.234.234-01")
	if err != nil {
		t.Error("return data should be nil")
	}

	personExists = func(cpf string) bool {
		return false
	}

	err = checkinService.Register("234.234.234-01")
	if err == nil {
		t.Error("return data should be not nil")
	}
}
