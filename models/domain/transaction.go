package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Date      time.Time
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Association
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
