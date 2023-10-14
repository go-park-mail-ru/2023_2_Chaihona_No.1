package handlers

import (
	"net/http"
	"strings"
)

const (
	AccessControlAllowOriginHeader      = "Access-Control-Allow-Origin"
	AccessControlAllowMethodsHeader     = "Access-Control-Allow-Methods"
	AccessControlAllowHeadersHeader     = "Access-Control-Allow-Headers"
	ContentTypeHeader                   = "Content-Type"
	AccessControlMaxAgeHeader           = "Access-Control-Max-Age"
	CookieHeader                        = "Cookie"
	AccessControlAllowCredentialsHeader = "Access-Control-Allow-Credentials"
)

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(AccessControlAllowOriginHeader, FrontendServerIP+FrontendServerPort)
	w.Header().
		Add(AccessControlAllowMethodsHeader, strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodOptions}, ", "))
	w.Header().
		Add(AccessControlAllowHeadersHeader, strings.Join([]string{ContentTypeHeader, CookieHeader}, ", "))
	w.Header().Add(AccessControlMaxAgeHeader, "86400")
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
	w.WriteHeader(http.StatusOK)
}

func AddAllowHeaders(w http.ResponseWriter) {
	w.Header().Add(AccessControlAllowOriginHeader, FrontendServerIP+FrontendServerPort)
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
}
