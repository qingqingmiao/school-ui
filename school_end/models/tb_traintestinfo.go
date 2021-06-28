package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type traintestinfo struct {
	Traintest_id int `gorm:"column:traintest_id"`
	Trainplan_id string `gorm:"column:trainplan_id"`
	Player_id string `gorm:"column:player_id"`
	Height int `gorm:"column:height"`
	Weight int `gorm:"column:weight"`
	Left_leg string `gorm:"column:left_leg"`
	Right_leg string `gorm:"column:right_leg"`
	Bodyfat string `gorm:"column:bodyfat"`
	Flexibility string `gorm:"column:flexibility"`
	Verticaljump string `gorm:"column:verticaljump"`
	Thirtymeter_sprint string `gorm:"column:thirtymeter_sprint"`
	Arrowrun string `gorm:"column:arrowrun"`
	Yoyoir2 string `gorm:"column:yoyoir2"`
	Push_up string `gorm:"column:push_up"`
	Pull_up string `gorm:"column:pull_up"`
	Test_position string `gorm:"column:test_position"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type traintestinfoTemp struct {
	Traintest_id int `gorm:"column:traintest_id"`
	Trainplan_id string `gorm:"column:trainplan_id"`
	Player_id string `gorm:"column:player_id"`
	Chinese_name string `gorm:"column:chinese_name"`
	Height int `gorm:"column:height"`
	Weight int `gorm:"column:weight"`
	Left_leg string `gorm:"column:left_leg"`
	Right_leg string `gorm:"column:right_leg"`
	Bodyfat string `gorm:"column:bodyfat"`
	Flexibility string `gorm:"column:flexibility"`
	Verticaljump string `gorm:"column:verticaljump"`
	Thirtymeter_sprint string `gorm:"column:thirtymeter_sprint"`
	Arrowrun string `gorm:"column:arrowrun"`
	Yoyoir2 string `gorm:"column:yoyoir2"`
	Push_up string `gorm:"column:push_up"`
	Pull_up string `gorm:"column:pull_up"`
	Test_position string `gorm:"column:test_position"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

func GetTraintestinfoByID(trainplan_id string) (error, []traintestinfoTemp) {
	var traintestinfoTempdata []traintestinfoTemp
	err := db.Table("tb_traintestinfo").Select("tb_traintestinfo.traintest_id, tb_traintestinfo.trainplan_id, tb_traintestinfo.player_id, tb_player.chinese_name, " +
		"tb_traintestinfo.height, tb_traintestinfo.weight, tb_traintestinfo.left_leg, tb_traintestinfo.right_leg, tb_traintestinfo.bodyfat, tb_traintestinfo.flexibility, " +
		"tb_traintestinfo.verticaljump, tb_traintestinfo.thirtymeter_sprint, tb_traintestinfo.arrowrun, tb_traintestinfo.yoyoir2, tb_traintestinfo.push_up, " +
		"tb_traintestinfo.pull_up, tb_traintestinfo.test_position, tb_traintestinfo.is_del, tb_traintestinfo.update_userid, tb_traintestinfo.update_time").
		Where("trainplan_id = ?", trainplan_id).Joins("left join tb_player on tb_player.player_id = tb_traintestinfo.player_id").Find(&traintestinfoTempdata).Error
	return err, traintestinfoTempdata
}

func AddTraintestinfo(a *gin.Context) (error, traintestinfo) {
	var traintestinfodata traintestinfo
	a.BindJSON(&traintestinfodata)
	traintestinfodata.Is_del = false
	traintestinfodata.Update_userid = "0"
	traintestinfodata.Update_time = time.Now()
	err := db.Table("tb_traintestinfo").Select("trainplan_id","player_id","height","weight","left_leg","right_leg","bodyfat","flexibility","verticaljump",
		"thirtymeter_sprint","arrowrun","yoyoir2","push_up","pull_up","test_position","is_del","update_userid","update_time").Create(&traintestinfodata).Error
	return err, traintestinfodata
}

func UpdateTraintestinfo(a *gin.Context) (error, traintestinfo) {
	var traintestinfodata traintestinfo
	a.BindJSON(&traintestinfodata)
	traintestinfodata.Update_time = time.Now()
	var traintest_id = traintestinfodata.Traintest_id
	err := db.Table("tb_traintestinfo").Where("traintest_id = ?", traintest_id).Updates(&traintestinfodata).Error
	return err, traintestinfodata
}