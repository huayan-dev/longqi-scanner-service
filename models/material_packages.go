package models

/**
  @author:pandi
  @date:2022-11-03
  @note:
**/

type MaterialPackages struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	MaterialID      uint    `json:"materialID" gorm:"not null"`
	PackageStatus   string  `json:"packageStatus" gorm:"size:32;not null"`
	Mark            string  `json:"mark" gorm:"size:64;default:''"`
	Quantity        uint    `json:"quantity" gorm:"not null"`
	PackageType     string  `json:"packageType" gorm:"size:32;not null"`
	PalletSn        string  `json:"palletSn" gorm:"size:64;not null"`
	InsidePackageID uint    `json:"insidePackageID" gorm:"default:0"`
	CreatedAt       *MyTime `json:"createdAt"`
	UpdatedAt       *MyTime `json:"updatedAt"`
}
