package model

import (
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

func CheckUser(name string) (code int)  {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}

	return errmsg.SUCCSE //200
}

func CreateUser(data *User) int  {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR  // 500
	}
	
	return errmsg.SUCCSE // 200
}