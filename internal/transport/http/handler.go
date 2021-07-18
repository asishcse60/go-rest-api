package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asishcse60/go-rest-api/internal/comment"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
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
// validate access token
func validateToken(accessToken string) bool{
	var mySigningKey = []byte("mission impossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("token parse got an error")
		}
		return mySigningKey, nil
	})
	if err!=nil {
		return false
	}
	return token.Valid
}

// JWTAuth - a decorator function for jwt validation for endpoint
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt authentication endpoints hits")
		authHeader := r.Header["Authorization"]
		if authHeader == nil{
			sendErrorResponse(w, "jwt not authenticated", errors.New("jwt not authenticated"))
			return
		}
		authoredParts := strings.Split(authHeader[0]," ")
		if len(authoredParts) != 2 || strings.ToLower(authoredParts[0])!="bearer"{
			sendErrorResponse(w, "jwt not authenticated", errors.New("jwt not authenticated"))
			return
		}
		if validateToken(authoredParts[1]) {
			original(w, r)
		}else{
			log.Info("user or password is not correct")
			sendErrorResponse(w, "jwt not authenticated", errors.New("jwt not authenticated"))
			return
		}

	}
}
// SetUpRoutes set up routes- all routes for our application
func (h *Handler) SetUpRoutes(){
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
    h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", BasicAuth(h.GetAllComments)).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.GetComment)).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        var responses = &Response{Message: "I am Alive"}
		if err := json.NewEncoder(w).Encode(responses);err !=nil{
			panic(err)
		}
	})
}
