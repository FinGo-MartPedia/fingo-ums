package services

import "github.com/fingo-martPedia/fingo-ums/internal/interfaces"

type Healthcheck struct {
	HealthcheckRepository interfaces.IHealthcheckRepo
}

func NewHealthcheckService() *Healthcheck {
	return &Healthcheck{}
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "OK", nil
}
