package main

import (
	"app/api/infrastructure/database"
	"app/api/llog"
	"app/api/presentation/handler"
	"app/api/presentation/server"
)

func main() {
	sqlHandler, err := database.New()
	if err != nil {
		llog.Fatal(err)
	}

	appHandler := handler.NewAppHandler(sqlHandler)

	srv := server.New(":8080")
	srv.Route(appHandler)
	srv.Serve()
}
