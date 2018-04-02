package dao

import (
	"eXmapi/model"
	//"encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	"strconv"
	// "strings"
)

var (
	cmd_sel_frm_exm   string = "SELECT * FROM EXAMEN order by DATE_HEURE_ACTIVATION desc"
	cmd_sel_count     string = "select count(*) as nb from GENID"
	cmd_insr_to_exm   string = "INSERT INTO EXAMEN (EXAMENID, SEMESTRE, COURS_GROUPE, DESCRIPTION, LANGUE, NOMBRE_QUESTIONS, NOMBRE_POINTS_TOTAL, DUREE_MAXIMALE, DATE_HEURE_ACTIVATION, DATE_DERNIERE_MODIFICATION) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')"
	cmd_updt_exm_code string = "UPDATE EXAMEN SET SEMESTRE = '%s' WHERE EXAMENID = '%s'; "
	cmd_del_exm_code  string = "delete from EXAMEN where EXAMENID = '%s'"
	cmd_generate_id   string = "INSERT INTO GENID (COLUMNGENID) VALUES ('GENID' )"
)

func GetExamen() ([]model.Examen, string) {
	var exmlist []model.Examen
	var exmelem model.Examen

	var errorStr string = ""
	conn, err := ConnectToDB()
	defer conn.Close()

	rows, err := conn.Query(cmd_sel_frm_exm)
	for rows.Next() {
		err := rows.Scan(&exmelem.EXAMENID, &exmelem.SEMESTRE, &exmelem.COURS_GROUPE, &exmelem.DESCRIPTION, &exmelem.LANGUE, &exmelem.NOMBRE_QUESTIONS, &exmelem.NOMBRE_POINTS_TOTAL, &exmelem.DUREE_MAXIMALE, &exmelem.DATE_HEURE_ACTIVATION, &exmelem.DATE_DERNIERE_MODIFICATION)
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

func GetCode() (int, string) {
	var nb string

	var errorStr string = ""
	conn, err := ConnectToDB()
	defer conn.Close()

	rows, err := conn.Query(cmd_sel_count)
	for rows.Next() {
		err := rows.Scan(&nb)
		if err != nil {
			errorStr += err.Error()
		}
	}
	if err != nil {
		errorStr += err.Error()
	}
	intnb, err := strconv.Atoi(nb)
	if err != nil {
		errorStr += err.Error()
	}

	_, err = conn.Exec(cmd_generate_id)
	if err != nil {
		errorStr += err.Error()
	}

	return intnb + 1, errorStr
}

func SetExamen(objExam model.Examen) (string, string) {
	cmd := fmt.Sprintf(cmd_insr_to_exm, objExam.EXAMENID, objExam.SEMESTRE, objExam.COURS_GROUPE, objExam.DESCRIPTION, objExam.LANGUE, objExam.NOMBRE_QUESTIONS, objExam.NOMBRE_POINTS_TOTAL, objExam.DUREE_MAXIMALE, objExam.DATE_HEURE_ACTIVATION, objExam.DATE_DERNIERE_MODIFICATION)
	var errorStr string = ""
	conn, err := ConnectToDB()
	fmt.Println(cmd)

	defer conn.Close()
	_, err = conn.Exec(cmd)

	if err != nil {
		errorStr += err.Error()
	}
	return objExam.EXAMENID, errorStr
}

func SetSemestreExamen(varSmestre string, varCode string) string {
	cmd := fmt.Sprintf(cmd_updt_exm_code, varSmestre, varCode)
	var errorStr string = ""
	conn, err := ConnectToDB()

	defer conn.Close()
	_, err = conn.Exec(cmd)

	if err != nil {
		errorStr += err.Error()
	}
	return errorStr
}

func DelExamenByCode(varCode string) string {
	cmd := fmt.Sprintf(cmd_del_exm_code, varCode)
	var errorStr string = ""
	conn, err := ConnectToDB()

	defer conn.Close()
	_, err = conn.Exec(cmd)

	if err != nil {
		errorStr += err.Error()
	}
	return errorStr
}
