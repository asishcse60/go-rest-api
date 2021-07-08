package main

import (
	"fmt"
	"github.com/asishcse60/go-rest-api/internal/comment"
	"github.com/asishcse60/go-rest-api/internal/database"
	transportHttp "github.com/asishcse60/go-rest-api/internal/transport/http"
	"net/http"
)

//App - the struct which contains things like pointers
//to do database connections

type App struct {}

// Run Sets up of our application
func (a *App) Run() error {
	fmt.Println("Setting Up our App")
	var err error
	db,err:=database.NewDatabase()
	if err != nil{
		return err
	}
	commentService := comment.NewService(db)
	handler := transportHttp.NewHandler(commentService)
	handler.SetUpRoutes()

	if err:=http.ListenAndServe(":8000", handler.Router);err!=nil{
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}
func main(){
	fmt.Println("GO rest api")
	app := App{}
	if err := app.Run();err != nil{
		fmt.Println("Error starting of our rest api")
		fmt.Println(err)
	}
}
