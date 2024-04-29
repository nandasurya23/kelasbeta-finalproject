package mockstruct

import (
	"kelasbeta/finalproject/models"
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID         uint              `json:"id"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  gorm.DeletedAt    `json:"deleted_at"`
	Identifier string            `json:"identifier"`
	Name       string            `json:"name"`
	Questions  []models.Question `json:"questions"`
}
