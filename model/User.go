package model

import (
	"encoding/base64"
	"ginBlog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	//data.Password = ScryptPassword(data.Password)
	data.BeforeSave()
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

func DeleteUser(id int)  int  {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE
}

func EditUser(id int, data *User) int  {
	var user User
	var maps = make(map[string]interface{})

	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE
}

func (u *User) BeforeSave() {
	u.Password = ScryptPassword(u.Password)
}

func ScryptPassword(password string) string  {
	const KeyLen = 10
	salt := make([]byte,8)
	salt = []byte{12,32,4,6,66,22,222,11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8,1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}

	finalPassword:= base64.StdEncoding.EncodeToString(HashPw)
	return finalPassword

}