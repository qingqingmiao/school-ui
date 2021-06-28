package models

import (
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/gredis"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysMenu struct {
	Id        int       `gorm:"primary_key; column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Icon      string    `gorm:"column:icon" json:"icon"`
	Order     int       `grom:"column:order" json:"order"`
	ParentId  int       `grom:"column:parent_id" json:"parent_id"`
	Path      string    `grom:"column:path" json:"path"`
	Redirect  string    `grom:"column:redirect" json:"redirect"`
	Name      string    `grom:"column:name" json:"name"`
	Component string    `grom:"column:component" json:"component"`
	Children  []SysMenu `grom:"-" json:"children"`
}

func (m *SysMenu) InsertOneMenu() error {
	return db.Table("sys_menu").Create(m).Error
}

func (m *SysMenu) DelOneMenu() error {
	var menu SysMenu
	err := db.Table("sys_menu").First(&menu, m.Id).Error
	if err != nil {
		return err
	}
	err = db.Table("sys_menu").Where("parent_id = ?", m.Id).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		err = db.Delete(m).Error
		return err
	}
	return errors.New("拥有子类，无法删除")
}

func (m *SysMenu) EditOne() error {
	pa := map[string]interface{}{}
	if m.ParentId != 0 {
		pa["parent_id"] = m.ParentId
	}
	if m.Name != "" {
		pa["name"] = m.Name
	}
	if m.Title != "" {
		pa["title"] = m.Title
	}
	if m.Icon != "" {
		pa["icon"] = m.Icon
	}
	if m.Component != "" {
		pa["component"] = m.Component
	}
	if m.Redirect != "" {
		pa["redirect"] = m.Redirect
	}
	if m.Path != "" {
		pa["path"] = m.Path
	}
	if m.Order != -1 {
		pa["order"] = m.Order
	}
	if len(pa) != 0 {
		return db.Table("sys_menu").Where("id = ?", m.Id).Updates(pa).Error
	}
	return nil
}

func GetMenuTree() []SysMenu {
	var menu_tree []SysMenu
	err := gredis.RedisConn.GetObject(e.MenuRedisKey, &menu_tree)
	if err == nil {
		return menu_tree
	}
	db.Table("sys_menu").Find(&menu_tree)
	if len(menu_tree) == 0 {
		return menu_tree
	}
	parent_infos := getParentListData(menu_tree)
	tree_data := data2TreeData(menu_tree,parent_infos)
	gredis.RedisConn.Set(e.MenuRedisKey,tree_data,0)
	return tree_data
}

func getParentListData(menus []SysMenu) []SysMenu {
	var parent_menu []SysMenu
	for _, each := range menus {
		if each.ParentId == 0 {
			parent_menu = append(parent_menu, each)
		}
	}
	return parent_menu
}

func data2TreeData(menus, parent []SysMenu) []SysMenu {
	for p,each_p := range parent{
		children := []SysMenu{}
		for _,each := range menus{
			if each.ParentId == each_p.Id{
				children = append(children, each)
			}
		}
		parent[p].Children = children
		if len(children) > 0{
			data2TreeData(menus,children)
		}
	}
	return parent
}
