package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type tb_new struct {
	New_id int `gorm:"column:new_id"`
	New_title string `gorm:"column:new_title"`
	New_time time.Time `gorm:"column:new_time"`
	New_content string `gorm:"column:new_content"`
	New_attachment string `gorm:"column:new_attachment"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}


func AddNew(a *gin.Context) (error, tb_new) {
	var newdata tb_new
	a.ShouldBind(&newdata)
	newdata.Update_userid = "0"
	newdata.Is_del = false
	newdata.Update_time = time.Now()
	err := db.Table("tb_new").Select("new_title", "new_time", "new_content", "new_attachment", "is_del", "update_userid", "update_time").Create(&newdata).Error
	return err, newdata
}

func GetAllNew(messageParam map[string]interface{}) (error, []tb_new, int64) {
	var news []tb_new
	page := messageParam["page"].(int)
	pageSize := messageParam["limit"].(int)
	searchText := messageParam["searchText"].(string)
	var total int64
	err := db.Table("tb_new").Where("new_title like ? and is_del = false", "%"+searchText+"%").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&news).Error
	return err, news, total
}

func DelNew(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_new").Where("new_id = ?", id).Update("is_del", true).Error
	return err
}

func UpdateNew(a *gin.Context) (error, tb_new) {
	var newdata tb_new
	a.BindJSON(&newdata)
	newdata.Is_del = false
	newdata.Update_userid = "0"
	newdata.Update_time = time.Now()
	err := db.Table("tb_new").Select("new_title", "new_time", "new_content", "new_attachment", "is_del", "update_userid", "update_time").Where("new_id = ?", newdata.New_id).Updates(&newdata).Error
	return err, newdata
}

func GetFourNew() (error, []tb_new) {
	var news []tb_new
	//page := messageParam["page"].(int)
	//pageSize := messageParam["limit"].(int)
	//var total int64
	//err := db.Table("tb_message").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&messages).Error
	err := db.Table("tb_new").Select("new_id", "new_title", "new_time", "new_content", "new_attachment", "is_del", "update_userid", "update_time").Where("is_del = false").Order("new_time DESC").Limit(4).Find(&news).Error
	return err, news
}