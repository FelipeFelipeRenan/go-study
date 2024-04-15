package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieexample/movie/internal/controller/movie"
	"net/http"
)

// Handler defines a movie handler
type Handler struct {
	ctrl *movie.Controller
}

// New creates a new movie HTTP handler
func New(ctrl *movie.Controller) *Handler  {
	return &Handler{ctrl}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request){
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound){
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil{
		log.Printf("Repository get error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil{
		log.Printf("Response encode error: %v", err)
	}
}