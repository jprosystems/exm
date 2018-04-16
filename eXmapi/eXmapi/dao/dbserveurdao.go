package dao

import (
	"database/sql"
	"eXmapi/model"
	"encoding/json"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"io/ioutil"
)

var (
	n           int    = -1
	errorString string = " "
	dbpath      string
	DbFileName  string = ""
)

func GetDbParameters(fn string) model.DbServer {
	plan, _ := ioutil.ReadFile(fn)
	var data model.DbServer
	err := json.Unmarshal(plan, &data)
	if err != nil {
		defer fmt.Println(err)
	}
	return data
}

func GetDbPath(db model.DbServer) string {
	return db.User + ":" + db.Pass + "@" + db.Host + "/" + db.DbPath
}

func ConnectToDB() (*sql.DB, error) {
	dbpath := GetDbPath(GetDbParameters(DbFileName))
	conn, err := sql.Open("firebirdsql", dbpath)
	return conn, err
}
