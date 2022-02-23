package context

import (
	"github.com/kaiaverkvist/dployr/pkg/deploy"
	"github.com/labstack/echo/v4"
	"sync"
)

type DployrContext struct {
	echo.Context

	Container *WebDataContainer
}

// WebDataContainer is a reference to
type WebDataContainer struct {
	Deployment  *deploy.Deployment
	Environment map[string]string

	mtx sync.Mutex
}

func NewWebDataContainer(deployment *deploy.Deployment) WebDataContainer {
	return WebDataContainer{
		Deployment:  deployment,
		Environment: make(map[string]string),
	}
}
