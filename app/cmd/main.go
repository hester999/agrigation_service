package main

import (
	_ "app/docs"
	"app/internal/config"
	"app/internal/db"
	_ "app/internal/handlers/subscription"
	handler "app/internal/handlers/subscription"
	"app/internal/logger"
	repo "app/internal/repo/subscription"
	"app/internal/router"
	_ "app/internal/usecases/subscription"
	uc "app/internal/usecases/subscription"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Aggregation Service API
// @version 1.0
// @description Сервис агрегации данных

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
		handlerLoger = logger.NewLogger("[subscription]")
	)

	repo := repo.NewSubscriptionRepo(database, repoLoger)
	usecases := uc.NewSubscriptionUsecases(repo, usecaseLoger)
	handler := handler.NewSubscriptionHandler(usecases, handlerLoger)

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	app := router.NewRouter(
		handler,
	)

	r.PathPrefix("/").Handler(app)
	log.Println("Starting server on " + serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
