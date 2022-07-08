package server

type AccountForResponse struct {
	Id    uint    `json:"id"`
	Money float32 `json:"money"` // Stored in kopeks
}
