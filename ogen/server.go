package ogen

import (
	"context"
	"net/http"

	"github.com/ogen-go/ogen/middleware"
	api "github.com/releaseband/test-openapi/ogen/gen"
	"github.com/rs/zerolog"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target gen --clean integration.yml --convenient-errors off

type server struct {
}

func StartServer(log *zerolog.Logger) {
	apiLog := log.With().Str("api", "http").Logger()

	opts := middleware.ChainMiddlewares(
		Logger(&apiLog),
	)

	srv, err := api.NewServer(&server{}, api.WithMiddleware(opts))
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		panic(err)
	}
}

func Logger(logger *zerolog.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		ctx := req.Context

		log := logger.With().
			Str("method", req.Raw.Method).
			Str("params", req.Raw.URL.Path).
			Logger()

		req.Context = log.WithContext(ctx)

		resp, err := next(req)
		if err != nil {
			return resp, err
		}
		return resp, err
	}
}

func (s *server) LaunchGame(ctx context.Context, params api.LaunchGameParams) error {
	return nil
}

func (s *server) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return nil
}

func (s *server) Bet(ctx context.Context, req api.OptBetRequest) (api.BetOK, error) {
	return api.BetOK{}, nil
}
