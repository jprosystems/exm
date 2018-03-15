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

	app.Run(iris.Addr(":8080"))
}

// func main() {
// 	r, _ := json.Marshal(service.GetExamen())
// 	fmt.Println(string(r))
// }
