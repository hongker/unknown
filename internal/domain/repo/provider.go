package repo

import "github.com/ebar-go/ego/utils/structure"

type provider struct {
	user *UserRepo
}

var providerOnce = structure.NewSingleton(func() *provider {
	return &provider{user: NewUserRepo()}
})

func Provider() *provider {
	return providerOnce.Get()
}

func (p *provider) User() *UserRepo {
	return p.user
}
