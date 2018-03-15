package dao

import (
	"eXmapi/model"
	//"encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "strconv"
	// "strings"
)

var (
	cmd_sel_frm_exm string = "SELECT * FROM EXAMEN"
)

func GetExamen() ([]model.Examen, string) {
	var exmlist []model.Examen
	var exmelem model.Examen

	var errorStr string = " "
	conn, err := ConnectToDB()
	defer conn.Close()

	rows, err := conn.Query(cmd_sel_frm_exm)
	for rows.Next() {
		err := rows.Scan(&exmelem.CODE, &exmelem.SEMESTRE, &exmelem.COURS_GROUPE, &exmelem.DESCRIPTION, &exmelem.LANGUE, &exmelem.NOMBRE_QUESTIONS, &exmelem.NOMBRE_POINTS_TOTAL, &exmelem.DUREE_MAXIMALE, &exmelem.DATE_HEURE_ACTIVATION, &exmelem.DATE_DERNIERE_MODIFICATION)
		exmlist = append(exmlist, exmelem)
		if err != nil {
			errorStr += err.Error()
		}
	}
	if err != nil {
		errorStr += err.Error()
	}
	return exmlist, errorStr
}
