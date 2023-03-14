package model

import (
	"gorm.io/gorm"
	"time"
)

// todo 调用接口统计剩余次数

type UserInterface struct {
	gorm.Model
	UserId          uint `json:"user_id" gorm:"comment:调用的用户id"`
	InterfaceInfoId uint `json:"interface_info_id" gorm:"comment:调用的接口id"`
	TotalNum        uint `json:"total_num" gorm:"comment:总调用次数"`
	LeftNum         uint `json:"left_num" gorm:"comment:剩余的调用次数"`
	Status          bool `json:"status" gorm:"comment:用户是否可继续调用该接口的判断"`
}

type NewUserInterface struct {
	UserId          uint `json:"user_id"`
	InterfaceInfoId uint `json:"interface_info_id"`
	TotalNum        uint `json:"total_num"`
	LeftNum         uint `json:"left_num"`
}

type QueryUserInterface struct {
	ID              uint      `json:"id"`
	UserId          uint      `json:"user_id"`
	InterfaceInfoId uint      `json:"interface_info_id"`
	TotalNum        uint      `json:"total_num"`
	LeftNum         uint      `json:"left_num"`
	Status          bool      `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UpdateUserInterface struct {
	ID              uint `json:"id"`
	InterfaceInfoId uint `json:"interface_info_id"`
	TotalNum        uint `json:"total_num"`
	LeftNum         uint `json:"left_num"`
	Status          bool `json:"status"`
}
