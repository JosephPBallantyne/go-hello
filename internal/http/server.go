package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/josephpballantyne/hello/internal/config"
	log "github.com/sirupsen/logrus"
)

type server struct {
	router *chi.Mux
}

var Config *config.Constants
var Server *server

func SetupRoutes(h *Handler, c *config.Constants) {
	Config = c
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type"},
		MaxAge:         3600,
	}))

	r.Route("/v1/api", func(r chi.Router) {
		r.Post("/user", h.HelloUser())
	})

	Server = &server{
		router: r,
	}
}

func StartServer() {
	log.WithFields(log.Fields{"PORT": Config.PORT}).Info("Server starting")
	log.Fatal(http.ListenAndServe(":"+Config.PORT, Server.router))
}
