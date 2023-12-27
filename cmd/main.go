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
  Id int
}

type PageIndex struct {
	Timings []Timing
  TotalTime time.Duration
}

var timings []Timing
var timingType string
var id int

func main() {
	timings = make([]Timing, 0)
  timingType = ""
  id  = 0

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
    addTiming("increase")
    data := getPageData()

    return c.Render(200, "index.html", data)
	})

	e.POST("/decrease", func(c echo.Context) error {
    addTiming("decrease")
    data := getPageData()

    return c.Render(200, "index.html", data)
	})


	e.POST("/clear", func(c echo.Context) error {
    timingType = ""
    timings = make([]Timing, 0)

    return c.Render(200, "index.html", PageIndex{Timings: timings})
	})

	e.Logger.Fatal(e.Start(":42069"))
}

func incrementId() {
  id += 1
}

func addTiming(t string) {
  if timingType != "" {
    timings[len(timings)-1].Stop = time.Now()
  }

  if timingType == t {
    timingType = ""
    return
  }

  timingType = t
  timings = append(timings, Timing{Type: t, Start: time.Now(), Id: id})
  incrementId()
}

func getPageData() PageIndex {
  var totalTime time.Duration = 0

  for _, timing := range timings {
    if !timing.Start.IsZero() && !timing.Stop.IsZero() {
      if timing.Type == "increase" {
        totalTime += timing.Stop.Sub(timing.Start)
      } else {
        totalTime -= timing.Stop.Sub(timing.Start)
      }
    }
  }

  return PageIndex{Timings: timings, TotalTime: totalTime}
}
