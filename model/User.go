package model

import (
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUser(name string) (code int)  {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}

	return errmsg.SUCCSE //200
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

func CreateUser(data *User) int  {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR  // 500
	}
	
	return errmsg.SUCCSE // 200
}


func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return users
}

