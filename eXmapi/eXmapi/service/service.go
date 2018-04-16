package service

import (
	"eXmapi/dao"
	"eXmapi/model"
	// "encoding/json"
	// "fmt"
	// "log"
	"strconv"
	"strings"
	"time"
)

func GetExamen() []model.Examen {
	exmlist, _ := dao.GetExamen()
	return exmlist
}

func SetExamen(varSEMESTRE string, varCOURS_GROUPE string, varDESCRIPTION string, varLANGUE string, varDUREE_MAXIMALE string, varDATE_HEURE_ACTIVATION string) string {
	nb, _ := dao.GetCode()
	tim := time.Now().Format("02.01.2006 15:04")

	var objExam model.Examen

	objExam.EXAMENID = "EX-" + strconv.Itoa(nb)
	objExam.SEMESTRE = varSEMESTRE
	objExam.COURS_GROUPE = varCOURS_GROUPE
	objExam.DESCRIPTION = varDESCRIPTION
	objExam.LANGUE = varLANGUE
	objExam.NOMBRE_QUESTIONS = "0"
	objExam.NOMBRE_POINTS_TOTAL = "0"
	objExam.DUREE_MAXIMALE = varDUREE_MAXIMALE
	objExam.DATE_HEURE_ACTIVATION = varDATE_HEURE_ACTIVATION
	objExam.DATE_DERNIERE_MODIFICATION = tim

	code, errorStr := dao.SetExamen(objExam)

	var varReturn string = ""
	if strings.Compare(errorStr, "") == 0 {
		varReturn = code
	} else {
		varReturn = "erreur"
	}
	return varReturn
}

func SetSemestreExamen(semestre string, cdoe string) string {
	err := dao.SetSemestreExamen(semestre, cdoe)
	return err
}

func DelExamenByCode(cdoe string) string {
	err := dao.DelExamenByCode(cdoe)
	return err
}

func Authen(username string, password string) model.Utilisateur {
	user := dao.Authen(username, password)
	return user
}
