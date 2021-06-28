package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User_Account struct {
	Account string `gorm:"column:account"`
	Password string `gorm:"column:password"`
}

type user struct {
	User_id int `gorm:"column:user_id"`
	User_type string `gorm:"column:user_type"`
	Account string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	Username string `gorm:"column:username"`
	Sex string `gorm:"column:sex"`
	Email string `gorm:"column:email"`
	Telephone string `gorm:"column:telephone"`
	Create_founder string `gorm:"column:create_founder"`
	Create_time time.Time `gorm:"column:create_time"`
	Portrait string `gorm:"column:portrait"`
	Thumbnail_portrait string `gorm:"column:thumbnail_portrait"`
	Focus_teamid postgres.Jsonb `gorm:"column:focus_teamid"`
	Focus_playerid postgres.Jsonb `gorm:"column:focus_playerid"`
	Due_time string `gorm:"column:due_time"`
	Id_code string `gorm:"column:id_code"`
	State string `gorm:"column:state"`
	Nationality string `gorm:"column:nationality"`
	Team_id string `gorm:"column:team_id"`
	Birthday string `gorm:"column:birthday"`
	Affiliatedunits string `gorm:"column:affiliatedunits"`
	Address string `gorm:"column:address"`
	Module_authority postgres.Jsonb `gorm:"column:module_authority"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type userfocus struct {
	User_id int `gorm:"column:user_id"`
	User_type string `gorm:"column:user_type"`
	Focus_teamid postgres.Jsonb `gorm:"column:focus_teamid"`
	Focus_playerid postgres.Jsonb `gorm:"column:focus_playerid"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type userfocusteamid struct {
	Teamid []int
}
type userfocusplayerid struct {
	Playerid []int
}

type userfocusteam struct {
	Focus_teamid int
	Focus_teamname string
}

type userfocusplayer struct {
	Focus_playerid int
	Focus_playername string
}

type dashboardnum_view struct {
	Player_count int64
	Train_count int64
	Injuries_count int64
	Coach_count int64
}

func (u *User_Account) CheckPasswd() (bool, error) {
	var su User_Account
	u.Password = EncodeMD5(u.Password)
	err := db.Table("tb_user").Select("account").Where("account = ? and password = ?", u.Account, u.Password).First(&su).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return true, err
	}

	if su.Account != "" {
		return true, nil
	}
	return false, nil
}

func GetUserByID(user_id string) (error, user) {
	var userdata user
	err := db.Table("tb_user").Where("user_id = ?", user_id).First(&userdata).Error
	return err, userdata
}

func UpdateUser(a *gin.Context) (error, user) {
	var userdata user
	a.BindJSON(&userdata)
	var passwordtemp string
	if userdata.Password != ""{
		passwordtemp = EncodeMD5(userdata.Password)
	}
	userdata.Update_time = time.Now()
	userdata.Password = passwordtemp
	var user_id = userdata.User_id
	err := db.Table("tb_user").Where("user_id = ?", user_id).Updates(&userdata).Error
	return err, userdata
}

func UpdateUserPortrait(a *gin.Context, user_portrait string) (error, user) {
	var userdata user
	user_id := a.Params.ByName("user_id")
	userdata.User_id, _ = strconv.Atoi(user_id)
	userdata.Portrait = user_portrait
	userdata.Thumbnail_portrait = user_portrait
	userdata.Update_userid = "0"
	userdata.Update_time = time.Now()
	err := db.Table("tb_user").Where("user_id = ?", user_id).Updates(&userdata).Error
	return err, userdata
}

func InsertUser(a *gin.Context) (error, user) {
	var userdata user
	a.BindJSON(&userdata)
	err := db.Table("tb_user").Select("user_id","account","password").Create(&userdata).Error
	return err, userdata
}

func DeleteUser(a *gin.Context) error {
	user_id := a.Params.ByName("id")
	err := db.Table("tb_user").Where("user_id = ?", user_id).Update("is_del", true).Error
	return err
}

//func UpdateUserRole(a *gin.Context) (error, user) {
//	var userdata user
//	a.BindJSON(&userdata)
//	userdata.User_type = time.Now()
//	var user_id = userdata.User_id
//	err := db.Table("tb_user").Where("user_id = ?", user_id).Updates(&userdata).Error
//	return err, userdata
//}

func GetAllUser(userParam map[string]interface{}) (error, []user, int64) {
	var users []user
	page := userParam["page"].(int)
	pageSize := userParam["limit"].(int)
	var total int64
	err := db.Table("tb_user").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&users).Error
	return err, users, total
}

func GetUserAuthByAccount(account string) (error, user) {
	var userdata user
	err := db.Table("tb_user").Where("account = ?", account).First(&userdata).Error
	return err, userdata
}


func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func GetUserFocus(user_id string) (error, []userfocusteam, []userfocusplayer) {
	var userfocusdata userfocus
	err := db.Table("tb_user").Select("user_id, user_type, focus_teamid, focus_playerid").Where("user_id = ?", user_id).Find(&userfocusdata).Error
	Focus_teamid, _ := json.Marshal(userfocusdata.Focus_teamid)
	userfocusteamidTemp := userfocusteamid{}
	json.Unmarshal(Focus_teamid, &userfocusteamidTemp)
	Focus_playerid, _ := json.Marshal(userfocusdata.Focus_playerid)
	Focus_playeridTemp := userfocusplayerid{}
	json.Unmarshal(Focus_playerid, &Focus_playeridTemp)
	var userfocusteamdata []userfocusteam
	var userfocusplayerdata []userfocusplayer
	err = db.Table("tb_team").Select("team_id as focus_teamid, team_name as focus_teamname").Where("team_id in (?)", userfocusteamidTemp.Teamid).Find(&userfocusteamdata).Error
	err = db.Table("tb_player").Select("player_id as focus_playerid, chinese_name as focus_playername").Where("player_id in (?)", Focus_playeridTemp.Playerid).Find(&userfocusplayerdata).Error
	return err, userfocusteamdata, userfocusplayerdata
}

func UpdateUserFocus(a *gin.Context) (error, userfocus) {
	var userfocusdate userfocus
	a.BindJSON(&userfocusdate)
	userfocusdate.Update_userid = "0"
	userfocusdate.Update_time = time.Now()
	err := db.Table("tb_user").Where("user_id = ?", userfocusdate.User_id).Updates(&userfocusdate).Error
	return err, userfocusdate
}

func GetDashboardnum() (error, dashboardnum_view) {
	var dashboardnum_viewdata dashboardnum_view
	err := db.Table("tb_player").Where("is_del = false").Count(&dashboardnum_viewdata.Player_count).Error
	err = db.Table("tb_trainplan").Where("is_del = false").Count(&dashboardnum_viewdata.Train_count).Error
	err = db.Table("tb_injury").Distinct("player_id").Where("is_del = false").Count(&dashboardnum_viewdata.Injuries_count).Error
	err = db.Table("tb_user").Where("user_type = '6' and is_del = false").Count(&dashboardnum_viewdata.Coach_count).Error
	return err, dashboardnum_viewdata
}
