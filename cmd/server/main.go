package main

import (
	"fmt"
	"github.com/asishcse60/go-rest-api/internal/comment"
	"github.com/asishcse60/go-rest-api/internal/database"
	transportHttp "github.com/asishcse60/go-rest-api/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)
//https://www.youtube.com/watch?v=Z3SYDTMP3ME
//App - contain application information
//to do database connections

type App struct {
	Name string
	version string
}

// Run Sets up of our application
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{"AppName":a.Name, "AppVersion": a.version}).Info("Setting Up our App")
	fmt.Println("Setting Up our App")
	var err error
	db,err:=database.NewDatabase()
	if err != nil{
		return err
	}

	err = database.MigrateDB(db)
	if err != nil{
		return err
	}

	commentService := comment.NewService(db)
	handler := transportHttp.NewHandler(commentService)
	handler.SetUpRoutes()

	if err:=http.ListenAndServe(":8080", handler.Router);err!=nil{
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}
func main(){
	fmt.Println("GO rest api")
	app := App{Name: "go-rest-api", version: "1.0"}
	if err := app.Run();err != nil{
		log.Error("Error starting of our rest api")
		log.Fatal(err)
	}
}
