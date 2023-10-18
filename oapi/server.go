package oapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/releaseband/go-micro-tools/http/server/middleware/logger"
	"github.com/releaseband/test-openapi/oapi/gen"
	"github.com/rs/zerolog"
)

type server struct {
}

func StartServer(log *zerolog.Logger) {
	apiLog := log.With().Str("api", "http").Logger()

	swagger, err := gen.GetSwagger()
	if err != nil {
		panic(fmt.Errorf("Error loading swagger spec\n: %w", err))
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(logger.New(&apiLog))

	gen.HandlerFromMux(&server{}, r)

	s := &http.Server{
		Handler: r,
		Addr:    "localhost:5000",
	}

	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (s *server) LaunchGame(w http.ResponseWriter, r *http.Request, params gen.LaunchGameParams) {
	fmt.Println("params", params)
}
