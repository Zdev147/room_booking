package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Zdev147/room_booking/pkg/config"
	"github.com/Zdev147/room_booking/pkg/models"
)

const templatePath = "./templates/"

var appConfig *config.AppConfig

func SetConfig(a *config.AppConfig) {
	appConfig = a
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data *models.TemplateData) {
	var cache map[string]*template.Template
	if appConfig.UseCache {
		cache = appConfig.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}

	cache[tmplName].Execute(w, data)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pageMatches, err := filepath.Glob(fmt.Sprintf("%s*.page.tmpl", templatePath))
	if err != nil {
		return cache, err
	}

	for _, match := range pageMatches {
		name := filepath.Base(match)
		parsedTemplate, err := template.New(name).ParseFiles(match)
		if err != nil {
			return cache, err
		}

		layoutMatches, err := filepath.Glob(fmt.Sprintf("%s*.layout.tmpl", templatePath))
		if err != nil {
			return cache, err
		}

		if len(layoutMatches) > 0 {
			parsedTemplate, err = parsedTemplate.ParseGlob(fmt.Sprintf("%s*.layout.tmpl", templatePath))
			if err != nil {
				return cache, err
			}
		}
		cache[name] = parsedTemplate
	}

	return cache, nil
}
