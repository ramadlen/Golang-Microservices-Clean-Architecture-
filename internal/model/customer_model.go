package model

type CustomerResponse struct {
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

type VerifyCustomerRequest struct {
	Token string `validate:"required"`
}

type CreateCustomerRequest struct {
	CustomerID       	 string `json:"id" validate:"required"`
	Nama_Lengkap     	 string `json:"nama_lengkap" validate:"required,max=100"`
	Alamat     			 string `json:"alamat" validate:"required,max=100"`
	NoTelepon     		 string `json:"no_telepon" validate:"required,max=100"`
	Email     			 string `json:"email" validate:"required,max=100"`
	TanggalLahir    	 string `json:"tanggal_lahir" validate:"required,max=100"`
	TanggalBergabung     string `json:"tanggal_bergabung" validate:"required,max=100"`
}

type UpdateCustomerRequest struct {
	CustomerID       	 string `json:"-" validate:"required"`
	Nama_Lengkap     	 string `json:"nama_lengkap" validate:"required,max=100"`
	Alamat     			 string `json:"alamat" validate:"required,max=100"`
	NoTelepon     		 string `json:"no_telepon" validate:"required,max=100"`
	Email     			 string `json:"email" validate:"required,max=100"`
	TanggalLahir    	 string `json:"tanggal_lahir" validate:"required,max=100"`
	TanggalBergabung     string `json:"tanggal_bergabung" validate:"required,max=100"`
}

type SearchCustomerRequest struct {
	CustomerID       	 string `json:"-" validate:"required"`
	Nama_Lengkap     	 string `json:"nama_lengkap" validate:"required,max=100"`
	Alamat     			 string `json:"alamat" validate:"required,max=100"`
	NoTelepon     		 string `json:"no_telepon" validate:"required,max=100"`
	Email     			 string `json:"email" validate:"required,max=100"`
	TanggalLahir    	 string `json:"tanggal_lahir" validate:"required,max=100"`
	TanggalBergabung     string `json:"tanggal_bergabung" validate:"required,max=100"`
	Page   int    `json:"page" validate:"min=1"`
	Size   int    `json:"size" validate:"min=1,max=100"`
}

type GetCustomerRequest struct {
	CustomerID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteCustomerRequest struct {
	CustomerID     string `json:"-" validate:"required,max=100,uuid"`
}