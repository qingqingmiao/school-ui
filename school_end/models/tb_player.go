package models

import (
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/gredis"
	"chinasoccer/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type player struct {
	Player_id int `gorm:"column:player_id"`
	Chinese_name string `gorm:"column:chinese_name"`
	English_name string `gorm:"column:english_name"`
	Id_code string `gorm:"column:id_code"`
	Birthday string `gorm:"column:birthday"`
	Age string `gorm:"column:age"`
	Height string `gorm:"column:height"`
	Weight string `gorm:"column:weight"`
	Position string `gorm:"column:position"`
	Team_name string `gorm:"column:team_name"`
	State string `gorm:"column:state"`
	Technical_level string `gorm:"column:technical_level"`
	Thumbnail_portrait string `gorm:"column:thumbnail_portrait"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type team struct {
	Team_id int `gorm:"column:team_id"`
	Team_name string `gorm:"column:team_name"`
	Club_name string `gorm:"column:club_name"`
	Founded_time time.Time `gorm:"column:founded_time"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type teamplayer struct {
	Player_id int `gorm:"column:player_id"`
	Team_id int `gorm:"column:team_id"`
	Position string `gorm:"column:position"`
	Is_ok bool `gorm:"column:is_ok"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type LocalDate time.Time
// MarshalJSON satify the json marshal interface
func (l LocalDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(l).Format("2006-01-02"))
	return []byte(stamp), nil
}

func GetAllPlayer(playerParam map[string]interface{}) (error, []player, int64) {
	var players []player
	page := playerParam["page"].(int)
	pageSize := playerParam["limit"].(int)
	initialLowerCase := strings.ToLower(playerParam["initials"].(string))
	searchText := playerParam["searchText"].(string)
	var total int64
	err := db.Table("playerList").Select("player_id, chinese_name, english_name, id_code, birthday, age, height, weight, position, team_name, state, technical_level, " +
		"thumbnail_portrait, is_del").Where("age >= ? and age <= ? and height >= ? and height <= ? and weight >= ? and weight <= ? and (english_name like ? or english_name like ? or ? = '') " +
		"and (team_name = ? or ? = '') and (position = ? or ? = '') and (state = ? or ? = '') and (chinese_name like ? or position like ? or team_name like ? or state like ? or Id_code like ?) and is_del = false",
		playerParam["ageGroup[]"].([]string)[0], playerParam["ageGroup[]"].([]string)[1], playerParam["heightGroup[]"].([]string)[0], playerParam["heightGroup[]"].([]string)[1],
		playerParam["weightGroup[]"].([]string)[0], playerParam["weightGroup[]"].([]string)[1], playerParam["initials"].(string)+"%", initialLowerCase+"%", playerParam["initials"].(string),
		playerParam["team_name"].(string), playerParam["team_name"].(string), playerParam["position"].(string), playerParam["position"].(string),
		playerParam["state"].(string), playerParam["state"].(string), "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%", searchText+"%").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&players).Error

	//err := db.Table("playerList").Where("age >= ? and age <= ? and height >= ? and height <= ? and " +
	//	"weight >= ? and weight <= ?",
	//	playerParam["ageGroup[]"].([]string)[0], playerParam["ageGroup[]"].([]string)[1], playerParam["heightGroup[]"].([]string)[0], playerParam["heightGroup[]"].([]string)[1],
	//	playerParam["weightGroup[]"].([]string)[0], playerParam["weightGroup[]"].([]string)[1]).Offset((page-1)*pageSize).Limit(pageSize).Find(&players).Error

	return err, players, total
}

func AddPlayer(a *gin.Context) (error, player) {
	var playerdata player
	a.ShouldBind(&playerdata)
	playerdata.Is_del = false
	playerdata.Update_userid = "0"
	playerdata.Update_time = time.Now()
	err := db.Table("tb_player").Select("chinese_name","birthday","state","height","weight","id_code","is_del","update_userid","update_time").Create(&playerdata).Error
	err =  db.Table("tb_player").Last(&playerdata).Error
	//判断Team_name是否在tb_team中有则返回team_id没有就插入并返回team_id
	var teamdata team
	var teamNameCnt int64
	err = db.Table("tb_team").Where("team_name = ?", playerdata.Team_name).Count(&teamNameCnt).Error
	if(teamNameCnt == 0) {
		teamdata.Team_name = playerdata.Team_name
		teamdata.Club_name = playerdata.Team_name
		timeTemp := "2021-05-17"
		var timeLayoutStr = "2006-01-02"
		st, _ := time.Parse(timeLayoutStr, timeTemp)
		teamdata.Founded_time = st
		teamdata.Is_del = false
		teamdata.Update_userid = "0"
		teamdata.Update_time = time.Now()
		err = db.Table("tb_team").Select("team_name","club_name","founded_time","is_del","update_userid","update_time").Create(&teamdata).Error
		err =  db.Table("tb_team").Last(&teamdata).Error
	} else {
		err = db.Table("tb_team").Where("team_name = ?", playerdata.Team_name).First(&teamdata).Error
	}
	//需要插入player_id、team_id、positon
	var teamplayerdata teamplayer
	teamplayerdata.Player_id = playerdata.Player_id
	teamplayerdata.Team_id = teamdata.Team_id
	teamplayerdata.Position = playerdata.Position
	teamplayerdata.Is_ok = false
	teamplayerdata.Update_userid = "0"
	teamplayerdata.Update_time = time.Now()
	err = db.Table("tb_teamplayer").Select("player_id","Team_id", "Position", "Is_ok", "Update_userid", "Update_time").Create(&teamplayerdata).Error
	return err, playerdata
}
func UpdatePlayer(a *gin.Context) (error, player) {
	var playerdata player
	a.ShouldBind(&playerdata)
	playerdata.Update_userid = "0"
	playerdata.Update_time = time.Now()
	err := db.Table("tb_player").Where("player_id = ?", playerdata.Player_id).Updates(&playerdata).Error
	err =  db.Table("tb_player").Last(&playerdata).Error
	//判断Team_name是否在tb_team中有则返回team_id没有就插入并返回team_id
	var teamdata team
	var teamId int
	err = db.Table("tb_team").Where("team_name = ?", playerdata.Team_name).First(&teamdata).Error
	teamId = teamdata.Team_id
	teamdata.Update_userid = "0"
	teamdata.Update_time = time.Now()
	err = db.Table("tb_team").Where("Team_id = ?", teamId).Updates(&teamdata).Error
	err = db.Table("tb_team").Where("team_name = ?", playerdata.Team_name).First(&teamdata).Error
	//需要插入player_id、team_id、positon
	var teamplayerdata teamplayer
	teamplayerdata.Player_id = playerdata.Player_id
	teamplayerdata.Team_id = teamdata.Team_id
	teamplayerdata.Position = playerdata.Position
	teamplayerdata.Update_userid = "0"
	teamplayerdata.Update_time = time.Now()
	err = db.Table("tb_teamplayer").Where("player_id = ?","Team_id = ?", playerdata.Player_id, teamdata.Team_id).Updates(&teamplayerdata).Error
	return err, playerdata
}

func DelPlayer(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_player").Where("player_id = ?", id).Update("is_del", true).Error
	return err
}

type playerKeyData struct {
	Player_id int `gorm:"column:player_id"`
	Chinese_name string `gorm:"column:chinese_name"`
	English_name string `gorm:"column:english_name"`
}

func GetAllPlayerKeyData() (error , []playerKeyData) {
	var playerKeyDatas []playerKeyData
	exist, err := gredis.RedisConn.Exists(e.AllPlayerRedisKey)
	if err != nil {
		logger.Error(err.Error())
	}
	if exist {
		gredis.RedisConn.GetObject(e.AllPlayerRedisKey, &playerKeyDatas)
		return err, playerKeyDatas
	}
	err = db.Table("tb_player").Select("player_id, chinese_name, english_name").Where("is_del = false").Order("player_id ASC").Find(&playerKeyDatas).Error
	if len(playerKeyDatas) != 0 {
		gredis.RedisConn.Set(e.AllPlayerRedisKey, playerKeyDatas, 0)
	}
	return err, playerKeyDatas
}