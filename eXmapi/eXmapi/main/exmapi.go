package main

import (
	"eXmapi/service"
	// "encoding/json"
	// "fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	// Method:   GET, Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("eXm Works")
	})

	// Method:   GET, Resource: http://localhost:8080/exmapi/examen/list
	app.Get("/exmapi/examen/list", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")
		ctx.JSON(service.GetExamen())
	})

	app.Get("/exmapi/examen/create/{SEMESTRE}/{COURS_GROUPE}/{DESCRIPTION}/{LANGUE}/{DUREE_MAXIMALE}/{DATE_HEURE_ACTIVATION}", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		varSEMESTRE := ctx.Params().Get("SEMESTRE")
		varCOURS_GROUPE := ctx.Params().Get("COURS_GROUPE")
		varDESCRIPTION := ctx.Params().Get("DESCRIPTION")
		varLANGUE := ctx.Params().Get("LANGUE")
		varDUREE_MAXIMALE := ctx.Params().Get("DUREE_MAXIMALE")
		varDATE_HEURE_ACTIVATION := ctx.Params().Get("DATE_HEURE_ACTIVATION")
		ctx.Writef(service.SetExamen(varSEMESTRE, varCOURS_GROUPE, varDESCRIPTION, varLANGUE, varDUREE_MAXIMALE, varDATE_HEURE_ACTIVATION))
	})

	//http://localhost:8080/exmapi/examen/HV-2018/EX-1
	app.Get("/exmapi/examen/update/{semestre}/{code}", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		semestre := ctx.Params().Get("semestre")
		cdoe := ctx.Params().Get("code")
		ctx.Writef(service.SetSemestreExamen(semestre, cdoe))
	})

	//http://localhost:8080/exmapi/examen/del/EX-1
	app.Get("/exmapi/examen/del/{code}", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		cdoe := ctx.Params().Get("code")
		ctx.Writef(service.DelExamenByCode(cdoe))
	})

	app.Run(iris.Addr(":8080"))
}

// func main() {

// 	// r, _ := json.Marshal(service.GetExamen())
// 	// fmt.Println(string(r))
// 	jsonStr := "{\"EXAMENID\":\"EX-1\",\"SEMESTRE\":\"HV-2018\",\"COURS_GROUPE\":\"GEL500-010\",\"DESCRIPTION\":\"Description\",\"LANGUE\":\"EN\",\"NOMBRE_QUESTIONS\":\"20\",\"NOMBRE_POINTS_TOTAL\":\"100\",\"DUREE_MAXIMALE\":\"120\",\"DATE_HEURE_ACTIVATION\":\"27.03.2018 05:08\",\"DATE_DERNIERE_MODIFICATION\":\"2014-04-12 11:37:12Z\"}"
// 	service.SetExamen(jsonStr)
// }
