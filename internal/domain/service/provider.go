package service

import (
	"github.com/ebar-go/ego/utils/structure"
	"unknown/internal/domain/application"
)

type provider struct {
	web *WebService
}

var providerOnce = structure.NewSingleton(func() *provider {
	return &provider{
		web: NewWebService(application.Provider().User()),
	}
})

func Provider() *provider {
	return providerOnce.Get()
}

func (p *provider) Web() *WebService {
	return p.web
}
