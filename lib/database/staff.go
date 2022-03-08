package database

import (
	"errors"

	"github.com/adicipta/Technical-Test-Klik-Digital-Sinergi/models"
	"gorm.io/gorm"
)

type StaffDB struct {
	db *gorm.DB
}

type StaffModel interface {
	CreateStaff(staff models.Staff) (models.Staff, error)
	GetAllStaff([]models.ResponseGetStaff, error)
	GetStaffByID(staffID int) (models.ResponseGetStaff, error)
	UpdateStaff(id int, newStaff models.Staff) (models.Staff, error)
}

func (m *StaffDB) CreateStaff(staff models.Staff) (models.Staff, error) {
	if err := m.db.Save(&staff).Error; err != nil {
		return staff, err
	}

	return staff, nil
}

func (m *StaffDB) GetAllStaff() ([]models.ResponseGetStaff, error) {
	var allstaff []models.ResponseGetStaff

	if err := m.db.Raw("SELECT s.id, s.phone_number, s.name, l.username, l.role FROM staffs s JOIN logins l ON s.id = l.staff_id").Scan(&allstaff).Error; err != nil {
		return nil, err
	} else if len(allstaff) == 0 {
		err := errors.New("is empty")
		return nil, err
	}

	return allstaff, nil
}

func (m *StaffDB) GetStaffByID(staffID int) (models.ResponseGetStaff, error) {
	var staff models.ResponseGetStaff

	if err := m.db.Raw("SELECT s.id, s.phone_number, s.name, l.username, l.role FROM staffs s JOIN logins l ON s.id = l.staff_id WHERE l.staff_id = ?", staffID).Scan(&staff).Error; err != nil {
		return staff, err
	}

	return staff, nil
}

func (m *StaffDB) UpdateStaff(id int, newStaff models.Staff) (models.Staff, error) {
	var staff models.Staff

	if err := m.db.First(&staff, id).Error; err != nil {
		return staff, err
	}

	staff.Name = newStaff.Name
	staff.PhoneNumber = newStaff.PhoneNumber

	if err := m.db.Save(&staff).Error; err != nil {
		return staff, err
	}

	return staff, nil
}
