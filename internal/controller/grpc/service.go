package grpc

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/account_v1"
)

type Implementation struct {
	account_v1.UnimplementedAccountServiceServer
	disable usecases.Disable
}

func NewImplementation(disable usecases.Disable) *Implementation {
	return &Implementation{
		disable: disable,
	}
}
