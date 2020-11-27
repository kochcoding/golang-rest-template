package types

import "time"

type GetItemsResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Checked   bool      `json:"checked"`
	UpdatedAt time.Time `json:"updatedAt"`
}
