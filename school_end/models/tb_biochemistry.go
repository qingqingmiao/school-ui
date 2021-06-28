package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type biochemistry struct {
	Biochemistry_id int `gorm:"column:biochemistry_id"`
	Trainplan_id string `gorm:"column:trainplan_id"`
	Player_id string `gorm:"column:player_id"`
	Player_code string `gorm:"column:player_code"`
	Player_name string `gorm:"column:player_name"`
	Urine_gravity string `gorm:"column:urine_gravity"`
	Ph string `gorm:"column:ph"`
	Urobilinogen string `gorm:"column:urobilinogen"`
	Urine_protein string `gorm:"column:urine_protein"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

func GetBiochemistryByID(trainplan_id string) (error, []biochemistry) {
	var biochemistrydata []biochemistry
	err := db.Table("tb_biochemistry").Where("trainplan_id = ?", trainplan_id).Find(&biochemistrydata).Error
	return err, biochemistrydata
}

func AddBiochemistry(a *gin.Context) (error, biochemistry) {
	var biochemistrydata biochemistry
	a.BindJSON(&biochemistrydata)
	biochemistrydata.Update_time = time.Now()
	err := db.Table("tb_biochemistry").Select("trainplan_id","player_id","player_code","player_name","urine_gravity","ph","urobilinogen",
		"urine_protein","is_del","update_userid","update_time").Create(&biochemistrydata).Error
	return err, biochemistrydata
}

func UpdateBiochemistry(a *gin.Context) (error, biochemistry) {
	var biochemistrydata biochemistry
	a.BindJSON(&biochemistrydata)
	biochemistrydata.Update_time = time.Now()
	var biochemistry_id = biochemistrydata.Biochemistry_id
	err := db.Table("tb_biochemistry").Where("biochemistry_id = ?", biochemistry_id).Updates(&biochemistrydata).Error
	return err, biochemistrydata
}