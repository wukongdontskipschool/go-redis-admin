package go_redis_admin

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	Name   string `gorm:"unique; not null; size:64"`
	Pass   string `gorm:"not null; size:32"`
	RoleId uint   `gorm:"not null; comment:角色id"`
}

// 角色表
type Role struct {
	gorm.Model
	Name string `gorm:"unique; not null; size:64"`
}

// 规则表
type Rule struct {
	gorm.Model
	Rule string `gorm:"not null; index; size:256"`
	Act  string `gorm:"not null; size:16"`
	Desc string `gorm:"not null; size:256"`
}

type RedisList struct {
	gorm.Model
	MenuId uint   `gorm:"not null; index; comment:菜单id"`
	Desc   string `gorm:"not null; size:256"`
	Host   string `gorm:"not null; size:64"`
	Port   string `gorm:"not null; size:16"`
	Auth   string `gorm:"not null; size:64"`
}

type RedisListTypes struct {
	gorm.Model
	Name string `gorm:"not null; size:64"`
}

type Menu struct {
	gorm.Model
	Pid   uint   `gorm:"not null; comment:父id"`
	Name  string `gorm:"not null; size:64;"`
	Icon  string `gorm:"not null; size:64; comment:图标"`
	Url   string `gorm:"not null; size:256; comment:跳转地址"`
	Rule  string `gorm:"not null; size:256;"`
	State int8   `gorm:"not null; comment:1正常0删除"`
}
