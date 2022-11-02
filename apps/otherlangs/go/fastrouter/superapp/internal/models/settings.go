package models

import ()

type Settings struct {
	Id int 		`json:"id" gorm:"primaryKey"`
	Tag string		`json:"tag"`
	Value string	`json:"value"`
}

//
type SettingsValueUnpacked struct {
	Host string		`json:"host"`
	Percent float64	`json:"percent" `
}

type SettingsUnpack struct {
	Id int 		`json:"id"`
	Tag string		`json:"tag"`
	Value []SettingsValueUnpacked `json:"value"`
}
