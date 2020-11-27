package router

import (
	"github.com/kochcoding/golang-rest-template/app/handler/item"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	itemsPrefix = "/items"
)

// ServiceInterface ...
type ServiceInterface interface {
}

// Service ...
type Service struct {
	item item.ServiceInterface
}

func NewService(g *echo.Group) *Service {

	srv := &Service{
		item: item.NewService(),
	}

	srv.setItemRoutes(g.Group(itemsPrefix))

	return srv
}
