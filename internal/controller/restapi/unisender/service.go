package unisender

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases"
)

type Impl struct {
	unisenderUseCase usecases.Unisender
	contactUseCase   usecases.Contact
	emailUseCase     usecases.Email
}

func NewImpl(unisenderUseCase usecases.Unisender, contactUseCase usecases.Contact, emailUseCase usecases.Email) *Impl {
	return &Impl{
		unisenderUseCase,
		contactUseCase,
		emailUseCase,
	}
}
