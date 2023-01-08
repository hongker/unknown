package service

import (
	"github.com/ebar-go/ego/utils/structure"
	"unknown/internal/domain/application"
)

type provider struct {
	web  *WebService
	gate *GateService
}

var providerOnce = structure.NewSingleton(func() *provider {
	return &provider{
		web:  NewWebService(application.Provider().User()),
		gate: NewGateService(application.Provider().User()),
	}
})

func Provider() *provider {
	return providerOnce.Get()
}

func (p *provider) Web() *WebService {
	return p.web
}

func (p *provider) Gate() *GateService {
	return p.gate
}
