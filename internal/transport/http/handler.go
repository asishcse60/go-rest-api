package http

import (
	"encoding/json"
	"fmt"
	"github.com/asishcse60/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//Handler - stores pointer to our comments service
type Handler struct{
	Router *mux.Router
	Service *comment.Service
}

// NewHandler returns a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{Service: service}
}

// SetUpRoutes set up routes- all routes for our application
func (h *Handler) SetUpRoutes(){
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment)
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id:=vars["id"]
	idInt,err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		response(w, err)
		return
	}
	comment, err := h.Service.GetComment(uint(idInt))
	if err!=nil{
		response(w, err)
		return
	}
	response(w, comment)
}
func response(w http.ResponseWriter, res interface{}) {
	if err, ok := res.(error); ok {
		errorResponse(w, err)
		return
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

// errorResponse writes out an error to the client as plaintext
func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

