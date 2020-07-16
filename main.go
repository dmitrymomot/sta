package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TV4/graceful"
	"github.com/dmitrymomot/sta/httpinterface"
	"github.com/dmitrymomot/sta/interactor"
	"github.com/dmitrymomot/sta/interactor/h"
	"github.com/dmitrymomot/sta/interactor/k"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

// Build tag is set up while compiling
var buildTag string

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(30*time.Second),
		middleware.Throttle(30),
		cors.AllowAll().Handler,
		middleware.AllowContentType(httpinterface.ContentTypeJSON),
	)

	r.NotFound(httpinterface.NotFound)
	r.MethodNotAllowed(httpinterface.MethodNotAllowed)

	r.Get("/", httpinterface.Root(buildTag))
	r.Get("/health", httpinterface.HealthCheck)
	r.Post("/interact", httpinterface.Interactor(
		interactor.New(
			h.NewCustom2(h.NewBase(nil)),
			k.NewCustom2(k.NewCustom1(k.NewBase(nil))),
		),
	))

	graceful.LogListenAndServe(&http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
	})
}
