package model

import (
	"gorm.io/gorm"
	"time"
)

type InterfaceInfo struct {
	gorm.Model

	UserId         uint   `json:"user_id" gorm:"comment:创建人的ID;not null"`
	Name           string `json:"name" gorm:"comment:接口名称;size:256;not null"`
	Description    string `json:"description" gorm:"comment:接口描述;size:256"`
	Url            string `json:"url" gorm:"comment:接口地址;size:512;not null"`
	Method         string `json:"method" gorm:"comment:请求类型;not null"`
	RequestHeader  string `json:"request_header" gorm:"comment:请求头"`
	ResponseHeader string `json:"response_header" gorm:"comment:响应头"`
	RequestParams  string `json:"request_params" gorm:"comment:请求参数"`
	Status         bool   `json:"status" gorm:"comment:接口状态：0关闭，1开启;not null"`
}

type NewInterface struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	Url            string `json:"url" binding:"required"`
	Method         string `json:"method" binding:"required"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	RequestParams  string `json:"request_params"`
}

type QueryInterface struct {
	ID             uint      `json:"id"`
	UserId         uint      `json:"user_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Url            string    `json:"url"`
	Method         string    `json:"method"`
	RequestHeader  string    `json:"request_header"`
	ResponseHeader string    `json:"response_header"`
	RequestParams  string    `json:"request_params"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         bool      `json:"status"`
}

type UpdateInterface struct {
	ID             uint   `json:"id" binding:"required"`
	UserId         uint   `json:"user_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	Url            string `json:"url" binding:"required"`
	Method         string `json:"method" binding:"required"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	RequestParams  string `json:"request_params"`
	Status         bool   `json:"status" binding:"required"`
}

type InvokeInterface struct {
	AccessKey string `json:"access_key" binding:"required"`
	SecretKey string `json:"secret_key" binding:"required"`
	Body      string `json:"body"`
}
