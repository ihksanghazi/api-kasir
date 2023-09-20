package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	TransactionID uuid.UUID `gorm:"foreignKey;type:uuid"`
	ProductID     uuid.UUID `gorm:"foreignKey;type:uuid"`
	Amount        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
