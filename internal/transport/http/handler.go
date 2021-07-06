package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//Handler - stores pointer to our comments service
type Handler struct{
	Router *mux.Router
}

//returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//set up routes- all routes for our application
func (h *Handler) SetUpRoutes(){
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}


