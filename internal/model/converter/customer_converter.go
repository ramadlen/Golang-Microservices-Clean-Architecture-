package converter

import (
	"Customer/internal/entity"
	"Customer/internal/model"
)

func CustomerToResponse(contact *entity.Customer) *model.CustomerResponse {
	return &model.CustomerResponse{
		CustomerID:        contact.CustomerID,
		Nama_Lengkap: contact.Nama_Lengkap,
		Alamat:  contact.Alamat,
		Email:     contact.Email,
		NoTelepon:     contact.NoTelepon,
		TanggalLahir: contact.TanggalLahir,
		TanggalBergabung: contact.TanggalBergabung,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}
}

func CustomerToEvent(contact *entity.Customer) *model.CustomerEvent {
	return &model.CustomerEvent{
		CustomerID:        contact.CustomerID,
		Nama_Lengkap: contact.Nama_Lengkap,
		Alamat:  contact.Alamat,
		Email:     contact.Email,
		NoTelepon:     contact.NoTelepon,
		TanggalLahir: contact.TanggalLahir,
		TanggalBergabung: contact.TanggalBergabung,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}
}
