package main

import (
	"app/internal/config"
	"app/internal/db"
	"app/internal/handlers/service_handler"
	"app/internal/logger"
	"app/internal/repo/service_repo"
	"app/internal/router"
	"app/internal/usecases/service_usecases"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.LoadConfig("../cfg.yaml")
	if err != nil {
		fmt.Println(err)
	}
	serverAddr, dbConn := config.CfgStringBuilder(*cfg)

	database, err := db.Connection(dbConn)
	if err != nil {
		panic(err)
	}

	var (
		repoLoger    = logger.NewLogger("[repo]")
		usecaseLoger = logger.NewLogger("[usecase]")
		handlerLoger = logger.NewLogger("[handler]")
	)

	repo := service_repo.NewServiceRepo(database, repoLoger)
	usecases := service_usecases.NewService(repo, usecaseLoger)
	handler := service_handler.NewServiceRepo(usecases, handlerLoger)

	r := mux.NewRouter()
	//r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	app := router.NewRouter(
		handler,
	)

	r.PathPrefix("/").Handler(app)
	log.Println("Starting server on " + serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
