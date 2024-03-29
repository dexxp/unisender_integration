package auth

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases"
)

type Impl struct {
	authUseCase usecases.Auth
}

func NewImpl(authUseCase usecases.Auth) *Impl {
	return &Impl{
		authUseCase,
	}
}
