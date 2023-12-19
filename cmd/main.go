package main

import "log"
import "io"
import "html/template"
import "github.com/labstack/echo/v4"
import "github.com/labstack/echo/v4/middleware"

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Timing struct {
	Start time.Duration
	Stop time.Duration
}

type PageIndex struct {
	timings Timing[]
}

var timings []Timing

func main() {
	timings = make([]Timing, 0)
	templates, err := template.ParseGlob("templates/*.html")

	if err != nil {
		log.Fatalf("Failed: %s", err)
	}

	e := echo.New()

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	e.Use(middleware.Logger())

	e.Static("/dist", "./dist")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
	})

	e.POST("/timing", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
