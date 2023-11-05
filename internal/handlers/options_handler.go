package handlers

import (
	"net/http"
	"strings"

	conf "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
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
	w.Header().Add(AccessControlAllowOriginHeader, conf.FrontendServerIP+conf.FrontendServerPort)
	w.Header().
		Add(AccessControlAllowMethodsHeader, strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodDelete}, ", "))
	w.Header().
		Add(AccessControlAllowHeadersHeader, strings.Join([]string{ContentTypeHeader, CookieHeader}, ", "))
	w.Header().Add(AccessControlMaxAgeHeader, "86400")
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
	w.WriteHeader(http.StatusOK)
}

func AddAllowHeaders(w http.ResponseWriter) {
	w.Header().Add(AccessControlAllowOriginHeader, conf.FrontendServerIP+conf.FrontendServerPort)
	w.Header().Add(AccessControlAllowCredentialsHeader, "true")
}