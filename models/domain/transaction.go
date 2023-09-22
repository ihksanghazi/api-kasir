package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Total     float64   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Association
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
