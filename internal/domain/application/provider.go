package application

import (
	"github.com/ebar-go/ego/utils/structure"
	"unknown/internal/domain/repo"
)

type provider struct {
	user *UserApp
}

var providerOnce = structure.NewSingleton(func() *provider {
	return &provider{user: NewUserApp(repo.Provider().User())}
})

func Provider() *provider {
	return providerOnce.Get()
}

func (p *provider) User() *UserApp {
	return p.user
}
