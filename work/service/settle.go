package service

import (
	jgorm "github.com/jinzhu/gorm"
)

type SettleStruct struct {
	tx *jgorm.DB
	tm int64
}

func NewSettleStruct() *SettleStruct {
	return &SettleStruct{}
}

func (ts *SettleStruct) Init() {

}
func (ts *SettleStruct) Run() {
	Glock.Lock()
	defer func() {
		Glock.Unlock()
	}()

}
