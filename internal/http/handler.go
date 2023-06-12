package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"github.com/josephpballantyne/hello/internal/hello"
)

type Handler struct {
	V *validator.Validate
}

func (h *Handler) HelloUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u *hello.User

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := h.V.Struct(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		render.JSON(w, r, fmt.Sprintf("Hello, %s", u.Name))
	}
}
