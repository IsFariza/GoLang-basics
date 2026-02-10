package domain

type Emulation struct {
	ID   string `json:"_id"`
	Name string `json:"name" binding:"required"`
}
