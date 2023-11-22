package authorization

import (
	"context"
	"net/http"
)

type writerKey struct{}
type sessionIDKey struct{}
type routesVarsKey struct{}

func GetSession(ctx context.Context) *http.Cookie {
	if ctx.Value(sessionIDKey{}) != nil {
		return ctx.Value(sessionIDKey{}).(*http.Cookie)
	}

	return &http.Cookie{}
}

func AddSession(ctx context.Context, cookie *http.Cookie) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, cookie)
}

func GetVars(ctx context.Context) map[string]string {
	if ctx.Value(routesVarsKey{}) != nil {
		return ctx.Value(routesVarsKey{}).(map[string]string)
	}

	return nil
}

func AddVars(ctx context.Context, vars map[string]string) context.Context {
	return context.WithValue(ctx, routesVarsKey{}, vars)
}

func AddWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, writerKey{}, w)
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	if ctx.Value(writerKey{}) != nil {
		return ctx.Value(writerKey{}).(http.ResponseWriter)
	}

	return nil
}
