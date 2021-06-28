package models

import (
	"gorm.io/gorm"
)

type SysUserInfo struct {
	Model
	Userid        string   `gorm:"primary_key;column:userid" json:"userid" `
	Name          string   `json:"name" gorm:"column:name"`
	DepartmentIds string   `json:"department_ids" gorm:"column:department_ids"`
	Departments   []string `gorm:"-"`
	NamePinyin    string   `json:"name_pinyin" gorm:"column:name_pinyin"`
	TagIds        string   `json:"tag_ids"  gorm:"column:tag_ids"`
	Passwd        string   `json:"passwd"  gorm:"column:passwd"`
	Tel           string   `json:"tel"  gorm:"column:tel"`
	Email         string   `json:"email"  gorm:"column:email"`
	Gender        int      `json:"gender"  gorm:"column:gender"`
	IconUrl       string   `json:"icon_url"  gorm:"column:icon_url"`
	Status        int      `json:"status"  gorm:"column:status"`
	RoleId        int      `json:"role_id"  gorm:"column:role_id"`
}

func (u *SysUserInfo) Add() error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *SysUserInfo) CheckPasswd2() (bool, error) {
	var su SysUserInfo
	err := db.Table("sys_user_info").Select("userid").Where("userid = ? and passwd = ?", u.Userid, u.Passwd).First(&su).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return true, err
	}

	if su.Userid != "" {
		return true, nil
	}
	return false, nil
}
