package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type trainplan struct {
	Trainplan_id int `gorm:"column:trainplan_id"`
	Trainstart_time string `gorm:"column:trainstart_time"`
	Trainsend_time string `gorm:"column:trainend_time"`
	Train_position string `gorm:"column:train_position"`
	Train_name string `gorm:"column:train_name"`
	Train_playerlist postgres.Jsonb `gorm:"column:train_playerlist"`
	Train_content postgres.Jsonb `gorm:"column:train_content"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type traincontent struct {
	Train_content postgres.Jsonb `gorm:"column:train_content"`
}
func GetAllTrainplan(playerParam map[string]interface{}) (error, []trainplan, int64) {
	var trainplans []trainplan
	pageSize := playerParam["limit"].(int)
	page := playerParam["page"].(int)
	var total int64
	err := db.Table("tb_trainplan").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&trainplans).Error

	return err, trainplans, total
}

func GetTrainplanByID(id int) (error, postgres.Jsonb) {
	var trainplandata trainplan
	err := db.Table("tb_trainplan").Where("trainplan_id = ?", id).First(&trainplandata).Error
	return err, trainplandata.Train_playerlist
}

func AddTrainplan(a *gin.Context) (error, trainplan) {
	var trainplandata trainplan
	a.ShouldBind(&trainplandata)
	err := db.Table("tb_trainplan").Create(&trainplandata).Error
	return err, trainplandata
}

func AddSchedule(a *gin.Context, trainplan_id int) (error, traincontent) {
	//trainplan_id := a.Params.ByName("trainplan_id")
	var traincontentdata traincontent
	a.BindJSON(&traincontentdata)
	temp, _ := json.Marshal(traincontentdata.Train_content)
	temp2 := map[string]interface{}{}
	json.Unmarshal(temp, &temp2)
	var traincontentdatatemp traincontent
	err := db.Table("tb_trainplan").Select("train_content").Where("trainplan_id = ? and is_del = false", trainplan_id).Find(&traincontentdatatemp).Error
	temp3, _ := json.Marshal(traincontentdatatemp.Train_content)
	temp4 := map[string]interface{}{}
	json.Unmarshal(temp3, &temp4)
	for k, v := range temp2 {
		temp4[k] = v
	}
	err = db.Table("tb_trainplan").Where("trainplan_id = ?", trainplan_id).Update("train_content", temp4).Error
	return err, traincontentdatatemp
}

func DelSchedule(a *gin.Context, trainplan_id int) (error, traincontent) {
	//trainplan_id := a.Params.ByName("trainplan_id")
	var traincontentdata traincontent
	a.BindJSON(&traincontentdata)
	temp, _ := json.Marshal(traincontentdata.Train_content)
	temp2 := map[string]interface{}{}
	json.Unmarshal(temp, &temp2)
	var traincontentdatatemp traincontent
	err := db.Table("tb_trainplan").Select("train_content").Where("trainplan_id = ? and is_del = false", trainplan_id).Find(&traincontentdatatemp).Error
	temp3, _ := json.Marshal(traincontentdatatemp.Train_content)
	temp4 := map[string]interface{}{}
	json.Unmarshal(temp3, &temp4)
	var temp2key =""
	for k := range temp2 {
		temp2key = k
	}
	for index := 0; index < len(temp4); index++{
		for k :=range temp4{
			if k == temp2key {
				delete(temp4, k)
			}
		}
	}
	//fmt.Print(temp2key)
	err = db.Table("tb_trainplan").Where("trainplan_id = ?", trainplan_id).Update("train_content", temp4).Error
	return err, traincontentdatatemp
}