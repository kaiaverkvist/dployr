package routes

import (
	"github.com/kaiaverkvist/dployr/internal/web/context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strings"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {

		cc := c.(*context.DployrContext)

		data := map[string]interface{}{
			"hostname": cc.Container.Deployment.FullHostName,
			"envs":     cc.Container.Deployment.ExpectedEnvs,
			"dir":      cc.Container.Deployment.Dir,
		}

		return c.Render(200, "assets/templates/deploy.html", data)
	})

	e.POST("/", func(c echo.Context) error {
		cc := c.(*context.DployrContext)

		params, _ := c.FormParams()

		envs := make(map[string]string)

		for k, v := range params {
			envs[k] = strings.Join(v, "")
		}

		go func() {
			log.Info("Performing deployment ::: ")
			_, err := cc.Container.Deployment.PerformDeployment(envs)
			if err != nil {
				log.Error("Unable to deploy: ", err.Error())
			}
		}()

		data := map[string]interface{}{
			"output": cc.Container.Deployment.OutputBuffer.String(),
		}

		return c.Render(200, "assets/templates/deployment_result.html", data)
	})

	e.GET("/output", func(c echo.Context) error {
		cc := c.(*context.DployrContext)

		data := map[string]interface{}{
			"output": cc.Container.Deployment.OutputBuffer.String(),
		}

		return c.Render(200, "assets/templates/deployment_result.html", data)
	})
}
