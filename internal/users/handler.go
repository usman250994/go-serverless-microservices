package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

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

	authHeader := r.Header.Get("Authorization")

	userId, err := parseUserIDFromAuthHeader(authHeader)
	if err != nil {
		http.Error(w, "unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	user := profile{
		Id:       userId,
		Name:     req.Name,
		Email:    req.Email,
		UserType: req.UserType,
		Rating:   req.Ratings,
	}

	resp, err := h.service.updateProfile(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func parseUserIDFromAuthHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	tokenStr := parts[1]
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	userId, ok := claims["sub"].(string)
	if !ok || userId == "" {
		return "", fmt.Errorf("user ID not found in token")
	}
	return userId, nil
}
