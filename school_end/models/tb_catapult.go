package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type catapult struct {
	Catapult_id int `gorm:"column:catapult_id"`
	Trainplan_id string `gorm:"column:trainplan_id"`
	Player_id string `gorm:"column:player_id"`
	Player_name string `gorm:"column:player_name"`
	Period_name string `gorm:"column:period_name"`
	Period_number string `gorm:"column:period_number"`
	Position_name string `gorm:"column:position_name"`
	Unix_start_time string `gorm:"column:unix_start_time"`
	Unix_end_time string `gorm:"column:unix_end_time"`
	Total_time  string `gorm:"column:total_time"`
	Total_distance  string `gorm:"column:total_distance"`
	Total_player_load  string `gorm:"column:total_player_load"`
	Load_intensity  string `gorm:"column:load_intensity"`
	Load_permeter  string `gorm:"column:load_permeter"`
	Meter_permin  string `gorm:"column:meter_permin"`
	Max_heartrate  string `gorm:"column:max_heartrate"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

func GetCatapultByID(trainplan_id string) (error, []catapult) {
	var catapultdata []catapult
	err := db.Table("tb_catapult").Where("trainplan_id = ?", trainplan_id).Find(&catapultdata).Error
	return err, catapultdata
}

func AddCatapult(a *gin.Context) (error, catapult) {
	var catapultdata catapult
	a.BindJSON(&catapultdata)
	catapultdata.Update_time = time.Now()
	err := db.Table("tb_catapult").Select("trainplan_id","player_id","player_name","period_name","period_number","position_name","unix_start_time","unix_end_time",
		"total_time","total_distance","total_player_load","load_intensity","load_permeter","meter_permin","max_heartrate","is_del","update_userid","update_time").Create(&catapultdata).Error
	return err, catapultdata
}

func UpdateCatapult(a *gin.Context) (error, catapult) {
	var catapultdata catapult
	a.BindJSON(&catapultdata)
	catapultdata.Update_time = time.Now()
	var catapult_id = catapultdata.Catapult_id
	err := db.Table("tb_catapult").Where("catapult_id = ?", catapult_id).Updates(&catapultdata).Error
	return err, catapultdata
}