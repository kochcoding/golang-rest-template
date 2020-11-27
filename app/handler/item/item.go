package item

import (
	"net/http"
	"strconv"

	"github.com/kochcoding/golang-rest-template/app/repos/item"
	"github.com/kochcoding/golang-rest-template/types"
	"github.com/kochcoding/golang-rest-template/vars"
	"github.com/labstack/echo"
)

type ServiceInterface interface {
	GetItems() echo.HandlerFunc
	AddItem() echo.HandlerFunc
	RemoveItem() echo.HandlerFunc
	ChangeItemStatus() echo.HandlerFunc
	GetNumberOfItems() echo.HandlerFunc
}

func NewService() *Service {

	return &Service{
		itemRepo: item.NewRepo(),
	}
}

type Service struct {
	itemRepo item.RepoInterface
}

func (s *Service) GetItems() echo.HandlerFunc {
	return func(c echo.Context) error {

		items, err := s.itemRepo.GetItems()
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.GetItems(): failed to query the DB (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to query DB")
		}

		return c.JSON(http.StatusOK, items)
	}
}

func (s *Service) AddItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var addItemReq types.AddItemRequest
		if err := c.Bind(&addItemReq); err != nil || addItemReq.Title == "" {
			vars.LoggerErr.Printf("[WRN] item.AddItem(): Invalid request (%s)", err)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request")
		}

		err := s.itemRepo.AddItem(addItemReq.Title)
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.AddItem(): Failed to add item (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add item to DB")
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *Service) RemoveItem() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			vars.LoggerErr.Printf("[WRN] item.RemoveItem(): Invalid id (%s)", err)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid id")
		}

		err = s.itemRepo.RemoveItem(id)
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.RemoveItem(): Failed to remove item from DB (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to remove item from DB")
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *Service) ChangeItemStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var changeItemStatusReq types.ChangeItemStatus
		if err := c.Bind(&changeItemStatusReq); err != nil {
			vars.LoggerErr.Printf("[WRN] item.ChangeItemStatus(): Invalid request (%s)", err)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			vars.LoggerErr.Printf("[WRN] item.ChangeItemStatus(): Invalid id (%s)", err)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid id")
		}

		err = s.itemRepo.ChangeItemStatus(id, changeItemStatusReq.Checked)
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.RemoveItem(): Failed to change item status (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to change item status")
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *Service) GetNumberOfItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := s.itemRepo.GetNumberOfItems()
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.GetNumberOfItems(): Failed to query the DB (%s)", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to query the DB")
		}

		return c.JSON(http.StatusOK, count)
	}
}
