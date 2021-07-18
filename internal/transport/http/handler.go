package http

import (
	"encoding/json"
	"errors"
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

// BasicAuth Basic Auth - a handy middleware function
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("basic auth endpoint hits")
		   user,pass,ok := r.BasicAuth()
		   if user =="admin" && pass=="password" && ok{
			   original(w,r)
		   }else{
		   	 log.Info("user or password is not correct")
		   	 w.Header().Set("Content-Type","application/json; charset=UTH-8")
		   	 sendErrorResponse(w, "not authenticated", errors.New("not authenticated"))
		   }

	}
}
// SetUpRoutes set up routes- all routes for our application
func (h *Handler) SetUpRoutes(){
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
    h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", BasicAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        var responses = &Response{Message: "I am Alive"}
		if err := json.NewEncoder(w).Encode(responses);err !=nil{
			panic(err)
		}
	})
}
