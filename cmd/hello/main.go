package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/josephpballantyne/hello/internal/config"
	"github.com/josephpballantyne/hello/internal/http"
)

var constants *config.Constants
var h http.Handler

func init() {
	constants, _ = config.InitViper()
	v := validator.New()
	h.V = v
	http.SetupRoutes(&h, constants)
}

func main() {
	http.StartServer()
}
