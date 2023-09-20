package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	ProductName   string
	PurchasePrice float64
	SellingPrice  float64
	Stock         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	// Association
	TransactionDetails []TransactionDetail `gorm:"foreignKey:ProductID"`
}
