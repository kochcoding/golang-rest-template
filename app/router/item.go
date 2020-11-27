package router

import (
	"github.com/labstack/echo"
)

func (s *Service) setItemRoutes(g *echo.Group) {
	g.GET("/", s.item.GetItems())
	g.POST("/", s.item.AddItem())
	g.GET("/count", s.item.GetNumberOfItems())
	g.DELETE("/:id", s.item.RemoveItem())
	g.PATCH("/:id", s.item.ChangeItemStatus())
}
