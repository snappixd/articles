package main

import (
	rest "articles_psql/internal/handler"
	"articles_psql/internal/repository"
	"articles_psql/internal/server"
	service2 "articles_psql/internal/service"
	"articles_psql/pkg/db"
	"log"
)

func main() {

	db, err := db.InitDB("mysql", "", "go_articles")
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
