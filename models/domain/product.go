package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	ProductName   string    `gorm:"index;not null;unique"`
	PurchasePrice float64   `gorm:"not null"`
	SellingPrice  float64   `gorm:"not null"`
	Stock         int       `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	// Association
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ProductID"`
}
