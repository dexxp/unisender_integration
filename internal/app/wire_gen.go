// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/controller/grpc"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/controller/restapi/auth"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/controller/restapi/disable"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/controller/restapi/unisender"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases/repo"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/pkg/mysql"
	"github.com/google/wire"
	"sync"
)

// Injectors from service_provider.go:

func Wire(db mysql.Database, client usecases.AmoClient, uniClient usecases.UnisenderClient, cfg config.Auth) *ServiceProvider {
	repoAccountRepo := ProvideAccountRepository(db)
	unisenderUseCase := ProvideUnisenderUC(repoAccountRepo, client, uniClient, cfg)
	repoContactRepo := ProvideContactRepository(db)
	contactUseCase := ProvideContactUC(repoContactRepo)
	repoEmailRepo := ProvideEmailRepository(db)
	emailUseCase := ProvideEmailUC(repoEmailRepo)
	impl := ProvideUnisenderHandler(unisenderUseCase, contactUseCase, emailUseCase)
	repoIntegrationRepo := ProvideIntegrationRepository(db)
	authUseCase := ProvideAuthUC(repoAccountRepo, repoIntegrationRepo, client, cfg)
	impl2 := ProvideAuthHandler(authUseCase)
	disableUseCase := ProvideDisableUC(repoAccountRepo, repoContactRepo, repoEmailRepo)
	impl3 := ProvideDisableHandler(disableUseCase)
	serviceProvider := &ServiceProvider{
		uImpl: impl,
		aImpl: impl2,
		dImpl: impl3,
	}
	return serviceProvider
}

func WireGrpc(db mysql.Database) *grpc.Implementation {
	repoAccountRepo := ProvideAccountRepository(db)
	repoContactRepo := ProvideContactRepository(db)
	repoEmailRepo := ProvideEmailRepository(db)
	disableUseCase := ProvideDisableUC(repoAccountRepo, repoContactRepo, repoEmailRepo)
	implementation := ProvideDImplHandler(disableUseCase)
	return implementation
}

// service_provider.go:

type ServiceProvider struct {
	uImpl *unisender.Impl
	aImpl *auth.Impl
	dImpl *disable.Impl
}

var (
	// Implementations
	authImpl     *auth.Impl
	authImplOnce sync.Once

	disableImpl     *disable.Impl
	disableImplOnce sync.Once

	unisenderImpl     *unisender.Impl
	unisenderImplOnce sync.Once

	dImpl     *grpc.Implementation
	dImplOnce sync.Once

	// UseCases
	authUC     *usecases.AuthUseCase
	authUCOnce sync.Once

	contactUC     *usecases.ContactUseCase
	contactUCOnce sync.Once

	emailUC     *usecases.EmailUseCase
	emailUCOnce sync.Once

	unisenderUC     *usecases.UnisenderUseCase
	unisenderUCOnce sync.Once

	disableUC     *usecases.DisableUseCase
	disableUCOnce sync.Once

	// Repositories
	accountRepo     *repo.AccountRepo
	accountRepoOnce sync.Once

	integrationRepo     *repo.IntegrationRepo
	integrationRepoOnce sync.Once

	contactRepo     *repo.ContactRepo
	contactRepoOnce sync.Once

	emailRepo     *repo.EmailRepo
	emailRepoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUnisenderHandler,
		ProvideAuthHandler,
		ProvideDisableHandler,

		ProvideAuthUC,
		ProvideUnisenderUC,
		ProvideEmailUC,
		ProvideContactUC,
		ProvideDisableUC,

		ProvideAccountRepository,
		ProvideIntegrationRepository,
		ProvideContactRepository,
		ProvideEmailRepository, wire.Bind(new(usecases.Auth), new(*usecases.AuthUseCase)), wire.Bind(new(usecases.Unisender), new(*usecases.UnisenderUseCase)), wire.Bind(new(usecases.Contact), new(*usecases.ContactUseCase)), wire.Bind(new(usecases.Email), new(*usecases.EmailUseCase)), wire.Bind(new(usecases.Disable), new(*usecases.DisableUseCase)), wire.Bind(new(usecases.AccountRepo), new(*repo.AccountRepo)), wire.Bind(new(usecases.IntegrationRepo), new(*repo.IntegrationRepo)), wire.Bind(new(usecases.ContactRepo), new(*repo.ContactRepo)), wire.Bind(new(usecases.EmailRepo), new(*repo.EmailRepo)), wire.Struct(new(ServiceProvider), "*"),
	)

	GRPCProviderSet wire.ProviderSet = wire.NewSet(
		ProvideDImplHandler,

		ProvideDisableUC,

		ProvideAccountRepository,
		ProvideContactRepository,
		ProvideEmailRepository, wire.Bind(new(usecases.Disable), new(*usecases.DisableUseCase)), wire.Bind(new(usecases.AccountRepo), new(*repo.AccountRepo)), wire.Bind(new(usecases.ContactRepo), new(*repo.ContactRepo)), wire.Bind(new(usecases.EmailRepo), new(*repo.EmailRepo)),
	)
)

func ProvideDImplHandler(uc usecases.Disable) *grpc.Implementation {
	dImplOnce.Do(func() {
		dImpl = grpc.NewImplementation(uc)
	})

	return dImpl
}

func ProvideAuthHandler(uc usecases.Auth) *auth.Impl {
	authImplOnce.Do(func() {
		authImpl = auth.NewImpl(uc)
	})

	return authImpl
}

func ProvideUnisenderHandler(uniUC usecases.Unisender, contactUC2 usecases.Contact, emailUC2 usecases.Email) *unisender.Impl {
	unisenderImplOnce.Do(func() {
		unisenderImpl = unisender.NewImpl(uniUC, contactUC2, emailUC2)
	})

	return unisenderImpl
}

func ProvideDisableHandler(disableUC2 usecases.Disable) *disable.Impl {
	disableImplOnce.Do(func() {
		disableImpl = disable.NewImpl(disableUC2)
	})

	return disableImpl
}

func ProvideAuthUC(accountRepo2 usecases.AccountRepo, integrationRepo2 usecases.IntegrationRepo, client usecases.AmoClient, cfg config.Auth) *usecases.AuthUseCase {
	authUCOnce.Do(func() {
		authUC = usecases.NewAuthUseCase(accountRepo2, integrationRepo2, client, cfg)
	})

	return authUC
}

func ProvideDisableUC(accountRepo2 usecases.AccountRepo, contactRepo2 usecases.ContactRepo, emailRepo2 usecases.EmailRepo) *usecases.DisableUseCase {
	disableUCOnce.Do(func() {
		disableUC = usecases.NewDisableUseCase(accountRepo2, contactRepo2, emailRepo2)
	})

	return disableUC
}

func ProvideContactUC(contactRepo2 usecases.ContactRepo) *usecases.ContactUseCase {
	contactUCOnce.Do(func() {
		contactUC = usecases.NewContactUseCase(contactRepo2)
	})

	return contactUC
}

func ProvideEmailUC(emailRepo2 usecases.EmailRepo) *usecases.EmailUseCase {
	emailUCOnce.Do(func() {
		emailUC = usecases.NewEmailUseCase(emailRepo2)
	})

	return emailUC
}

func ProvideUnisenderUC(accountRepo2 usecases.AccountRepo, amoClient usecases.AmoClient, uniClient usecases.UnisenderClient, cfg config.Auth) *usecases.UnisenderUseCase {
	unisenderUCOnce.Do(func() {
		unisenderUC = usecases.NewUnisenderUseCase(accountRepo2, amoClient, uniClient, cfg)
	})

	return unisenderUC
}

func ProvideAccountRepository(db mysql.Database) *repo.AccountRepo {
	accountRepoOnce.Do(func() {
		accountRepo = repo.NewAccountRepo(db)
	})

	return accountRepo
}

func ProvideIntegrationRepository(db mysql.Database) *repo.IntegrationRepo {
	integrationRepoOnce.Do(func() {
		integrationRepo = repo.NewIntegrationRepo(db)
	})

	return integrationRepo
}

func ProvideContactRepository(db mysql.Database) *repo.ContactRepo {
	contactRepoOnce.Do(func() {
		contactRepo = repo.NewContactRepo(db)
	})

	return contactRepo
}

func ProvideEmailRepository(db mysql.Database) *repo.EmailRepo {
	emailRepoOnce.Do(func() {
		emailRepo = repo.NewEmailRepo(db)
	})

	return emailRepo
}
