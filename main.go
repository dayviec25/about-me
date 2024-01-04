package main

import (
	"about-me/data"
	"about-me/handlers"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TemplateRegistry is a custom HTML template renderer
type TemplateRegistry struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", name, err)
		return err
	}
	return nil
}

func parseTemplates(patterns []string) (*template.Template, error) {
	tmpl := template.New("")
	for _, pattern := range patterns {
		_, err := tmpl.ParseGlob(pattern)
		if err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}

func main() {
	e := echo.New()
	e.Static("/dist", "dist")

	// Initialize the template engine
	patterns := []string{"templates/**/*.tmpl"}
	templates, err := parseTemplates(patterns)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Debug = true
	e.Renderer = &TemplateRegistry{templates: templates}
	handlers.RegisterReviewHandlers(e)
	e.GET("/", func(c echo.Context) error {
		pageData, err := data.GetAboutMeData()
		if err != nil {
			log.Printf("Error getting project data: %v", err)
			return nil
		}
		return c.Render(http.StatusOK, "index.tmpl", pageData)
	})
	e.GET("/:page", func(c echo.Context) error {
		page := c.Param("page")
		data := getPageData(page)
		return c.Render(http.StatusOK, page+".tmpl", data)
	})

	// Custom error handler
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Start HTTP server
	go func() {
		if err := e.Start(":80"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the HTTP server")
		}
	}()

	// Start HTTPS server
	e.Logger.Fatal(e.StartTLS(":443", "/etc/letsencrypt/live/dchung.dev/cert.pem", "/etc/letsencrypt/live/dchung.dev/privatekey.pem"))
}

func getPageData(page string) interface{} {
	if page == "index" {
		pageData, err := data.GetAboutMeData()
		if err != nil {
			log.Printf("Error getting project data: %v", err)
			return nil
		}
		return pageData
	}
	return nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Printf("Error occurred: %v", err)
	if he, ok := err.(*echo.HTTPError); ok {
		c.String(he.Code, he.Message.(string))
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
