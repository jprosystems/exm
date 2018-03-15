package service

import (
	"eXmapi/dao"
	"eXmapi/model"
)

func GetExamen() []model.Examen {
	exmlist, _ := dao.GetExamen()
	return exmlist
}
