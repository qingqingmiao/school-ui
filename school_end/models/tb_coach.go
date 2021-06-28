package models

import (
	"github.com/gin-gonic/gin"
)

type coach struct {
	Id int `gorm:"column:id"`
	Coach_name string `gorm:"column:coach_name"`
	Coach_club string `gorm:"column:coach_club"`
}

func GetAllCoach() (error, []coach) {
	var coachs []coach
	//exist,err:=gredis.RedisConn.Exists(cache_key)
	//if err != nil{
	//	logger.Error(err.Error())
	//}
	//if exist {
	//	gredis.RedisConn.GetObject(cache_key,&stus)
	//
	//}
	err := db.Table("tb_coach").Find(&coachs).Error
	if len(coachs) != 0{
		for index,each := range coachs{
			//json.Unmarshal([]byte(each.Memo),&each.Memos)
			//each.Memo = ""
			//stus[index] = each
			coachs[index] = each
		}
		//gredis.RedisConn.Set(cache_key,stus,0)
	}

	return err, coachs
}

func AddCoach(a *gin.Context) (error, coach) {
	var coachdata coach
	a.BindJSON(&coachdata)
	err := db.Table("tb_coach").Create(&coachdata).Error
	return err, coachdata
}

func UpdateCoach(a *gin.Context) (error, coach) {
	var coachdata coach
	a.BindJSON(&coachdata)
	var id = coachdata.Id
	err := db.Table("tb_coach").Where("Id = ?", id).Updates(&coachdata).Error
	return err, coachdata
}

func DelCoach(a *gin.Context) (error, coach) {
	id := a.Params.ByName("id")
	var coachdata coach
	err := db.Table("tb_coach").Where("Id = ?", id).Delete(&coachdata).Error
	return err, coachdata
}
