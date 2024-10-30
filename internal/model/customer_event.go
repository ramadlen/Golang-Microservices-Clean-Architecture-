package model

type CustomerEvent struct {
	CustomerID 	 	 string `json:"id,omitempty"`
	Nama_Lengkap 	 string `json:"nama_lengkap,omitempty"`
	Alamat		 	 string `json:"alamat,omitempty"`
	NoTelepon	 	 string `json:"no_telepon,omitempty"`
	Email		 	 string `json:"email,omitempty"`
	TanggalLahir 	 string `json:"tanggal_lahir,omitempty"`
	TanggalBergabung string `json:"tanggal_bergabung,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (u *CustomerEvent) GetId() string {
	return u.CustomerID
}
