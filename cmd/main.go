package main

import (
	rest "articles_psql/internal/handler"
	"articles_psql/internal/repository"
	"articles_psql/internal/server"
	service2 "articles_psql/internal/service"
	"articles_psql/pkg/db"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description First try to make Swagger

// @host 127.0.0.1:8000
// @BasePath /

func main() {
	db, err := db.InitDB("sql4499495", "kxfp2x7qNW", "sql4.freemysqlhosting.net", "3306", "sql4499495")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()
	log.Println("Connected to db!")

	repos := repository.NewRepositories(db)
	service := service2.NewService(service2.Deps{
		Repos: repos,
	})

	handler := rest.NewHandler(service)

	srv := server.NewServer(handler.InitRoutes())

	if err := srv.Run(); err != nil {
		log.Fatalf(err.Error())
	}

}
