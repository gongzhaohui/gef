package main

import (
	"gef/pkg/components"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// var webFiles embed.FS

func main() {
	app.Route("/", func() app.Composer { return &components.DataTable{} })
	// app.HandleStatic("/", webFiles)
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

// 前端服务（开发模式）
// func runHTTPServer() {
// 	handler := &app.Handler{
// 		Name: "My App",
// 		// Resources: http.FileSystem{webFiles},
// 		// AutoUpdate: true,
// 	}

// 	if err := http.ListenAndServe(":8080", handler); err != nil {
// 		log.Fatal(err)
// 	}
// }
