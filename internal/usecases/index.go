package usecases

import "github.com/google/wire"

var AccountUseCaseSet = wire.NewSet(NewAccountUseCase)
var EmailUseCaseSet = wire.NewSet(NewEmailUseCase)
var ContactUseCaseSet = wire.NewSet(NewContactUseCase)
var AuthUseCaseSet = wire.NewSet(NewAuthUseCase)
var IntegrationUseCaseSet = wire.NewSet(NewIntegrationUseCase)
var UnisenderUseCaseSet = wire.NewSet(NewUnisenderUseCase)
