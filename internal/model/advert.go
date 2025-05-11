package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advert struct {
	gorm.Model
	ID                 uuid.UUID
	CarModel           Car
	Status             bool
	SellerID           uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
	Images             []string
	Year               int
	MileageKm          int
	OwnersCount        int
	EnsuranceValidTill *time.Time
	InteriorMaterial   InteriorMaterial
	LicensePlateNumber *string
	Color              Color
	ExchangePossible   bool
	TestDriveAvailable bool
	VIN                string
	CustomsCleared     bool
	PTS                string // Есть такое вообще?
	Addons             *AddonsList
	City               int
	LastService        *time.Time
	Condition          StateType
}
