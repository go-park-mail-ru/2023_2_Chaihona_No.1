package handlers

import "net/http"

const (
	AccessControlAllowOriginHeader      = "Access-Control-Allow-Origin"
	AccessControlAllowMethodsHeader     = "Access-Control-Allow-Methods"
	AccessControlAllowHeadersHeader     = "Access-Control-Allow-Headers"
	ContentTypeHeader                   = "Content-Type"
	AccessControlMaxAgeHeader           = "Access-Control-Max-Age"
	CookieHeader                        = "Cookie"
	AccessControlAllowCredentialsHeader = "Access-Control-Allow-Credentials"
)

func headersToString(headers ...string) string {
	res := ""
	for _, header := range headers {
		res += header + ", "
	}
	if len(res) > 0 {
		res = res[:len(res)-2]
	}
	return res
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(AccessControlAllowOriginHeader, FrontendServerIP+FrontendServerPort)
	w.Header().Add(AccessControlAllowMethodsHeader, headersToString(http.MethodGet, http.MethodPost, http.MethodOptions))
	w.Header().Add(AccessControlAllowHeadersHeader, headersToString(ContentTypeHeader, CookieHeader))
	w.Header().Add(AccessControlMaxAgeHeader, "86400")
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
}

func AddAllowHeaders(w http.ResponseWriter) {
	w.Header().Add(AccessControlAllowOriginHeader, FrontendServerIP+FrontendServerPort)
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
}
