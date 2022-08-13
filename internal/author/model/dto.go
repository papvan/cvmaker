package model

import "time"

type CreateAuthorDTO struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	IsAlive   bool      `json:"is_alive"`
	CreatedAt time.Time `json:"created_at"`
}
