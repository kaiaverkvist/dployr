package web

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/CloudyKit/jet/v6/loaders/httpfs"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templateFolder string
	jetviews       *jet.Set
}

func NewTemplateRenderer(templateFolder string, fs http.FileSystem) TemplateRenderer {
	loader, _ := httpfs.NewLoader(fs)

	jetset := jet.NewSet(
		loader,
		jet.InDevelopmentMode(), // remove in production
	)

	return TemplateRenderer{
		templateFolder: templateFolder,
		jetviews:       jetset,
	}
}

// Render renders a template document
func (t TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	var datamap map[string]interface{}
	if data == nil {
		datamap = make(map[string]interface{})
	} else {
		datamap = data.(map[string]interface{})
	}

	view, err := t.jetviews.GetTemplate(name)
	if err != nil {
		return echo.NewHTTPError(500, "Template rendering error: "+err.Error())
	}

	return view.Execute(w, nil, datamap)
}
