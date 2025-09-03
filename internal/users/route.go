package user

import (
	"github.com/go-chi/chi/v5"
)

const user string = "http://localhost:8080/user"

// defines a post call to accept url like this user?id=250994

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Post("/user", h.update)
}
