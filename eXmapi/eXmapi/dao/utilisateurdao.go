package dao

import (
	"eXmapi/model"
	//"encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	// "strconv"
	// "strings"
)

var (
	cmd_sel_frm_util string = "SELECT * FROM UTILISATEUR where USERNAME = '%s' and MOT_DE_PASSE = '%s'"
)

func Authen(username string, password string) model.Utilisateur {
	cmd := fmt.Sprintf(cmd_sel_frm_util, username, password)
	var utilelem model.Utilisateur
	utilelem.UTILISATEURID = ""
	utilelem.USERNAME = ""
	utilelem.MOT_DE_PASSE = ""
	utilelem.TYPE_UTILISATEUR = ""

	var errorStr string = ""
	conn, err := ConnectToDB()
	defer conn.Close()

	rows, err := conn.Query(cmd)
	for rows.Next() {
		err := rows.Scan(&utilelem.UTILISATEURID, &utilelem.USERNAME, &utilelem.MOT_DE_PASSE, &utilelem.TYPE_UTILISATEUR, &utilelem.MAIL)
		if err != nil {
			errorStr += err.Error()
		}
	}
	if err != nil {
		errorStr += err.Error()
	}
	return utilelem
}
