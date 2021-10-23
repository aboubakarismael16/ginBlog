package model

import (
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
}



func CreateArticle(data *Article) int  {

	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR  // 500
	}

	return errmsg.SUCCSE // 200
}

func GetCategoryArticle(id int,pageSize int, pageNum int)  ([]Article, int){
	var catArticleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&catArticleList).Error
	if err != nil {
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}

	return catArticleList, errmsg.SUCCSE
}

func GetArticleInfo(id int) (Article, int)  {
	var article Article
	err := db.Preload("Category").Where("id = ?",id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}

	return article, errmsg.SUCCSE
}

func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}

	return articleList,errmsg.SUCCSE
}

func DeleteArticle(id int)  int  {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE
}

func EditArticle(id int, data *Article) int  {
	var article Article
	var maps = make(map[string]interface{})

	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCSE
}