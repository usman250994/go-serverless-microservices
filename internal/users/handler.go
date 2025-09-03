package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// how to test this go function locally via a postman?
// how to apply validation on fields?

type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service, validate: validator.New()}
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {

	var req updateProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, "validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := profile{
		//get id from req auth headers
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		UserType: req.UserType,
		Rating:   req.Ratings,
	}

	user, err := h.service.updateProfile(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
