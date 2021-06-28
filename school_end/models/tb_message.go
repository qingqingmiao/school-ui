package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type tb_message struct {
	Message_id int `gorm:"column:message_id"`
	Message_title string `gorm:"column:message_title"`
	Message_time time.Time `gorm:"column:message_time"`
	Message_content string `gorm:"column:message_content"`
	Message_attachment string `gorm:"column:message_attachment"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}


func AddMessage(a *gin.Context) (error, tb_message) {
	var messagedata tb_message
	a.ShouldBind(&messagedata)
	messagedata.Update_userid = "0"
	messagedata.Is_del = false
	messagedata.Update_time = time.Now()
	//err := db.Table("tb_message").Select("message_title","message_time","message_content","is_del","update_userid","update_time").Create(&messagedata).Error
	err := db.Table("tb_message").Select("message_title","message_time", "message_content", "message_attachment", "is_del", "update_userid", "update_time").Create(&messagedata).Error
	return err, messagedata
}

func GetAllMessage(messageParam map[string]interface{}) (error, []tb_message, int64) {
	var messages []tb_message
	page := messageParam["page"].(int)
	pageSize := messageParam["limit"].(int)
	searchText := messageParam["searchText"].(string)
	var total int64
	err := db.Table("tb_message").Where("message_title like ? and is_del = false", "%"+searchText+"%").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&messages).Error
	return err, messages, total
}

func DelMessage(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_message").Where("message_id = ?", id).Update("is_del", true).Error
	return err
}

func UpdateMessage(a *gin.Context) (error, tb_message) {
	var messagedata tb_message
	a.BindJSON(&messagedata)
	messagedata.Is_del = false
	messagedata.Update_userid = "0"
	messagedata.Update_time = time.Now()
	err := db.Table("tb_message").Select("message_title", "message_time", "message_content", "message_attachment", "is_del", "update_userid", "update_time").Where("message_id = ?", messagedata.Message_id).Updates(&messagedata).Error
	return err, messagedata
}

func GetFourMessage() (error, []tb_message) {
	var messages []tb_message
	//page := messageParam["page"].(int)
	//pageSize := messageParam["limit"].(int)
	//var total int64
	//err := db.Table("tb_message").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&messages).Error
	err := db.Table("tb_message").Select("message_id", "message_title", "message_time", "message_content", "message_attachment", "is_del", "update_userid", "update_time").Where("is_del = false").Order("message_time DESC").Limit(4).Find(&messages).Error
	return err, messages
}