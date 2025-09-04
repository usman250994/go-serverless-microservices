package user

import (
	"github.com/go-chi/chi/v5"
)

const user string = "user"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Post("/"+user, h.update)
	r.Get("/"+user, h.get)
}
