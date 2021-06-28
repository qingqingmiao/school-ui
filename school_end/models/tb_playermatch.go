package models

import (
	"time"
)

type playersdetailachieve struct {
	Match_sum int
	Goal_sum int
	Assist_sum int
	Yellowcard_sum int
	Redcard_sum int
}

type playermatch struct {
	Player_id int `gorm:"column:player_id"`
	Chinese_name string `gorm:"column:chinese_name"`
	Match_id int `gorm:"column:match_id"`
	Team_id int `gorm:"column:team_id"`
	Team_name string `gorm:"column:team_name"`
	Player_uniform int `gorm:"column:player_uniform"`
	Position string `gorm:"column:position"`
	Totaltime string `gorm:"column:totaltime"`
	Is_start string `gorm:"column:is_start"`
	Offreplaced int `gorm:"column:offreplaced"`
	Onreplaced int `gorm:"column:onreplaced"`
	Shoot int `gorm:"column:shoot"`
	Shoottarget int `gorm:"column:shoottarget"`
	Passball int `gorm:"column:passball"`
	Sucpassball int `gorm:"column:sucpassball"`
	Forepassball int `gorm:"column:forepassball"`
	Sucforepassball int `gorm:"column:sucforepassball"`
	Passmid int `gorm:"column:passmid"`
	Sucpassmid int `gorm:"column:sucpassmid"`
	Airfight int `gorm:"column:airfight"`
	Sucairfight int `gorm:"column:sucairfight"`
	Groundfight int `gorm:"column:groundfight"`
	Sucgroundfight int `gorm:"column:sucgroundfight"`
	Foul int `gorm:"column:foul"`
	Snatch int `gorm:"column:snatch"`
	Sucsnatch int `gorm:"column:sucsnatch"`
	Violated int `gorm:"column:violated"`
	Yellowcard int `gorm:"column:yellowcard"`
	Redcard int `gorm:"column:redcard"`
	Offside int `gorm:"column:offside"`
	Touchball int `gorm:"column:touchball"`
	Loseball int `gorm:"column:loseball"`
	Totalrun int `gorm:"column:totalrun"`
	Sprint int `gorm:"column:sprint"`
	Speed int `gorm:"column:speed"`
	Fast int `gorm:"column:fast"`
	Fastrun int `gorm:"column:fastrun"`
	Slowrun int `gorm:"column:slowrun"`
	Walk int `gorm:"column:walk"`
	Strength int `gorm:"column:strength"`
	Strengthnum int `gorm:"column:strengthnum"`
	Sprintnum int `gorm:"column:sprintnum"`
	Goal int `gorm:"column:goal"`
	Assist int `gorm:"column:assist"`
	Breach int `gorm:"column:breach"`
	Sucbreach int `gorm:"column:sucbreach"`
	Break int `gorm:"column:break"`
	Sucbreak int `gorm:"column:sucbreak"`
	Threaten int `gorm:"column:threaten"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type playermatch_view struct {
	Match_id int
	Team_id int
	Team_shoot int
	Team_shoottarget int
	Team_passball int
	Team_sucpassball int
	Team_forepassball int
	Team_sucforepassball int
	Team_passmid int
	Team_sucpassmid int
	Team_airfight int
	Team_sucairfight int
	Team_groundfight int
	Team_sucgroundfight int
	Team_foul int
	Team_snatch int
	Team_sucsnatch int
	Team_violated int
	Team_yellowcard int
	Team_redcard int
	Team_offside int
	Team_touchball int
	Team_loseball int
	Team_totalrun int
	Team_sprint int
	Team_speed int
	Team_fast int
	Team_fastrun int
	Team_slowrun int
	Team_walk int
	Team_strength int
	Team_strengthnum int
	Team_sprintnum int
	Team_goal int
	Team_assist int
	Team_breach int
	Team_sucbreach int
	Team_break int
	Team_sucbreak int
	Team_threaten int
}

func GetPlayermatchByPlayerID(player_id string) (error, playersdetailachieve) {
	var playersdetailachievedata playersdetailachieve
	err := db.Table("tb_playermatch").Select("count(tb_playermatch.player_id) as match_sum, sum(tb_playermatch.goal) as goal_sum, sum(tb_playermatch.assist) as assist_sum, " +
		"sum(tb_playermatch.yellowcard) as yellowcard_sum, sum(tb_playermatch.redcard) as redcard_sum").Where("tb_playermatch.player_id = ? and tb_playermatch.is_del = false", player_id).
		Find(&playersdetailachievedata).Error
	return err, playersdetailachievedata
}

func GetPlayermatchteamByTeamID(match_id string, team_id string) (error, playermatch_view) {
	var playermatch_viewdata playermatch_view
	err := db.Table("tb_playermatch").Select("match_id, team_id, sum(shoot) as Team_shoot, sum(shoottarget) as Team_shoottarget, sum(passball) as Team_passball, sum(sucpassball) as Team_sucpassball, " +
		"sum(forepassball) as Team_forepassball, sum(sucforepassball) as Team_sucforepassball, sum(passmid) as Team_passmid, sum(sucpassmid) as Team_sucpassmid, sum(airfight) as Team_airfight, sum(sucairfight) as Team_sucairfight, " +
		"sum(groundfight) as Team_groundfight, sum(sucgroundfight) as Team_sucgroundfight, sum(foul) as Team_foul, sum(snatch) as Team_snatch, sum(sucsnatch) as Team_sucsnatch, sum(violated) as Team_violated, " +
		"sum(yellowcard) as Team_yellowcard, sum(redcard) as Team_redcard, sum(offside) as Team_offside, sum(touchball) as Team_touchball, sum(totalrun) as Team_totalrun, sum(sprint) as Team_sprint, sum(speed) as Team_speed, " +
		"sum(fast) as Team_fast, sum(fastrun) as Team_fastrun, sum(slowrun) as Team_slowrun, sum(walk) as Team_walk, sum(strength) as Team_strength, sum(strengthnum) as Team_strengthnum, sum(sprintnum) as Team_sprintnum, " +
		"sum(goal) as Team_goal, sum(assist) as Team_assist, sum(breach) as Team_breach, sum(sucbreach) as Team_sucbreach, sum(break) as Team_break, sum(sucbreak) as Team_sucbreak, sum(threaten) as Team_threaten").
		Where("match_id = ? and team_id = ? and is_del = ?", match_id, team_id, false).Group("match_id, team_id").Find(&playermatch_viewdata).Error
	return err, playermatch_viewdata
}

func GetPlayermatchplayerByMatchID(match_id string) (error, []playermatch) {
	var playermatchdata []playermatch
	err := db.Table("tb_playermatch").Select("tb_playermatch.player_id, tb_player.chinese_name, tb_playermatch.match_id, tb_playermatch.team_id, tb_team.team_name, tb_playermatch.player_uniform, " +
		"tb_playermatch.position, tb_playermatch.totaltime, tb_playermatch.is_start, tb_playermatch.offreplaced, tb_playermatch.onreplaced, tb_playermatch.shoot, tb_playermatch.shoottarget, tb_playermatch.passball, " +
		"tb_playermatch.sucpassball, tb_playermatch.forepassball, tb_playermatch.sucforepassball, tb_playermatch.passmid, tb_playermatch.sucpassmid, tb_playermatch.airfight, tb_playermatch.sucairfight, " +
		"tb_playermatch.groundfight, tb_playermatch.sucgroundfight, tb_playermatch.foul, tb_playermatch.snatch, tb_playermatch.sucsnatch, tb_playermatch.violated, tb_playermatch.yellowcard, tb_playermatch.redcard, " +
		"tb_playermatch.offside, tb_playermatch.touchball, tb_playermatch.loseball, tb_playermatch.totalrun, tb_playermatch.sprint, tb_playermatch.speed, tb_playermatch.fast, tb_playermatch.fastrun, tb_playermatch.slowrun, " +
		"tb_playermatch.walk, tb_playermatch.strength, tb_playermatch.strengthnum, tb_playermatch.sprintnum, tb_playermatch.goal, tb_playermatch.assist, tb_playermatch.breach, tb_playermatch.sucbreach, tb_playermatch.break, " +
		"tb_playermatch.sucbreak, tb_playermatch.threaten, tb_playermatch.is_del, tb_playermatch.update_userid, tb_playermatch.update_time").
		Where("tb_playermatch.match_id = ? and tb_playermatch.is_del = ?", match_id, false).Joins("left join tb_player on tb_player.player_id = tb_playermatch.player_id").
		Joins("left join tb_team on tb_team.team_id = tb_playermatch.team_id").Find(&playermatchdata).Error
	return err, playermatchdata
}