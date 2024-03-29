package usecases

import (
	"errors"
	"fmt"
)

type DisableUseCase struct {
	accountRepo AccountRepo
	contactRepo ContactRepo
	emailRepo   EmailRepo
}

func NewDisableUseCase(accountRepo AccountRepo, contactRepo ContactRepo, emailRepo EmailRepo) *DisableUseCase {
	return &DisableUseCase{
		accountRepo: accountRepo,
		contactRepo: contactRepo,
		emailRepo:   emailRepo,
	}
}

func (d *DisableUseCase) Disable(accountID uint64) error {

	fmt.Println("Disable UseCase")
	account, _ := d.accountRepo.Get(accountID)

	err := d.accountRepo.Remove(account)

	if err != nil {
		fmt.Println("Disable Use Case: ", err)
		return errors.New("Не получилось отключить аккаунт!")
	}

	return nil
}
