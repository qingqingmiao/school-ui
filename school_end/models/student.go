package models

import (
	"chinasoccer/pkg/gredis"
	"encoding/json"
)

type Student struct {
	SID int `gorm:"column:SID"`
	Sname string `gorm:"column:sname"`
	Sage string `gorm:"column:sage"`
	Memo string `gorm:"column:memo;type(json)" `
	Memos MemoInfo `gorm:"-"`
	Sex bool `gorm:"column:sex"`
	Chengji string `gorm:"chengji"`

}

type MemoInfo struct {
	Guid string `json:"guid"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
}

var cache_key = "11111"

func GetAll() (error, []Student) {
	var stus []Student
	//exist,err:=gredis.RedisConn.Exists(cache_key)
	//if err != nil{
	//	logger.Error(err.Error())
	//}
	//if exist {
	//	gredis.RedisConn.GetObject(cache_key,&stus)
	//
	//}

	err := db.Table("student").Find(&stus).Error
	if len(stus) != 0{
		for index,each := range stus{
			json.Unmarshal([]byte(each.Memo),&each.Memos)
			each.Memo = ""
			stus[index] = each
		}
		gredis.RedisConn.Set(cache_key,stus,0)
	}
	return err, stus
}

func (s *Student) UpdateStudent()  {
	db.Table("student").Where("guid = ?",1).Update("name","123123213213213")
}
