package server

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
)

const tokenHeader = "Acme-Token"

var errNoToken = errors.New("no token provided, please login")

type authInterceptor struct{}

func NewAuthInterceptor() *authInterceptor {
	return &authInterceptor{}
}

// called when user sends a chat message
func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		if !req.Spec().IsClient && req.Header().Get(tokenHeader) == "" {
			fmt.Println(ctx, "LASKDJLASKDJLKASJDKLSAJ")
			return nil, connect.NewError(connect.CodeUnauthenticated, errNoToken)
		}
		return next(ctx, req)
	})
}

func (i *authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	})
}

// called when user joins a new chat and creates a stream
func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		if conn.RequestHeader().Get(tokenHeader) == "" {
			fmt.Println(ctx, "aslkjdaksl")
			return connect.NewError(connect.CodeUnauthenticated, errNoToken)
		}
		return next(ctx, conn)
	})
}
