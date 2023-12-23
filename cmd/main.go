package main

import "log"
import "time"
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
  Type string
	Start time.Time
	Stop time.Time
}

type PageIndex struct {
	Timings []Timing
}

var timings []Timing
var timingType string

func main() {
	timings = make([]Timing, 0)
  timingType = ""

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
    return c.Render(200, "index.html", PageIndex{Timings: timings})
	})

	e.POST("/increase", func(c echo.Context) error {
    log.Print("increase")

    if timingType != "" {
      timings[len(timings)-1].Stop = time.Now()
    }

    timingType = "increase"
    timings = append(timings, Timing{Type: "increase", Start: time.Now()})
  
    log.Print(timings)

    return c.Render(200, "timings.html", PageIndex{Timings: timings})
	})

	e.POST("/decrease", func(c echo.Context) error {
    if timingType != "" {
      timings[len(timings)-1].Stop = time.Now()
    }

    timingType = "decrease"
    timings = append(timings, Timing{Type: "decrease", Start: time.Now()})

    return c.Render(200, "timings.html", PageIndex{Timings: timings})
	})


	e.POST("/clear", func(c echo.Context) error {
    timingType = ""
    timings = make([]Timing, 0)

    return c.Render(200, "timings.html", PageIndex{Timings: timings})
	})

	e.Logger.Fatal(e.Start(":42069"))
}
