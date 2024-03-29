package repo

import "github.com/google/wire"

var AccountRepoSet = wire.NewSet(NewAccountRepo)
var ContactRepoSet = wire.NewSet(NewContactRepo)
var IntegrationRepoSet = wire.NewSet(NewIntegrationRepo)
var EmailRepoSet = wire.NewSet(NewEmailRepo)
