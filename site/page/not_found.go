package page

import (
	"github.com/labstack/echo/v4"
)

const (
	notFoundPageUri = "/*"
)

type NotFoundPage struct {
}

type NotFoundPageData struct {
}

func NewNotFoundPage() *Page {
	deps := &NotFoundPage{}

	return &Page{
		MenuID:      "not-found-page",
		Title:       "NotFound",
		Frame:       true,
		Path:        notFoundPageUri,
		Template:    "not-found.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *NotFoundPage) GetPageData(c echo.Context) any {
	return NotFoundPageData{}
}