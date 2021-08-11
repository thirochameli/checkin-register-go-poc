package service

import (
	"fmt"
	"github.com/thirocha/application/repositories"
)

type CheckinService struct {
	personRepository repositories.PersonRepository
}

func NewCheckinService(personRepository repositories.PersonRepository) *CheckinService {
	return &CheckinService{personRepository: personRepository}
}

func (checkinService *CheckinService) Register(cpf string) error {

	exists := checkinService.personRepository.PersonExists(cpf)
	if !exists {
		return fmt.Errorf("Cpf '%s' not found for check-in", cpf)
	}

	fmt.Printf("Registered sucessfuly, cpf: '%s'", cpf)

	return nil
}
