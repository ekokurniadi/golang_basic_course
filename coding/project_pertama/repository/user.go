package repository

import (
	"fmt"
	"project_pertama/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	UploadProfilePicture(user entity.User) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	FindAll(search string, page int, size int) ([]entity.User, error)
	TotalFetchData(search string, page int, size int) (int, error)
	DeleteUser(ID int) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	//perintah di MYSQL INSERT INTO users VALUES('Eko','Laki-Laki','','eko123','12345','administrator')
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UploadProfilePicture(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindAll(search string, page int, size int) ([]entity.User, error) {
	var users []entity.User
	searchQuery := search
	sql := "SELECT * FROM users WHERE 1=1 "
	if searchQuery != "" {
		sql = fmt.Sprintf("%s AND (nama_lengkap LIKE '%%%s%%' OR user_name LIKE '%%%s%%') ", sql, searchQuery, searchQuery)
	}
	start := (page * size)
	limit := 0
	if start < 0 {
		limit = 0
	} else {
		limit = start
	}
	length := size

	sql = fmt.Sprintf("%s LIMIT %d,%d", sql, limit, length)
	fmt.Println("query mysql =>", sql)
	err := r.db.Raw(sql).Scan(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil

}

func (r *userRepository) TotalFetchData(search string, page int, size int) (int, error) {
	var users []entity.User
	searchQuery := search
	sql := "SELECT * FROM users WHERE 1=1 "
	// SELECT * FROM users where 1=1 and (nama_lengkap LIKE '%joko%' OR user_name like '%joko%')
	if searchQuery != "" {
		sql = fmt.Sprintf("%s AND (nama_lengkap like '%%%s%%' OR user_name like '%%%s%%') ", sql, searchQuery, searchQuery)
	}

	err := r.db.Raw(sql).Scan(&users).Error
	if err != nil {
		return len(users), err
	}
	return len(users), nil
}

func (r *userRepository) DeleteUser(ID int) (bool, error) {
	var user entity.User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
