package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {

	var req addProductReq
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

	prod := product{
		Id:      generateUniqueID(),
		UserId:  userId,
		Name:    req.Name,
		Details: req.Details,
		Lat:     req.Lat,
		Lng:     req.Lng,
	}

	resp, err := h.service.addProduct(r.Context(), prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	// Parse query params into ProductQuery
	var q ProductQuery
	q.Name = r.URL.Query().Get("name")
	q.Details = r.URL.Query().Get("details")
	// Parse lat/lng as float64
	if lat := r.URL.Query().Get("lat"); lat != "" {
		fmt.Sscanf(lat, "%f", &q.Lat)
	}
	if lng := r.URL.Query().Get("lng"); lng != "" {
		fmt.Sscanf(lng, "%f", &q.Lng)
	}

	resp, err := h.service.getNearestProducts(r.Context(), &q)
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

// generateUniqueID generates a unique ID for a product.
func generateUniqueID() string {
	return fmt.Sprintf("%d", generateRandomInt64())
}

// generateRandomInt64 generates a random int64 value.
func generateRandomInt64() int64 {
	// You can use crypto/rand for better randomness in production.
	return int64(time.Now().UnixNano())
}
