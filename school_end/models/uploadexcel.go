package models

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"strconv"
	"time"
)

type tb_player struct {
	Chinese_name string `gorm:"column:chinese_name"`
	English_name string `gorm:"column:english_name"`
	Id_code string `gorm:"column:id_code"`
	Birthday time.Time `gorm:"column:birthday"`
	Sex string `gorm:"column:sex"`
	State string `gorm:"column:state"`
	Is_del bool `gorm:"column:is_del"`
	Former_name string `gorm:"column:former_name"`
	Height string `gorm:"column:height"`
	Weight string `gorm:"column:weight"`
	Technical_level string `gorm:"column:technical_level"`
	Native_place string `gorm:"column:native_place"`
	Certificate_type string `gorm:"column:certificate_type"`
	Update_userid string `gorm:"column:update_userid"`
	Update_time time.Time `gorm:"column:update_time"`
}

func AddExcel(r io.Reader) (error) {
	xlsx, err := excelize.OpenReader(r)

	if err != nil {
		return err
	}
	style, err := xlsx.NewStyle(`{"number_format": 0}`)
	if err != nil {
		return err
	}

	for  i:=2; i<=5000; i++ {
		xlsx.SetCellStyle("Sheet1", "D2", "D"+strconv.Itoa(i), style)
	}

	rows := xlsx.GetRows("Sheet" + "1")
	var err2 error
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			timeNow := time.Now()
			//var timeLayoutStr = "2006-01-02"
			birthday := excelDateToDate(data[3])
			//birthday, _ := time.Parse(timeLayoutStr, data[3])
			playerdata := tb_player{data[0], data[1], data[2], birthday, data[4], data[5], false, data[10], data[11], data[12],
				data[13], data[14], data[17], "0", timeNow}
			err2 = db.Table("tb_player").Create(&playerdata).Error
		}
	}
	return err2
}

func excelDateToDate(excelDate string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}