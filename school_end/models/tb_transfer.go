package models

type transfer_view struct {
	Transfer_id int
	Player_id int
	Transfer_time string
	From_teamid int
	From_teamname string
	To_teamid int
	To_teamname string
	Is_del bool
}

func GetTransfer(transferParam map[string]interface{}) (error, []transfer_view, int64) {
	var transfer_viewdata []transfer_view
	page := transferParam["page"].(int)
	pageSize := transferParam["limit"].(int)
	player_id := transferParam["player_id"].(int)
	var total int64
	err := db.Table("transferList").Select("transfer_id, player_id, transfer_time, from_teamid, from_teamname, to_teamid, to_teamname, is_del").Where("player_id = ? and is_del = false", player_id).Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&transfer_viewdata).Error
	return err, transfer_viewdata, total
}