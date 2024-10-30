package entity

import (
	"gorm.io/gorm"
)

// User is a struct that represents a user entity
type Customer struct {
	CustomerID        string    `gorm:"column:id;primaryKey"`
	Nama_Lengkap      string    `gorm:"column:nama_lengkap"`
	Alamat			  string	`gorm:"column:alamat"`
	NoTelepon		  string    `gorm:"column:no_telepon"`
	Email			  string    `gorm:"column:email"`
	TanggalLahir	  string	`gorm:"column:tanggal_lahir"`
	TanggalBergabung  string    `gorm:"column:tanggal_bergabung"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at"`
	
}

func (cus *Customer) TableName() string {
	return "customers"
}
