package main

import (
	"fmt"
	transportHttp "github.com/asishcse60/go-rest-api/internal/transport/http"
	"net/http"
)

//App - the struct which contains things like pointers
//to do database connections

type App struct {}

//Sets up of our application
func (a *App) Run() error {
	fmt.Println("Setting Up our App")

	handler := transportHttp.NewHandler()
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
