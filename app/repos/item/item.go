package item

import (
	"database/sql"
	"strings"
	"time"

	"github.com/kochcoding/golang-rest-template/types"
	"github.com/kochcoding/golang-rest-template/vars"
	_ "github.com/lib/pq"
)

type RepoInterface interface {
	GetItems() ([]types.GetItemsResponse, error)
	AddItem(title string) error
	ChangeItemStatus(id int, status bool) error
	RemoveItem(id int) error
	GetNumberOfItems() (int, error)
	CleanOldItems() error
}

type Repo struct {
	db *sql.DB
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) GetItems() ([]types.GetItemsResponse, error) {
	rows, err := vars.DB.Query("SELECT * FROM items")
	if err != nil {
		vars.LoggerErr.Printf("[ERR] item.GetItems(): failed to query the DB (%s)", err)
		return nil, err
	}

	items := make([]types.GetItemsResponse, 0, 8)

	for rows.Next() {
		var id int
		var title string
		var checked bool
		var updatedAt time.Time

		err := rows.Scan(&id, &title, &checked, &updatedAt)
		if err != nil {
			vars.LoggerErr.Printf("[ERR] item.GetItems(): failed to scan item (%s)", err)
			return nil, err
		}

		items = append(items, types.GetItemsResponse{
			Id:        id,
			Title:     title,
			Checked:   checked,
			UpdatedAt: updatedAt,
		})
	}

	return items, nil
}

func (r *Repo) AddItem(title string) error {
	_, err := vars.DB.Exec("INSERT INTO items(title, updated_at) VALUES ($1, $2)", strings.Title(title), time.Now())
	if err != nil {
		vars.LoggerErr.Printf("[ERR] item.AddItem(): failed to add item to DB (%s)", err)
		return err
	}
	return nil
}

func (r *Repo) ChangeItemStatus(id int, status bool) error {
	_, err := vars.DB.Exec("UPDATE items SET checked = $1, updated_at = NOW()  WHERE id = $2", status, id)
	if err != nil {
		vars.LoggerErr.Printf("[WRN] item.RemoveItem(): failed to change status of item (%s)", err)
		return err
	}

	return nil
}

func (r *Repo) RemoveItem(id int) error {
	_, err := vars.DB.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		vars.LoggerErr.Printf("[WRN] item.RemoveItem(): failed to remove item (%s)", err)
		return err
	}

	return nil
}

func (r *Repo) GetNumberOfItems() (int, error) {
	row := vars.DB.QueryRow("SELECT COUNT(*) FROM items WHERE checked = FALSE")
	var count int

	err := row.Scan(&count)
	if err != nil {
		vars.LoggerErr.Printf("[WRN] item.GetNumberOfItems(): failed to query the DB (%s)", err)
		return -1, err
	}

	return count, nil
}

func (r *Repo) CleanOldItems() error {
	_, err := vars.DB.Exec("DELETE FROM items WHERE checked = TRUE AND updated_at < $1", time.Now().Add(-24*time.Hour))
	if err != nil {
		vars.LoggerErr.Printf("[WRN] item.RemoveItem(): failed to remove item (%s)", err)
		return err
	}

	return nil
}
