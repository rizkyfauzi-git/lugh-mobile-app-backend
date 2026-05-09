package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name" binding:"required"`
	Type      string         `gorm:"type:enum('income','expense');not null" json:"type" binding:"required,oneof=income expense"`
	Icon      string         `json:"icon"`
	UserID    uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Wallet struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name" binding:"required"`
	UserID    uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Transaction struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id"`
	WalletID    uint           `json:"wallet_id" binding:"required"`
	CategoryID  uint           `json:"category_id" binding:"required"`
	Amount      float64        `gorm:"type:decimal(15,2);not null" json:"amount" binding:"required"`
	Type        string         `gorm:"type:enum('income','expense');not null" json:"type" binding:"required,oneof=income expense"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Wallet   Wallet   `gorm:"foreignKey:WalletID" json:"wallet"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

type TransactionRequest struct {
	WalletID    uint      `json:"wallet_id" binding:"required"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}
