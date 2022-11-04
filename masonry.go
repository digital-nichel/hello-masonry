package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	e := echo.New()
	e.Static("/static/", "static")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.gohtml")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		coverImage := fmt.Sprintf("/static/images/cat-0001%d.jpeg", rand.Intn(10))

		var images []string
		for i := 1; i < 20; i++ {
			idx := rand.Intn(10)
			images = append(images, fmt.Sprintf("/static/images/cat-0001%d.jpeg", idx))
		}

		return c.Render(http.StatusOK, "masonry.gohtml", map[string]interface{}{
			"name":   "masonry",
			"cover":  coverImage,
			"images": images,
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
