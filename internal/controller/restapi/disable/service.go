package disable

import "git.amocrm.ru/dmiroshnikov/unisender_integration/internal/usecases"

type Impl struct {
	disableUC usecases.Disable
}

func NewImpl(disable usecases.Disable) *Impl {
	return &Impl{
		disableUC: disable,
	}
}
