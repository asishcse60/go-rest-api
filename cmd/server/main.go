package main

import "fmt"

//App - the struct which contains things like pointers
//to do database connections

type App struct {}

//Sets up of our application
func (a *App) Run() error {
	fmt.Println("Setting Up our App")
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
