package authorization

import (
	"context"
	"net/http"
)

type writerKey struct{}
type sessionIDKey struct{}
type routesVarsKey struct{}

func GetSession(ctx context.Context) *http.Cookie {
	return ctx.Value(sessionIDKey{}).(*http.Cookie)
}

func AddSession(ctx context.Context, cookie *http.Cookie) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, cookie)
}

func GetVars(ctx context.Context) map[string]string {
	return ctx.Value(routesVarsKey{}).(map[string]string)
}

func AddVars(ctx context.Context, vars map[string]string) context.Context {
	return context.WithValue(ctx, routesVarsKey{}, vars)
}

func AddWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, writerKey{}, w)
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	return ctx.Value(writerKey{}).(http.ResponseWriter)
}
