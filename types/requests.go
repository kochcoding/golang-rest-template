package types

type AddItemRequest struct {
	Title string `json:"title"`
}

type ChangeItemStatus struct {
	Checked bool `json:"checked"`
}
