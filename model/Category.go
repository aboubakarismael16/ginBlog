package model

import (
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}


func CheckCategory(name string) (code int)  {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}

	return errmsg.SUCCESS //200
}


func CreateCategory(data *Category) int  {
	//data.Password = ScryptPassword(data.Password)
	//data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR  // 500
	}

	return errmsg.SUCCESS // 200
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}


func GetCategories(pageSize int, pageNum int) ([]Category, int64) {
	var categories []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&categories).Error
	db.Model(&categories).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}

	return categories, total
}

func DeleteCategory(id int)  int  {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

func EditCategory(id int, data *Category) int  {
	var category Category
	var maps = make(map[string]interface{})

	maps["name"] = data.Name

	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

