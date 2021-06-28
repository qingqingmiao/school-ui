package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

// 伤病中心页面的视图取名为injury_view
type injury_view struct {
	Injury_id int `gorm:"column:injury_id"`
	Player_id int `gorm:"column:player_id"`
	Chinese_name string `gorm:"column:chinese_name"`
	Injury_time string `gorm:"column:injury_time"`
	Type_name string `gorm:"column:type_name"`
	Injury_position string `gorm:"column:injury_position"`
	Injury_level string `gorm:"column:injury_level"`
	Injury_location string `gorm:"column:injury_location"`
	Injury_scene string `gorm:"column:injury_scene"`
	Nation_doctorname string `gorm:"column:nation_doctorname"`
	Nation_doctorid int `gorm:"column:nation_doctorid"`
	Club_doctorname string `gorm:"column:club_doctorname"`
	Club_doctorid int `gorm:"column:club_doctorid"`
	Recovery_time string `gorm:"column:recovery_time"`
	Injury_method string `gorm:"column:injury_method"`
}

type user_doctor struct {
	User_id int `gorm:"column:user_id"`
	User_type string `gorm:"column:user_type"`
	Account string `gorm:"column:account"`
	Username string `gorm:"column:username"`
}

type injury struct {
	Injury_id int `gorm:"column:injury_id"`
	Player_id int `gorm:"column:player_id"`
	Injury_time string `gorm:"column:injury_time"`
	Injury_typeid int `gorm:"column:injury_typeid"`
	Injury_position string `gorm:"column:injury_position"`
	Injury_location string `gorm:"column:injury_location"`
	Injury_level string `gorm:"column:injury_level"`
	Injury_scene string `gorm:"column:injury_scene"`
	Injury_method string `gorm:"column:injury_method"`
	Recovery_time string `gorm:"column:recovery_time"`
	Nation_doctorid int `gorm:"column:nation_doctorid"`
	Club_doctorid int `gorm:"column:club_doctorid"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type common struct{
	Group_id int `gorm:"column:group_id"`
	Group_name string `gorm:"column:group_name"`
	Type_id int `gorm:"column:type_id"`
	Type_name string `gorm:"column:type_name"`
	Is_del bool `gorm:"column:is_del"`
}

func GetAllInjury(injuryParam map[string]interface{}) (error, []injury_view, int64) {
	var injury_views []injury_view
	page := injuryParam["page"].(int)
	pageSize := injuryParam["limit"].(int)
	searchText := injuryParam["searchText"].(string)
	var total int64
	err := db.Table("injuryList").Where("chinese_name like ? and is_del = false", "%"+searchText+"%").Order("injury_id DESC").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&injury_views).Error
	return err, injury_views, total
}

func GetInjuryByID(player_id string) (error, []injury_view) {
	var injury_viewdata []injury_view
	err := db.Table("injuryList").Select("injury_id","player_id", "chinese_name", "injury_time", "type_name", "injury_position", "injury_level", "injury_location", "injury_scene").Where("player_id = ? and is_del = false", player_id).Find(&injury_viewdata).Error
	return err, injury_viewdata
}

func AddInjury(a *gin.Context) (error, injury_view) {
	var injury_viewdata injury_view
	a.BindJSON(&injury_viewdata)
	//查看该伤病类型有没有，没有就为-1，有就是非-1
	type_id := -1
	err := db.Table("tb_common").Select("type_id").Where("type_name = ?", injury_viewdata.Type_name).Find(&type_id).Error
	if(type_id == -1) {
		var maxValue int
		err = db.Table("tb_common").Select("max(type_id)").Where("group_id = 2").Find(&maxValue).Error
		var commondata common
		commondata.Group_id = 2
		commondata.Group_name = "伤病类型"
		commondata.Type_id = maxValue + 1
		commondata.Type_name = injury_viewdata.Type_name
		commondata.Is_del = false
		err = db.Table("tb_common").Create(&commondata).Error
		type_id = commondata.Type_id
	}
	var injurydata injury
	injurydata.Player_id = injury_viewdata.Player_id
	injurydata.Injury_time = injury_viewdata.Injury_time
	injurydata.Injury_typeid = type_id
	injurydata.Injury_position = injury_viewdata.Injury_position
	injurydata.Injury_location = injury_viewdata.Injury_location
	injurydata.Injury_level = injury_viewdata.Injury_level
	injurydata.Injury_scene = injury_viewdata.Injury_scene
	injurydata.Nation_doctorid = injury_viewdata.Nation_doctorid
	injurydata.Club_doctorid = injury_viewdata.Club_doctorid
	injurydata.Recovery_time = injury_viewdata.Recovery_time
	injurydata.Injury_method = injury_viewdata.Injury_method
	injurydata.Is_del = false
	injurydata.Update_userid = "0"
	injurydata.Update_time = time.Now()
	err = db.Table("tb_injury").Select("player_id", "injury_time", "injury_typeid", "injury_position", "injury_location", "injury_level", "injury_scene",
		"injury_method", "recovery_time", "nation_doctorid", "club_doctorid", "is_del", "update_userid", "update_time").Create(&injurydata).Error
	return err, injury_viewdata
}

func GetDoctor() (error, []user_doctor) {
	var user_doctors []user_doctor
	err := db.Table("tb_user").Where("user_type = '7' and is_del = false").Order("user_id ASC").Find(&user_doctors).Error
	return err, user_doctors
}

func UpdateInjury(a *gin.Context) (error, injury_view) {
	var injury_viewdata injury_view
	a.BindJSON(&injury_viewdata)
	//查看该伤病类型有没有，没有就为-1，有就是非-1
	type_id := -1
	err := db.Table("tb_common").Select("type_id").Where("type_name = ?", injury_viewdata.Type_name).Find(&type_id).Error
	if(type_id == -1) {
		var maxValue int
		err = db.Table("tb_common").Select("max(type_id)").Where("group_id = 2").Find(&maxValue).Error
		var commondata common
		commondata.Group_id = 2
		commondata.Group_name = "伤病类型"
		commondata.Type_id = maxValue + 1
		commondata.Type_name = injury_viewdata.Type_name
		commondata.Is_del = false
		err = db.Table("tb_common").Create(&commondata).Error
		type_id = commondata.Type_id
	}
	var injurydata injury
	injurydata.Player_id = injury_viewdata.Player_id
	injurydata.Injury_time = injury_viewdata.Injury_time
	injurydata.Injury_typeid = type_id
	injurydata.Injury_position = injury_viewdata.Injury_position
	injurydata.Injury_location = injury_viewdata.Injury_location
	injurydata.Injury_level = injury_viewdata.Injury_level
	injurydata.Injury_scene = injury_viewdata.Injury_scene
	injurydata.Nation_doctorid = injury_viewdata.Nation_doctorid
	injurydata.Club_doctorid = injury_viewdata.Club_doctorid
	injurydata.Recovery_time = injury_viewdata.Recovery_time
	injurydata.Injury_method = injury_viewdata.Injury_method
	injurydata.Is_del = false
	injurydata.Update_userid = "0"
	injurydata.Update_time = time.Now()
	err = db.Table("tb_injury").Select("player_id", "injury_time", "injury_typeid", "injury_position", "injury_location", "injury_level", "injury_scene",
		"injury_method", "recovery_time", "nation_doctorid", "club_doctorid", "is_del", "update_userid", "update_time").Where("injury_id = ?", injury_viewdata.Injury_id).Updates(&injurydata).Error
	return err, injury_viewdata
}

func DelInjury(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_injury").Where("injury_id = ?", id).Update("is_del", true).Error
	return err
}
