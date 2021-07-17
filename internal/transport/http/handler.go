package http

import (
	"encoding/json"
	"github.com/asishcse60/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Handler - stores pointer to our comments service
type Handler struct{
	Router *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
	Error string
}
// NewHandler returns a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{Service: service}
}

func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Endpoint hits!")
		log.WithFields(log.Fields{"Method":r.Method, "Path": r.URL.Path}).Info("handled request")
		next.ServeHTTP(w, r)
	})
}

// SetUpRoutes set up routes- all routes for our application
func (h *Handler) SetUpRoutes(){
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
    h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        var responses = &Response{Message: "I am Alive"}
		if err := json.NewEncoder(w).Encode(responses);err !=nil{
			panic(err)
		}
	})
}
