package http

import (
	"encoding/json"
	"github.com/asishcse60/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id:=vars["id"]
	idInt,err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		response(w, err)
		return
	}
	err = h.Service.DeleteComment(uint(idInt))
	if err!=nil{
		response(w, err)
		return
	}
	response(w, http.StatusOK)
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil{
		response(w, err)
		return
	}
	response(w, comments)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var input comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response(w, err)
	}
	c, err:= h.Service.PostComment(input)
	if err != nil {
		response(w, err)
		return
	}
	response(w, c)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	var input comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response(w, err)
	}
	vars := mux.Vars(r)
	id:=vars["id"]
	idInt,err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		response(w, err)
		return
	}
	comment, err := h.Service.UpdateComment(uint(idInt), input)
	if err != nil {
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

func sendErrorResponse(w http.ResponseWriter, message string, err error){
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type","application/json; charset=UTH-8")
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil{
		panic(err)
	}
}