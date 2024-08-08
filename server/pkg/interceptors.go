package server

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/api/idtoken"
)

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
		cookieHeader := req.Header().Get("Cookie")
		if !validateCredentials(cookieHeader) {
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
		cookieHeader := conn.RequestHeader().Get("Cookie")
		if !validateCredentials(cookieHeader) {
			return connect.NewError(connect.CodeUnauthenticated, errNoToken)
		}
		return next(ctx, conn)
	})
}

func parseCookies(cookieHeader string) map[string]string {
	cookies := make(map[string]string)
	pairs := strings.Split(cookieHeader, ";")
	for _, pair := range pairs {
		kv := strings.SplitN(strings.TrimSpace(pair), "=", 2)
		if len(kv) == 2 {
			cookies[kv[0]] = kv[1]
		}
	}
	return cookies
}

func validateCredentials(cookieHeader string) bool {
	if cookieHeader == "" {
		return false
	}

	cookies := parseCookies(cookieHeader)
	if _, ok := cookies["credential"]; !ok {
		return false
	}

	credential := cookies["credential"]
	ctx := context.Background()
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		return false
	}

	payload, err := validator.Validate(ctx, credential, googleClientID)
	if err != nil {
		return false
	}

	fmt.Println(payload)
	// verify JWT by checking aud, expiry, signed by Google, email, email verified, iss
	payloadChecks := []struct {
		have any
		want any
	}{
		{
			have: payload.Claims["aud"],
			want: googleClientID,
		},
		{
			have: payload.Claims["hd"],
			want: "khanacademy.org",
		},
		{
			have: payload.Claims["email_verified"],
			want: true,
		},
		{
			have: payload.Claims["iss"],
			want: "https://accounts.google.com",
		},
	}
	for _, check := range payloadChecks {
		if check.have != check.want {
			return false
		}
	}
	return true
}
