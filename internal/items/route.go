package item

import (
	"github.com/go-chi/chi/v5"
)

const prod string = "product"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Post("/"+prod, h.add)
	r.Get("/"+prod, h.get)
}
