package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
	require.Nil(t, err)

	personExists = func(cpf string) bool {
		return false
	}

	err = checkinService.Register("234.234.234-01")
	require.Error(t, err)
}
