package repository

import (
	"Customer/internal/entity"
	"Customer/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Repository[entity.Customer]
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}

func (r *CustomerRepository) FindById(db *gorm.DB, customer *entity.Customer, id string) error {
	return db.Where("id = ?", id).Take(customer).Error
}

func (r *CustomerRepository) Search(db *gorm.DB, request *model.SearchCustomerRequest) ([]entity.Customer, int64, error) {
	var contacts []entity.Customer
	if err := db.Scopes(r.FilterCustomer(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Customer{}).Scopes(r.FilterCustomer(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func (r *CustomerRepository) FilterCustomer(request *model.SearchCustomerRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.CustomerID)

		if name := request.Nama_Lengkap; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
		}

		if phone := request.NoTelepon; phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}

		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}
		if alamat := request.Alamat; alamat != "" {
			alamat = "%" + alamat + "%"
			tx = tx.Where("alamat LIKE ?", alamat)
		}

		return tx
	}
}
