package server

// The structure represents transaction data
type transaction struct {
	Sum      float32 `json:"sum"`
	Sender   uint    `json:"sender"`
	Receiver uint    `json:"receiver"`
}
