package main

import (
	"app/api/constants"
	"app/api/infrastructure/database"
	"app/api/llog"
	"app/api/presentation/handler"
	"app/api/presentation/server"
	"fmt"
)

func main() {
	sqlHandler, err := database.New()
	if err != nil {
		llog.Fatal(err)
	}

	appHandler := handler.NewAppHandler(sqlHandler)

	srv := server.New(fmt.Sprintf(":%s", constants.ServerPort))
	srv.Route(appHandler)
	srv.Serve()
}
