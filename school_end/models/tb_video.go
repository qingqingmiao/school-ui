package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/postgres"
	"strings"
	"time"
)

type video struct {
	Video_id int `gorm:"column:video_id"`
	Video_keyword postgres.Jsonb `gorm:"column:video_keyword"`
	Video_time string `gorm:"column:video_time"`
	Video_title string `gorm:"column:video_title"`
	Video_content string `gorm:"column:video_content"`
	Video_duration string `gorm:"column:video_duration"`
	Storage_path string `gorm:"column:storage_path"`
	Video_thumbnail string `gorm:"column:video_thumbnail"`
	Is_del bool `gorm:"column:is_del"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

type Videokeyword struct {
	Key []string
}

func GetAllVideo(videoParam map[string]interface{}) (error, []video, int) {
	var videos []video
	page := videoParam["page"].(int)
	pageSize := videoParam["limit"].(int)
	video_keyword := videoParam["video_keyword"].(string)
	//keys：前端搜索框传来的关键字
	keys := strings.Fields(video_keyword)
	var total int
	err := db.Table("tb_video").Where("is_del = false").Find(&videos).Error

	for index := 0; index < len(videos); index++ {
		keywords, _ := json.Marshal(videos[index].Video_keyword)
		//keywordsTemp：数据库中某个视频中包含的关键字
		keywordsTemp := Videokeyword{}
		json.Unmarshal(keywords, &keywordsTemp)
		keywordsTemp.Key = append(keywordsTemp.Key, videos[index].Video_title)
		cnt := 0
		for _, v1 := range keys {
			flag := false
			for _, v2 := range keywordsTemp.Key {
				if(strings.Contains(v2, v1)) {
					cnt++
					flag = true
					break
				}
			}
			//假如前端搜索框传来的某个关键字在数据库中不存在，那么剩下的也就不需要遍历了
			if(!flag) {
				break
			}
		}
		if cnt != len(keys) {
			videos = append(videos[:index], videos[index+1:]...)
			index--
		}
	}
	total = len(videos)
	start := (page-1)*pageSize
	end := start + pageSize
	if(end > total) {
		end = total
	}
	return err, videos[start:end], total
}

func AddVideo(videoAddData map[string]interface{}) (error, video) {
	var videodata video
	keys := videoAddData["video_keyword"].(string)

	var VideokeywordTemp Videokeyword
	json.Unmarshal([]byte(keys), &VideokeywordTemp)
	keywords, _ := json.Marshal(VideokeywordTemp)
	videodata.Video_keyword = postgres.Jsonb{keywords}
	videodata.Video_time = videoAddData["video_time"].(string)
	videodata.Video_title = videoAddData["video_title"].(string)
	videodata.Video_content = videoAddData["video_content"].(string)
	videodata.Video_duration = videoAddData["video_duration"].(string)
	videodata.Storage_path = videoAddData["storage_path"].(string)
	videodata.Video_thumbnail = videoAddData["video_thumbnail"].(string)
	videodata.Is_del = false
	videodata.Update_userid = "0"
	videodata.Update_time = time.Now()
	err := db.Table("tb_video").Select("video_keyword","video_time","video_title","video_content","video_duration",
		"storage_path","video_thumbnail","is_del","update_userid","update_time").Create(&videodata).Error
	return err, videodata
}

func UpdateVideo(videoUpdateDate map[string]interface{}) (error, video) {
	var videodata video
	keys := videoUpdateDate["video_keyword"].(string)

	var VideokeywordTemp Videokeyword
	json.Unmarshal([]byte(keys), &VideokeywordTemp)
	keywords, _ := json.Marshal(VideokeywordTemp)
	videodata.Video_keyword = postgres.Jsonb{keywords}
	videodata.Video_id = videoUpdateDate["video_id"].(int)
	videodata.Video_time = videoUpdateDate["video_time"].(string)
	videodata.Video_title = videoUpdateDate["video_title"].(string)
	videodata.Video_content = videoUpdateDate["video_content"].(string)
	videodata.Video_duration = videoUpdateDate["video_duration"].(string)
	videodata.Storage_path = videoUpdateDate["storage_path"].(string)
	videodata.Video_thumbnail = videoUpdateDate["video_thumbnail"].(string)
	videodata.Update_userid = "0"
	videodata.Update_time = time.Now()
	err := db.Table("tb_video").Select("video_keyword","video_time","video_title","video_content","video_duration",
		"storage_path","video_thumbnail","update_userid","update_time").Where("video_id = ?", videodata.Video_id).Updates(&videodata).Error
	return err, videodata
}

func DelVideo(a *gin.Context) error {
	id := a.Params.ByName("id")
	err := db.Table("tb_video").Where("video_id = ?", id).Update("is_del", true).Error
	return err
}