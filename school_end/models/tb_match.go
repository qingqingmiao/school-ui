package models

import (
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/gredis"
	"chinasoccer/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type match_view struct {
	Match_id int `gorm:"column:match_id"`
	Type_id int `gorm:"column:type_id"`
	Type_name string `gorm:"column:type_name"`
	Season string `gorm:"column:season"`
	Match_time string `gorm:"column:match_time"`
	Hometeam_id int `gorm:"column:hometeam_id"`
	Hometeam_name string `gorm:"column:hometeam_name"`
	Hometeam_score int `gorm:"column:hometeam_score"`
	Visitingteam_id int `gorm:"column:visitingteam_id"`
	Visitingteam_name string `gorm:"column:visitingteam_name"`
	Visitingteam_score int `gorm:"column:visitingteam_score"`
	Hometeam_formation string `gorm:"column:hometeam_formation"`
	Visitingteam_formation string `gorm:"column:visitingteam_formation"`
}

type teaminfo struct {
	Team_id int `gorm:"column:team_id"`
	Team_name string `gorm:"column:team_name"`
}

type playerinfo struct {
	Team_id int `gorm:"column:team_id"`
	Player_id int `gorm:"column:player_id"`
	Uniform_number string `gorm:"column:uniform_number"`
	Chinese_name string `gorm:"column:chinese_name"`
}

type match struct {
	Match_id int `gorm:"column:match_id"`
	Leaguematch_type int `gorm:"column:leaguematch_type"`
	Season string `gorm:"column:season"`
	Match_time string `gorm:"column:match_time"`
	Hometeam_id int `gorm:"column:hometeam_id"`
	Hometeam_formation string `gorm:"column:hometeam_formation"`
	Hometeam_score int `gorm:"column:hometeam_score"`
	Hformation_data postgres.Jsonb `gorm:"column:hformation_data"`
	Visitingteam_id int `gorm:"column:visitingteam_id"`
	Visitingteam_formation string `gorm:"column:visitingteam_formation"`
	Visitingteam_score int `gorm:"column:visitingteam_score"`
	Vformation_data postgres.Jsonb `gorm:"column:vformation_data"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type teamformation_data struct {
	Match_id int `gorm:"column:match_id"`
	Hformation_data postgres.Jsonb `gorm:"column:hformation_data"`
	Vformation_data postgres.Jsonb `gorm:"column:vformation_data"`
}

type season struct {
	Season string `gorm:"column:season"`
}

type achieveinfo_view struct {
	Start_sum int
	Match_sum int
	Goal_sum int
	Assist_sum int
	Yellowcard_sum int
	Redcard_sum int
}

type playerinjury_view struct {
	Injury_time string
	Type_name string
	Injury_position string
	Injury_level string
	Injury_location string
	Injury_scene string
	Season string
}

type transferinfo_view struct {
	Season string
	Type_name string
	Shoottarget int
	Sucsnatch int
	Redcard int
	Yellowcard int
	Touchball int
}

func GetAllMatch(matchParam map[string]interface{}) (error, []match_view, int64) {
	var matches []match_view
	page := matchParam["page"].(int)
	pageSize := matchParam["limit"].(int)
	searchText := matchParam["searchText"].(string)
	var total int64
	err := db.Table("matchList").Where("type_name like ? and is_del = false", "%"+searchText+"%").Order("match_id DESC").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&matches).Error
	return err, matches, total
}

func AddMatch(a *gin.Context) (error, match_view) {
	var match_viewdata match_view
	a.BindJSON(&match_viewdata)
	var matchdata match
	matchdata.Leaguematch_type = match_viewdata.Type_id
	matchdata.Season = match_viewdata.Season
	matchdata.Match_time = match_viewdata.Match_time
	matchdata.Hometeam_id = match_viewdata.Hometeam_id
	matchdata.Hometeam_formation = match_viewdata.Hometeam_formation
	matchdata.Hometeam_score = match_viewdata.Hometeam_score
	matchdata.Visitingteam_id = match_viewdata.Visitingteam_id
	matchdata.Visitingteam_formation = match_viewdata.Visitingteam_formation
	matchdata.Visitingteam_score = match_viewdata.Visitingteam_score
	matchdata.Is_del = false
	matchdata.Update_userid = "0"
	matchdata.Update_time = time.Now()
	err := db.Table("tb_match").Select("leaguematch_type", "season", "match_time", "hometeam_id", "hometeam_formation", "hometeam_score", "visitingteam_id",
		"visitingteam_formation", "visitingteam_score", "is_del", "update_userid", "update_time").Create(&matchdata).Error
	return err, match_viewdata
}

func GetCompetition() (error, []common) {
	var commondata []common
	err := db.Table("tb_common").Where("group_id = 3").Order("type_id ASC").Find(&commondata).Error
	return err, commondata
}

func GetFootballTeam() (error, []teaminfo) {
	var teaminfodata []teaminfo
	exist, err := gredis.RedisConn.Exists(e.AllTeamRedisKey)
	if err != nil {
		logger.Error(err.Error())
	}
	if exist {
		gredis.RedisConn.GetObject(e.AllTeamRedisKey, &teaminfodata)
		return err, teaminfodata
	}
	err = db.Table("tb_team").Select("team_id, team_name").Where("is_del = false").Order("team_id ASC").Find(&teaminfodata).Error
	if len(teaminfodata) != 0 {
		gredis.RedisConn.Set(e.AllTeamRedisKey, teaminfodata, 0)
	}
	return err, teaminfodata
}

func UpdateMatch(a *gin.Context) (error, match_view) {
	var match_viewdata match_view
	a.BindJSON(&match_viewdata)
	var matchdata match
	matchdata.Leaguematch_type = match_viewdata.Type_id
	matchdata.Season = match_viewdata.Season
	matchdata.Match_time = match_viewdata.Match_time
	matchdata.Hometeam_id = match_viewdata.Hometeam_id
	matchdata.Hometeam_formation = match_viewdata.Hometeam_formation
	matchdata.Hometeam_score = match_viewdata.Hometeam_score
	matchdata.Visitingteam_id = match_viewdata.Visitingteam_id
	matchdata.Visitingteam_formation = match_viewdata.Visitingteam_formation
	matchdata.Visitingteam_score = match_viewdata.Visitingteam_score
	matchdata.Is_del = false
	matchdata.Update_userid = "0"
	matchdata.Update_time = time.Now()
	err := db.Table("tb_match").Select("leaguematch_type", "season", "match_time", "hometeam_id", "hometeam_formation", "hometeam_score", "visitingteam_id",
		"visitingteam_formation", "visitingteam_score", "is_del", "update_userid", "update_time").Where("match_id = ?", match_viewdata.Match_id).Updates(&matchdata).Error
	return err, match_viewdata
}

func DelMatch(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_match").Where("match_id = ?", id).Update("is_del", true).Error
	return err
}

func GetTeamPlayerByID(a *gin.Context) (error, []playerinfo) {
	var playerinfodata []playerinfo
	id := a.Params.ByName("id")
	err := db.Table("tb_teamplayer").Select("tb_teamplayer.team_id", "tb_player.player_id", "tb_teamplayer.uniform_number", "tb_player.chinese_name").
		Where("team_id = ? and is_ok = false", id).Joins("left join tb_player on tb_player.player_id = tb_teamplayer.player_id").Order("player_id ASC").
		Find(&playerinfodata).Error
	return err, playerinfodata
}

func AddTeamFormation_data(a *gin.Context) (error, teamformation_data) {
	var teamformation_datadata teamformation_data
	a.BindJSON(&teamformation_datadata)
	team_type := a.Query("team_type")
	var err error
	if team_type ==  "home" {
		err = db.Exec("UPDATE tb_match SET hformation_data=?,is_del=?,update_userid=?,update_time=? WHERE match_id=?",
			teamformation_datadata.Hformation_data, false, "0", time.Now(), teamformation_datadata.Match_id).Error
	} else {
		err = db.Exec("UPDATE tb_match SET vformation_data=?,is_del=?,update_userid=?,update_time=? WHERE match_id=?",
			teamformation_datadata.Vformation_data, false, "0", time.Now(), teamformation_datadata.Match_id).Error
	}

	return err, teamformation_datadata
}

func GetMatchSeason() (error, []season) {
	var seasondata []season
	err := db.Table("tb_match").Distinct("season").Select("season").Where("season <> '' and is_del = false").Order("season DESC").Find(&seasondata).Error
	return err, seasondata
}

func GetPlayerdetailachieveinfo(searchParam map[string]interface{}) (error, achieveinfo_view) {
	var achieveinfo_viewdata []achieveinfo_view
	player_id := searchParam["player_id"].(int)
	season := searchParam["season"].(string)
	leaguematch_type := searchParam["leaguematch_type"].(int)
	is_hometeam := searchParam["is_hometeam"].(string)
	is_victory := searchParam["is_victory"].(string)
	com_number := searchParam["com_number"].(int)
	var err error
	//is_hometeam为1时筛选为主队的数据，为0时筛选为客队的数据，为空时筛选主队和客队的数据
	//is_victory为1时筛选胜的数据，为0时筛选负的数据，为2筛选平局的数据，为空时筛选所有数据
	if is_hometeam == "" && is_victory == "" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "" && is_victory == "0" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and ((tb_match.hometeam_id = ? and tb_match.hometeam_score < tb_match.visitingteam_score) or (tb_match.visitingteam_id = ? and tb_match.hometeam_score > tb_match.visitingteam_score)) and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "" && is_victory == "1" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and ((tb_match.hometeam_id = ? and tb_match.hometeam_score > tb_match.visitingteam_score) or (tb_match.visitingteam_id = ? and tb_match.hometeam_score < tb_match.visitingteam_score)) and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "" && is_victory == "2" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and ((tb_match.hometeam_id = ? and tb_match.hometeam_score = tb_match.visitingteam_score) or (tb_match.visitingteam_id = ? and tb_match.hometeam_score = tb_match.visitingteam_score)) and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "0" && is_victory == "" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.visitingteam_id = ? and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "0" && is_victory == "0" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.visitingteam_id = ? and tb_match.hometeam_score > tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "0" && is_victory == "1" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.visitingteam_id = ? and tb_match.hometeam_score < tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "0" && is_victory == "2" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.visitingteam_id = ? and tb_match.hometeam_score = tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "1" && is_victory == "" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.hometeam_id = ? and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "1" && is_victory == "0" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.hometeam_id = ? and tb_match.hometeam_score < tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "1" && is_victory == "1" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.hometeam_id = ? and tb_match.hometeam_score > tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	} else if is_hometeam == "1" && is_victory == "2" {
		err = db.Table("tb_playermatch").Select("count(tb_playermatch.is_start = '1' or null) as start_sum, count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and (tb_match.season = ? or ? = '') and (tb_match.leaguematch_type = ? or ? = -1) and tb_match.hometeam_id = ? and tb_match.hometeam_score = tb_match.visitingteam_score and tb_playermatch.is_del = false and tb_match.is_del = false", player_id, season, season, leaguematch_type, leaguematch_type, player_id).Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").Group("tb_match.match_time").Order("tb_match.match_time DESC").Limit(com_number).Find(&achieveinfo_viewdata).Error
	}
	var achieveinfodata achieveinfo_view
	for i, _ := range achieveinfo_viewdata {
		if achieveinfo_viewdata[i].Start_sum != 0 {
			achieveinfodata.Start_sum += 1
		}
		achieveinfodata.Match_sum += 1
		if achieveinfo_viewdata[i].Goal_sum != 0 {
			achieveinfodata.Goal_sum += achieveinfo_viewdata[i].Goal_sum
		}
		if achieveinfo_viewdata[i].Assist_sum != 0 {
			achieveinfodata.Assist_sum += achieveinfo_viewdata[i].Assist_sum
		}
		if achieveinfo_viewdata[i].Yellowcard_sum != 0 {
			achieveinfodata.Yellowcard_sum += achieveinfo_viewdata[i].Yellowcard_sum
		}
		if achieveinfo_viewdata[i].Redcard_sum != 0 {
			achieveinfodata.Redcard_sum += achieveinfo_viewdata[i].Redcard_sum
		}
	}
	return err, achieveinfodata
}

func GetPlayerdetailtransferinfo(transferinfoParam map[string]interface{}) (error, []transferinfo_view, int64) {
	var transferinfo_viewdata []transferinfo_view
	page := transferinfoParam["page"].(int)
	pageSize := transferinfoParam["limit"].(int)
	player_id := transferinfoParam["player_id"].(int)
	var total int64
	err := db.Table("tb_playermatch").Select("tb_match.season", "tb_common.type_name", "tb_playermatch.shoottarget", "tb_playermatch.sucsnatch", "tb_playermatch.redcard",
		"tb_playermatch.yellowcard", "tb_playermatch.touchball").Where("tb_playermatch.player_id = ? and tb_playermatch.is_del = false and tb_match.is_del = false and tb_common.is_del = false", player_id).
		Joins("left join tb_match on tb_match.match_id = tb_playermatch.match_id").Joins("left join tb_common on tb_common.group_id = 3 and tb_common.type_id = tb_match.leaguematch_type").
		Order("tb_match.match_time DESC").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&transferinfo_viewdata).Error
	return err, transferinfo_viewdata, total
}